package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"timer/internal/args"
	"timer/internal/cursor"
	"timer/internal/info"
	"timer/internal/progress"
	"timer/internal/render"
	"timer/internal/ticker"
	"timer/internal/version"
)

func main() {
	a, err := args.Parse()
	if err != nil {
		log.Fatal(err)
	}

	if a.ShowBuildInfo {
		version.PrintInfo()
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	w := os.Stdout
	c := cursor.New(w)
	r := render.New(w, c)
	t := ticker.New(ctx, a.Duration)

	p := progress.New()
	i := info.New(a.Duration,
		info.WithName(a.Name),
		info.WithStartTime(),
		info.WithCountdown(),
	)

	c.Hide()
	defer c.Show()

	go t.Start(func(t *ticker.Tick) {
		r.ClearScreen()
		r.RenderLineln(i.GetView())
		r.RenderLine(p.GetView(t.Percents))
	})

	var x bool
	select {
	case <-s: // Stop signal received
		x = true
		r.ClearLine()
	case <-t.StopChan: // Timer completed
		x = false
	}

	if a.ClearOnComplete {
		r.ClearScreen()
	}

	cancel()
	r.RenderLineln(i.GetEndView(x))
}
