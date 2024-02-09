package main

import (
	"log"
	"os"

	"timer/internal/args"
	"timer/internal/cursor"
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
		log.Panic("Error parins args")
		return
	}

	w := os.Stdout
	c := cursor.New(w)
	r := render.New(w, c)
	t := ticker.New(a.Duration)

	p := progress.New(EMPTY_SYMBOL, FULL_SYMBOL, GRADIENT_BEGIN, GRADIENT_END, EMPTY_COLOR)

	c.Hide()

	t.Start(func(t *ticker.Tick) {
		r.ClearScreen()
		r.RenderLineln(p.GetView(t.Percents))
	}, nil)

	c.Show()
}
