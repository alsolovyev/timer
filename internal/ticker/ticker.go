package ticker

import (
	"context"
	"time"
)

const (
	DEFAULT_INTERVAL = 100 * time.Millisecond
)

type Tick struct {
	Current int
	Total   int

	Percents float32

	EndTime   time.Time
	StartTime time.Time
}

type Ticker struct {
	Duration time.Duration
	EndTime  time.Time

	StopChan chan struct{}
	Ticker   *time.Ticker

	ctx context.Context
}

// New creates a new Ticker with the specified interval (in seconds) and durection (in seconds)
func New(ctx context.Context, d time.Duration) *Ticker {
	return &Ticker{
		Duration: DEFAULT_INTERVAL,
		EndTime:  time.Now().Add(d),
		StopChan: make(chan struct{}),

		ctx: ctx,
	}
}

func (t *Ticker) Start(tick func(*Tick)) {
	t.Ticker = time.NewTicker(t.Duration)
	defer t.Stop()

	tc, n := t.startTickInfo()

	tick(t.getTick(tc, 0, n))

	for c := 1; c <= tc; c++ {
		select {
		case <-t.Ticker.C:
			tick(t.getTick(tc, c, n))

		case <-t.ctx.Done():
			return
		}
	}
}

func (t *Ticker) Stop() {
	t.Ticker.Stop()
	close(t.StopChan)
}

func (t *Ticker) startTickInfo() (int, time.Time) {
	n := time.Now()
	tc := int(t.EndTime.Sub(n)/t.Duration + 1)

	return tc, n
}

func (t *Ticker) getTick(tc int, c int, s time.Time) *Tick {
	return &Tick{
		Current: c,
		Total:   tc,

		Percents: float32(c) / float32(tc) * 100,

		StartTime: s,
	}
}
