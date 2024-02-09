package main

import (
	"log"
	"os"

	"timer/internal/args"
	"timer/internal/cursor"
	"timer/internal/render"
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

	c.Hide()

	r.RenderLine(a)

	c.Show()
}
