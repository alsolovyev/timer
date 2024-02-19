package progress

import (
	"strings"

	"timer/internal/palette"
	"timer/internal/terminal"
	"timer/internal/termstyle"
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
}

func New() *Progress {
	p := Progress{
		Width: GetWidth(),

		EmptySymbol: "░",
		EmptyColor:  palette.Secondary,

		FullSymbol: "█",
		FullColor:  palette.Primary,
	}

	return &p
}

func (p *Progress) GetView(pr float32) string {
	c := int(float32(p.Width) / 100 * pr)

	return p.GenerateCompleteBarView(c) + p.GenerateRemainingBarView(p.Width-c)
}

func (p *Progress) GenerateRemainingBarView(c int) string {
	return termstyle.ToColor(strings.Repeat(p.EmptySymbol, c), p.EmptyColor)
}

func (p *Progress) GenerateCompleteBarView(c int) string {
	return termstyle.ToColor(strings.Repeat(p.FullSymbol, c), p.FullColor)
}

func GetWidth() int {
	s, err := terminal.GetSize()
	if err != nil {
		return defaultWidth
	}

	return s.Col()
}
