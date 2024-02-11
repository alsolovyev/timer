package progress

import (
	"strings"

	"timer/internal/terminal"
	"timer/internal/termstyle"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	DEFAULT_WIDTH = 80
)

type Progress struct {
	WinWidth int

	ColorA colorful.Color
	ColorB colorful.Color
	ColorC string

	EmptySymbol string
	FullSymbol  string

	p int
	c string
}

func New(e, f, ca, cb, cc string) *Progress {
	p := Progress{}

	p.WinWidth = p.GetWidth()

	p.ColorA, _ = colorful.Hex(ca)
	p.ColorB, _ = colorful.Hex(cb)
	p.ColorC = cc

	p.EmptySymbol = e
	p.FullSymbol = f

	return &p
}

func (p *Progress) GetView(pr float32) string {
	n := int(float32(p.WinWidth) / 100 * pr)

	if n != 0 && p.p == n {
		return p.c
	}

	p.c = p.GetBar(n, true, p.FullSymbol) + p.GetBar(p.WinWidth-n, false, p.EmptySymbol)
	p.p = n

	return p.c

}

func (p *Progress) GetBar(n int, g bool, s string) string {
	r := strings.Builder{}

	for i := 0; i < n; i++ {
		if g {
			cg := p.ColorA.BlendLuv(p.ColorB, float64(i)/float64(p.WinWidth-1)).Hex()
			r.WriteString(termstyle.ToColor(s, cg))
		} else {
			r.WriteString(termstyle.ToColor(s, p.ColorC))
		}
	}

	return r.String()
}

func (p *Progress) GetWidth() int {
	s, err := terminal.GetSize()
	if err != nil {
		return DEFAULT_WIDTH
	}

	return s.Col()
}
