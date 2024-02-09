package args

import (
	"os"
	"time"
)

type Args struct {
	Duration time.Duration
}

func Parse() (*Args, error) {
	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		return nil, err
	}

	return &Args{Duration: d}, nil
}
