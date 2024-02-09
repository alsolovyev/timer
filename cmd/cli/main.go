package main

import (
	"log"
	"os"

	"timer/internal/args"
	"timer/internal/cursor"
	"timer/internal/render"
	"timer/internal/ticker"
)

func main() {
	a, err := args.Parse()
	if err != nil {
		log.Panic("Error parins args")
		return
	}

	w := os.Stdout
	c := cursor.New(w)
	r := render.New(w, c)
	t := ticker.New(a.Duration)

	c.Hide()

	t.Start(func(t *ticker.Tick) {
		r.ClearScreen()
		r.RenderLine(a)
	}, nil)

	c.Show()
}
