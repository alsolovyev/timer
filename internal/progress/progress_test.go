package progress

import (
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestProgressNew(t *testing.T) {
	p := New()

	if p.useGradient {
		t.Error("The progress bar should be initialized in monochrome them")
	}

	if p.cachedView == "" {
		t.Error("The progress bar should have initial (cached) view")
	}
}

func TestProgressWithEmptySymbol(t *testing.T) {
	p := New(WithEmptySymbol("-"))
	if p.EmptySymbol != "-" {
		t.Errorf("Invalid empty symbol. Expected '-', got '%s'", p.EmptySymbol)
	}
}

func TestProgressFullEmptySymbol(t *testing.T) {
	p := New(WithFullSymbol("-"))
	if p.FullSymbol != "-" {
		t.Errorf("Invalid full symbol. Expected '-', got '%s'", p.EmptySymbol)
	}
}

func TestWithDefaultGradient(t *testing.T) {
	p := New(WithDefaultGradient())
	if !p.useGradient {
		t.Error("p.useGradient must be true")
	}
}

func TestWithGradient(t *testing.T) {
	s, _ := colorful.Hex("#5A56E0")
	e, _ := colorful.Hex("#EE6FF8")
	p := New(WithGradient("#5A56E0", "#EE6FF8"))

	if !p.useGradient {
		t.Error("p.useGradient must be true")
	}

	if p.gradientStart != s {
		t.Errorf("Envalid start gradient color. Expected %v, got: %v", s, p.gradientStart)
	}

	if p.gradientEnd != e {
		t.Errorf("Envalid end gradient color. Expected %v, got: %v", s, p.gradientEnd)
	}
}

func TestProgressGetView(t *testing.T) {
	cases := []struct {
		Complete int
		Percents float32
	}{
		{Complete: 0, Percents: 0},
		{Complete: 19, Percents: 24.7},
		{Complete: 80, Percents: 100},
	}

	p := New()
	p.Width = 80

	for _, c := range cases {
		b := p.GetView(c.Percents)

		if p.cachedView != b {
			t.Errorf("Should save generated view to cache. Expected '%s', got '%s'", b, p.cachedView)
		}

		if p.complete != c.Complete {
			t.Errorf("Should correctly count the number of completed symbols. Expected '%d', got '%d'", c.Complete, p.complete)
		}
	}
}

func TestProgressGenerateRemainingBarView(t *testing.T) {
	cases := []struct {
		NumOfSymbols int
	}{
		{NumOfSymbols: 0},
		{NumOfSymbols: 19},
		{NumOfSymbols: 80},
	}
	p := New(WithEmptySymbol("0"))
	for _, c := range cases {
		b := p.GenerateRemainingBarView(c.NumOfSymbols)
		n := countOccurrences(b, '0')
		if n != c.NumOfSymbols {
			t.Errorf("Incorent number of symbols. Expected %d, got %d", c.NumOfSymbols, n)
		}
	}
}

func TestProgressGenerateCompleteBarView(t *testing.T) {
	cases := []struct {
		NumOfSymbols int
	}{
		{NumOfSymbols: 0},
		{NumOfSymbols: 19},
		{NumOfSymbols: 80},
	}

	p := New(WithFullSymbol("0"))
	for _, c := range cases {
		b := p.GenerateCompleteBarView(c.NumOfSymbols)
		n := countOccurrences(b, '0')
		if n != c.NumOfSymbols {
			t.Errorf("Incorent number of symbols. Expected %d, got %d", c.NumOfSymbols, n)
		}
	}

	p = New(WithFullSymbol("0"), WithDefaultGradient())
	for _, c := range cases {
		b := p.GenerateCompleteBarView(c.NumOfSymbols)
		n := countOccurrences(b, '0')
		if n != c.NumOfSymbols {
			t.Errorf("Incorent number of symbols. Expected %d, got %d", c.NumOfSymbols, n)
		}
	}
}

func countOccurrences(s string, c rune) int {
	n := 0
	for _, char := range s {
		if char == c {
			n++
		}
	}
	return n
}
