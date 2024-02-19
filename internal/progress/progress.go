package progress

import (
	"strings"

	"timer/internal/palette"
	"timer/internal/terminal"
	"timer/internal/termstyle"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	defaultWidth = 80
)

type Progress struct {
	Width int

	EmptyColor  string
	EmptySymbol string

	FullColor  string
	FullSymbol string

	useGradient   bool
	gradientStart colorful.Color
	gradientEnd   colorful.Color

	complete   int
	cachedView string
}

type ProgressOption func(*Progress)

// WithEmptySymbol sets the symbol used to construct the empty components of the progress bar.
func WithEmptySymbol(s string) ProgressOption {
	return func(p *Progress) {
		p.EmptySymbol = s
	}
}

// WithFullSymbol sets the symbol used to construct the complete components of the progress bar.
func WithFullSymbol(s string) ProgressOption {
	return func(p *Progress) {
		p.FullSymbol = s
	}
}

// WithGradient sets predefined gradient colors for the complete components of the progress bar.
func WithDefaultGradient() ProgressOption {
	return WithGradient("#5A56E0", "#EE6FF8")
}

// WithGradient sets the gradient colors for the complete components of the progress bar.
func WithGradient(hb, he string) ProgressOption {
	return func(p *Progress) {
		b, _ := colorful.Hex(hb)
		e, _ := colorful.Hex(he)

		p.useGradient = true
		p.gradientStart = b
		p.gradientEnd = e
	}
}

// New creates and returns a new Progress instance with default settings.
func New(opts ...ProgressOption) *Progress {
	p := &Progress{
		Width: GetWidth(),

		EmptySymbol: "░",
		EmptyColor:  palette.Secondary,

		FullSymbol: "█",
		FullColor:  palette.Primary,
	}

	p.cachedView = p.GenerateRemainingBarView(p.Width)

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// GetView generates and returns a progress bar view based on the given completion percentage.
func (p *Progress) GetView(pr float32) string {
	c := int(float32(p.Width) / 100 * pr)

	if p.complete == c {
		return p.cachedView
	}

	b := p.GenerateCompleteBarView(c) + p.GenerateRemainingBarView(p.Width-c)
	p.cachedView = b
	p.complete = c

	return b
}

// GenerateRemainingBarView generates a progress bar view for the remaining part of the bar.
func (p *Progress) GenerateRemainingBarView(c int) string {
	return termstyle.ToColor(strings.Repeat(p.EmptySymbol, c), p.EmptyColor)
}

// GenerateCompleteBarView generates a progress bar view for the completed part of the bar.
func (p *Progress) GenerateCompleteBarView(c int) string {
	// Monochrome
	if !p.useGradient {
		return termstyle.ToColor(strings.Repeat(p.FullSymbol, c), p.FullColor)
	}

	// Gradient
	s := strings.Builder{}

	for i := 0; i < c; i++ {
		g := p.gradientStart.BlendLuv(
			p.gradientEnd,
			float64(i)/float64(p.Width-1),
		).Hex()

		s.WriteString(termstyle.ToColor(p.FullSymbol, g))
	}

	return s.String()
}

// GetWidth retrieves the width of the terminal.
func GetWidth() int {
	s, err := terminal.GetSize()
	if err != nil {
		return defaultWidth
	}

	return s.Col()
}
