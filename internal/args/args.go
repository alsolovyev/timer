package args

import (
	"flag"
	"fmt"
	"time"
)

type Args struct {
	Duration      time.Duration
	Name          string
	ShowBuildInfo bool
}

func Parse() (*Args, error) {
	n := flag.String("n", "", "The name of the timer")
	v := flag.Bool("v", false, "Display the app version")

	flag.Parse()

	d, err := parseDuration(flag.Args(), *v)
	if err != nil {
		return nil, err
	}

	return &Args{
		Duration:      d,
		Name:          *n,
		ShowBuildInfo: *v,
	}, nil
}

func parseDuration(args []string, r bool) (time.Duration, error) {
	var t time.Duration

	if r {
		return t, nil
	}

	if len(args) == 0 {
		return t, fmt.Errorf("first argument (time duration) is required")
	}

	var err error
	t, err = time.ParseDuration(args[0])

	return t, err
}
