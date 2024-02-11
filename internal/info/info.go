package info

import (
	"fmt"
	"time"
)

const (
	TIME_FORMAT = "03:04PM"
)

type Info struct {
	EndTime   time.Time
	Prefix    string
	StartTime time.Time
}

type InfoOption func(*Info)

func New(d time.Duration, opts ...InfoOption) *Info {
	n := time.Now()

	i := &Info{
		Prefix:    "",
		StartTime: n,
		EndTime:   n.Add(d),
	}

	for _, opt := range opts {
		opt(i)
	}

	return i
}

func (i *Info) GetView() string {
	return i.Prefix + i.GetRemainTime()
}

func WithName(n string) InfoOption {
	return func(i *Info) {
		if n == "" {
			return
		}

		i.Prefix += n + " "
	}
}

func WithStartTime() InfoOption {
	return func(i *Info) {
		i.Prefix += time.Now().Format(TIME_FORMAT) + " "
	}
}

func (i *Info) GetRemainTime() string {
	c := time.Now()
	r := i.EndTime.Sub(c)

	h := int(r.Hours())
	m := int(r.Minutes()) % 60
	s := int(r.Seconds()) % 60
	ms := (r.Milliseconds() + 50) / 100 * 100 // rounding to the nearest hundred

	if r.Hours() >= 24 {
		d := int(r.Hours() / 24)
		return fmt.Sprintf("%dd%dh%dm%ds", d, h, m, s)
	}

	if h > 0 {
		return fmt.Sprintf("%dh%dm%ds", h, m, s)
	}

	if m > 0 {
		return fmt.Sprintf("%dm%ds", m, s)
	}

	if s > 0 {
		return fmt.Sprintf("%ds", s)
	}

	return fmt.Sprintf("%dms", ms)
}
