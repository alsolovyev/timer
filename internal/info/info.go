package info

import (
	"fmt"
	"time"
)

const (
	TIME_FORMAT = "03:04PM"
)

type Info struct {
	StartTime string
	EndTime   time.Time
}

func New(s time.Time, d time.Duration) *Info {
	i := Info{}

	i.StartTime = s.Format(TIME_FORMAT)
	i.EndTime = s.Add(d)

	return &i
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

func (i *Info) GetView() string {
	return i.StartTime + ": " + i.GetRemainTime()
}
