package info

import (
	"time"
	"timer/internal/termstyle"
)

const (
	TIME_FORMAT = "03:04PM"
)

type Info struct {
	EndTime   time.Time
	Prefix    string
	StartTime time.Time

	WithCountdown bool
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
	if i.WithCountdown {
		return i.Prefix + termstyle.ToSecondary(i.GetRemainTime())
	}

	return i.Prefix
}

func (i *Info) GetEndView(e bool) string {
	if e {
		return fmt.Sprintf("%s %s",
			termstyle.ToBold("Stopped"),
			termstyle.ToSecondary("Total time elapsed: " + i.GetElapsedTime()),
		)
	}

	return ""
}

func WithName(n string) InfoOption {
	return func(i *Info) {
		if n == "" {
			return
		}

		i.Prefix += termstyle.ToPrimaryBold(n) + " "
	}
}

func WithStartTime() InfoOption {
	return func(i *Info) {
		i.Prefix += termstyle.ToSecondary(time.Now().Format(TIME_FORMAT)) + " "
	}
}

func WithCountdown() InfoOption {
	return func(i *Info) {
		i.WithCountdown = true
	}
}

func (i *Info) GetRemainTime() string {
	c := time.Now()
	r := i.EndTime.Sub(c)

	if r <= time.Second {
		return r.Round(100 * time.Millisecond).String()
	}

	return r.Round(time.Second).String()
}

func (i *Info) GetElapsedTime() string {
	return time.Since(i.StartTime).Round(time.Second).String()
}
