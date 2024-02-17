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
)

const (
	EMPTY_SYMBOL   = "░"
	EMPTY_COLOR    = "#454759"
	FULL_SYMBOL    = "█"
	GRADIENT_BEGIN = "#5A56E0"
	GRADIENT_END   = "#EE6FF8"
)

func main() {
	a, err := args.Parse()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	w := os.Stdout
	c := cursor.New(w)
	r := render.New(w, c)
	t := ticker.New(ctx, a.Duration)

	p := progress.New(EMPTY_SYMBOL, FULL_SYMBOL, GRADIENT_BEGIN, GRADIENT_END, EMPTY_COLOR)
	i := info.New(a.Duration,
		info.WithName(a.Name),
		info.WithStartTime(),
		info.WithCountdown(),
	)

	c.Hide()

	go t.Start(func(t *ticker.Tick) {
		r.ClearScreen()
		r.RenderLineln(i.GetView())
		r.RenderLine(p.GetView(t.Percents))
	})

	c.Show()

	select {
	case <-s:
		r.ClearLine()
		r.RenderLineln(i.GetEndView(true))
		cancel()

	case <-t.StopChan:
		r.RenderLineln(i.GetEndView(false))
		cancel()
	}
}
