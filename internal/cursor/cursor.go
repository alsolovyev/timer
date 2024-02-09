package cursor

import (
	"fmt"
	"io"
)

type Cursor struct {
	w io.Writer
}

func New(w io.Writer) *Cursor {
	return &Cursor{w: w}
}

// Up moves the cursor n lines up relative to the current position
func (c *Cursor) Up(n int) {
	if n > 0 {
		fmt.Fprintf(c.w, "\x1b[%dA", n)
	}
}

// Down moves the cursor n lines down relative to the current position
func (c *Cursor) Down(n int) {
	if n > 0 {
		fmt.Fprintf(c.w, "\x1b[%dB", n)
	}
}

// Right moves the cursor n characters to the right relative to the current position
func (c *Cursor) Right(n int) {
	if n > 0 {
		fmt.Fprintf(c.w, "\x1b[%dC", n)
	}
}

// Left moves the cursor n characters to the left relative to the current position
func (c *Cursor) Left(n int) {
	if n > 0 {
		fmt.Fprintf(c.w, "\x1b[%dD", n)
	}
}

// StartOfLine moves the cursor to the beginning of the line
func (c *Cursor) StartOfLine() {
	fmt.Fprint(c.w, "\x1b[0G")
}

// Show the cursor if it was hidden previously
func (c *Cursor) Show() {
	fmt.Fprint(c.w, "\x1b[?25h")
}

// Hide the cursor
func (c *Cursor) Hide() {
	fmt.Fprintf(c.w, "\x1b[?25l")
}

// ClearLine clears the current line and moves the cursor to it's start position
func (c *Cursor) ClearLine() {
	fmt.Fprintf(c.w, "\x1b[2K")
}
