package main

import (
	"fmt"
	"log"

	"timer/internal/args"
)

func main() {
	a, err := args.Parse()
	if err != nil {
		log.Panic("Error parins args")
		return
	}

	fmt.Println(a)
}
