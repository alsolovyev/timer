package render

import (
	"fmt"
	"io"
)

var (
  TEXT_BOLD = "\033[1m"
  TEXT_RESET = "\033[0m"
)

type Cursor interface {
	Up(n int)
	Down(n int)
	Left(n int)
	Right(n int)
	StartOfLine()

	Show()
	Hide()

	ClearLine()
}

type Render struct {
	Cursor Cursor
	Lines  int
	Writer io.Writer
}

func New(w io.Writer, c Cursor) *Render {
	return &Render{Cursor: c, Lines: 0, Writer: w}
}

func (r *Render) RenderLine(a ...any) {
	fmt.Fprint(r.Writer, a...)
}

func (r *Render) RenderLineln(a ...any) {
	fmt.Fprintln(r.Writer, a...)

  r.Lines++
}

func (r *Render) ClearLine() {
  r.Cursor.Up(1)
  r.Cursor.StartOfLine()
  r.Cursor.ClearLine()
  r.Lines--
}

func (r *Render) ClearScreen() {
	for r.Lines > 0 {
    r.ClearLine()
	}
}

