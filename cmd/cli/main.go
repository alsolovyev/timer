package main

import (
	"fmt"
	"log"
	"os"

	"timer/internal/args"
	"timer/internal/cursor"
)

func main() {
	a, err := args.Parse()
	if err != nil {
		log.Panic("Error parins args")
		return
	}

	w := os.Stdout
	c := cursor.New(w)

	c.Hide()
	fmt.Println(a)
	c.Show()
}
