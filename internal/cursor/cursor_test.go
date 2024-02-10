package cursor

import (
	"bytes"
	"testing"
)

func TestCursorUp(t *testing.T) {
	buffer := &bytes.Buffer{}
	c := New(buffer)

	cases := []struct {
		name string
		n    int
		want string
	}{
		{"MoveUpOneLine", 1, "\x1b[1A"},
		{"MoveUpThreeLines", 3, "\x1b[3A"},
		{"NonPositiveValue", -2, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer.Reset()
			c.Up(tc.n)
			got := buffer.String()
			if got != tc.want {
				t.Errorf("Up(%d) = %q, want %q", tc.n, got, tc.want)
			}
		})
	}
}

func TestCursorDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	c := New(buffer)

	cases := []struct {
		name string
		n    int
		want string
	}{
		{"MoveDownOneLine", 1, "\x1b[1B"},
		{"MoveDownThreeLines", 3, "\x1b[3B"},
		{"NonPositiveValue", -2, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer.Reset()
			c.Down(tc.n)
			got := buffer.String()
			if got != tc.want {
				t.Errorf("Down(%d) = %q, want %q", tc.n, got, tc.want)
			}
		})
	}
}

func TestCursorRight(t *testing.T) {
	buffer := &bytes.Buffer{}
	c := New(buffer)

	cases := []struct {
		name string
		n    int
		want string
	}{
		{"MoveRightOneCharacter", 1, "\x1b[1C"},
		{"MoveRightFiveCharacters", 5, "\x1b[5C"},
		{"NonPositiveValue", -2, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer.Reset()
			c.Right(tc.n)
			got := buffer.String()
			if got != tc.want {
				t.Errorf("Right(%d) = %q, want %q", tc.n, got, tc.want)
			}
		})
	}
}

func TestCursorLeft(t *testing.T) {
	buffer := &bytes.Buffer{}
	c := New(buffer)

	cases := []struct {
		name string
		n    int
		want string
	}{
		{"MoveLeftOneCharacter", 1, "\x1b[1D"},
		{"MoveLeftFiveCharacters", 5, "\x1b[5D"},
		{"NonPositiveValue", -2, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer.Reset()
			c.Left(tc.n)
			got := buffer.String()
			if got != tc.want {
				t.Errorf("Left(%d) = %q, want %q", tc.n, got, tc.want)
			}
		})
	}
}

func TestCursorStartOfLine(t *testing.T) {
	buffer := &bytes.Buffer{}
	c := New(buffer)

	cases := []struct {
		name string
		want string
	}{
		{"MoveToStartOfLine", "\x1b[0G"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer.Reset()
			c.StartOfLine()
			got := buffer.String()
			if got != tc.want {
				t.Errorf("StartOfLine() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestCursorShow(t *testing.T) {
	buffer := &bytes.Buffer{}
	c := New(buffer)

	cases := []struct {
		name string
		want string
	}{
		{"ShowCursor", "\x1b[?25h"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer.Reset()
			c.Show()
			got := buffer.String()
			if got != tc.want {
				t.Errorf("Show() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestCursorHide(t *testing.T) {
	buffer := &bytes.Buffer{}
	c := New(buffer)

	cases := []struct {
		name string
		want string
	}{
		{"HideCursor", "\x1b[?25l"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer.Reset()
			c.Hide()
			got := buffer.String()
			if got != tc.want {
				t.Errorf("Hide() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestCursorClearLine(t *testing.T) {
	buffer := &bytes.Buffer{}
	c := New(buffer)

	cases := []struct {
		name string
		want string
	}{
		{"ClearLine", "\x1b[2K"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer.Reset()
			c.ClearLine()
			got := buffer.String()
			if got != tc.want {
				t.Errorf("ClearLine() = %q, want %q", got, tc.want)
			}
		})
	}
}
