package renderer

import (
	"bytes"
	"testing"
)

type MockCursor struct {
	UpCalled          int
	DownCalled        int
	LeftCalled        int
	RightCalled       int
	StartOfLineCalled int
	ShowCalled        int
	HideCalled        int
	ClearLineCalled   int
}

// Up is a mock implementation for the Up method.
func (m *MockCursor) Up(n int) { m.UpCalled++ }

// Down is a mock implementation for the Down method.
func (m *MockCursor) Down(n int) { m.DownCalled++ }

// Left is a mock implementation for the Left method.
func (m *MockCursor) Left(n int) { m.LeftCalled++ }

// Right is a mock implementation for the Right method.
func (m *MockCursor) Right(n int) { m.RightCalled++ }

// StartOfLine is a mock implementation for the StartOfLine method.
func (m *MockCursor) StartOfLine() { m.StartOfLineCalled++ }

// Show is a mock implementation for the Show method.
func (m *MockCursor) Show() { m.ShowCalled++ }

// Hide is a mock implementation for the Hide method.
func (m *MockCursor) Hide() { m.HideCalled++ }

// ClearLine is a mock implementation for the ClearLine method.
func (m *MockCursor) ClearLine() { m.ClearLineCalled++ }

func TestRendererNew(t *testing.T) {
	var b bytes.Buffer

	c := MockCursor{}
	r := New(&b, &c)

	if r == nil {
		t.Error("Expected non-nil Renderer, got nil")
	}

	if r.Cursor != &c {
		t.Errorf("Expected Cursor to be %v, got %v", c, r.Cursor)
	}

	if r.Lines != 0 {
		t.Errorf("Expected Lines to be 0, got %d", r.Lines)
	}

	if r.Writer != &b {
		t.Error("Expected Writer to be the provided io.Writer, got a different io.Writer")
	}

	if b.String() != "" {
		t.Errorf("Expected buffer to be empty, got %s", b.String())
	}
}

func TestRendererRenderLine(t *testing.T) {
	var b bytes.Buffer

	c := MockCursor{}
	r := New(&b, &c)

	r.RenderLine("a")
	r.RenderLine("b")
	r.RenderLine("c")

	if r.Lines != 0 {
		t.Errorf("Expected Lines to be 0, got %d", r.Lines)
	}

	if b.String() != "abc" {
		t.Errorf("Expected Writter to be 'abs', got %s", b.String())
	}
}

func TestRendererRenderLineln(t *testing.T) {
	var b bytes.Buffer

	c := MockCursor{}
	r := New(&b, &c)

	r.RenderLineln("a")
	r.RenderLineln("b")
	r.RenderLineln("c")

	if r.Lines != 3 {
		t.Errorf("Expected Lines to be 3, got %d", r.Lines)
	}

	if b.String() != "a\nb\nc\n" {
		t.Errorf("Expected Writter to be 'a\nb\ns\n', got %s", b.String())
	}
}

func TestRendererClearLine(t *testing.T) {
	var b bytes.Buffer

	c := MockCursor{}
	r := New(&b, &c)

	r.RenderLine("a")
	r.ClearLine()
	r.RenderLine("a")
	r.ClearLine()

	if r.Lines != 0 {
		t.Errorf("Expected Lines to be 0, got %d", r.Lines)
	}

	if c.ClearLineCalled != 2 {
		t.Errorf("Expected Curosr.ClearLine to be called 2 times, called %d", c.ClearLineCalled)
	}

	if c.StartOfLineCalled != 2 {
		t.Errorf("Expected Curosr.StartOfLine to be called 2 times, called %d", c.StartOfLineCalled)
	}
}

func TestRendererClearLineln(t *testing.T) {
	var b bytes.Buffer

	c := MockCursor{}
	r := New(&b, &c)

	r.RenderLineln("a")
	r.DeleteLine()
	r.RenderLineln("a")
	r.DeleteLine()

	if r.Lines != 0 {
		t.Errorf("Expected Lines to be 0, got %d", r.Lines)
	}

	if c.ClearLineCalled != 2 {
		t.Errorf("Expected Curosr.ClearLine to be called 2 times, called %d", c.ClearLineCalled)
	}

	if c.StartOfLineCalled != 2 {
		t.Errorf("Expected Curosr.StartOfLine to be called 2 times, called %d", c.StartOfLineCalled)
	}

	if c.UpCalled != 2 {
		t.Errorf("Expected Curosr.Up to be called 2 times, called %d", c.StartOfLineCalled)
	}
}

func TestRendererClearScreen(t *testing.T) {
	var b bytes.Buffer

	c := MockCursor{}
	r := New(&b, &c)

	r.RenderLineln("a")
	r.RenderLineln("a")
	r.RenderLine("a")
	r.ClearScreen()

	if r.Lines != 0 {
		t.Errorf("Expected Lines to be 0, got %d", r.Lines)
	}

	if c.ClearLineCalled != 2 {
		t.Errorf("Expected Curosr.ClearLine to be called 2 times, called %d", c.ClearLineCalled)
	}

	if c.StartOfLineCalled != 2 {
		t.Errorf("Expected Curosr.StartOfLine to be called 2 times, called %d", c.StartOfLineCalled)
	}

	if c.UpCalled != 2 {
		t.Errorf("Expected Curosr.Up to be called 2 times, called %d", c.StartOfLineCalled)
	}
}
