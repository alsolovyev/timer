package ticker

import (
	"os"
	"os/signal"
	"syscall"
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
	s        chan struct{}
}

// New creates a new Ticker with the specified interval (in seconds) and durection (in seconds)
func New(d time.Duration) *Ticker {
	return &Ticker{
		Duration: DEFAULT_INTERVAL,
		EndTime:  time.Now().Add(d),
		s:        make(chan struct{}),
	}
}

func (t *Ticker) Start(tick func(*Tick), s chan os.Signal) {
	tkr := time.NewTicker(t.Duration)
	tt, ct, n := t.startTickInfo()

	if s == nil {
		s = make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	}

	tick(t.getTick(tt, ct, n))
	ct++

	for {
		select {
		case <-tkr.C:
			tick(t.getTick(tt, ct, n))
			ct++

			if time.Now().After(t.EndTime) {
				tkr.Stop()
				return
			}
		case <-t.s:
			tkr.Stop()
			return

		case <-s:
			tkr.Stop()
			return
		}
	}
}

func (t *Ticker) Stop() {
	close(t.s)
}

func (t *Ticker) startTickInfo() (int, int, time.Time) {
	n := time.Now()
	tc := int(t.EndTime.Sub(n)/t.Duration + 1)

	return tc, 0, n
}

func (t *Ticker) getTick(tc int, c int, s time.Time) *Tick {
	return &Tick{
		Current: c,
		Total:   tc,

		Percents: float32(c) / float32(tc) * 100,

		StartTime: s,
	}
}
