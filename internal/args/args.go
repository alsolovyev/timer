package args

import (
	"flag"
	"fmt"
	"time"
)

type Args struct {
	Duration time.Duration
	Name     string
}

func Parse() (*Args, error) {
	n := flag.String("n", "", "The name of the timer")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		return nil, fmt.Errorf("First argument (time duration) is required")
	}

	d, err := time.ParseDuration(args[0])
	if err != nil {
		return nil, err
	}

	return &Args{
		Duration: d,
		Name:     *n,
	}, nil
}
