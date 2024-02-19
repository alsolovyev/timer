package renderer

import (
	"fmt"
	"io"
)

const (
	TEXT_BOLD  = "\033[1m"
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

type Renderer struct {
	Cursor Cursor
	Lines  int
	Writer io.Writer
}

func New(w io.Writer, c Cursor) *Renderer {
	return &Renderer{Cursor: c, Lines: 0, Writer: w}
}

func (r *Renderer) RenderLine(a ...any) {
	fmt.Fprint(r.Writer, a...)
}

func (r *Renderer) RenderLineln(a ...any) {
	fmt.Fprintln(r.Writer, a...)
	r.Lines++
}

// ClearLine clears the content of the current line by moving the cursor to the beginning.
func (r *Renderer) ClearLine() {
	r.Cursor.StartOfLine()
	r.Cursor.ClearLine()
}

// DeleteLine clears the current line, moves the cursor up one line.
func (r *Renderer) DeleteLine() {
	r.Cursor.Up(1)
	r.Cursor.StartOfLine()
	r.Cursor.ClearLine()
	r.Lines--
}

// ClearScreen clears the entire screen.
func (r *Renderer) ClearScreen() {
	for r.Lines > 0 {
		r.DeleteLine()
	}
}
