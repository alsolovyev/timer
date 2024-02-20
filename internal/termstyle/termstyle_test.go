package termstyle

import (
	"testing"
)

func TestTermstyleSetPrimaryColor(t *testing.T) {
	o := primary
	SetPrimaryColor("")

	if primary != o {
		t.Errorf("Expected")
	}

	SetPrimaryColor("#000000")
	if primary != profile.Color("#000000") {
		t.Error("Unexpected primary color")
	}
}

func TestTermstyleSetSecondaryColor(t *testing.T) {
	o := secondary
	SetSecondaryColor("")

	if secondary != o {
		t.Errorf("Expected")
	}

	SetSecondaryColor("#000000")
	if secondary != profile.Color("#000000") {
		t.Error("Unexpected secondary color")
	}
}

func TestTermstyleToStyle(t *testing.T) {
	s := ToStyle("a")

	if s.String() != "a" {
		t.Errorf("Expected 'a', got %s", s.String())
	}
}

func TestTermstyleToColor(t *testing.T) {
	s := ToColor("a", "#000000")

	if s != "a" {
		t.Errorf("Expected 'a', got %s", s)
	}
}

func TestTermstyleToColorBold(t *testing.T) {
	s := ToColorBold("a", "#000000")

	if s != "a" {
		t.Errorf("Expected 'a', got %s", s)
	}
}

func TestTermstyleToPrimary(t *testing.T) {
	s := ToPrimary("a")

	if s != "a" {
		t.Errorf("Expected 'a', got %s", s)
	}
}

func TestTermstyleToPrimaryBold(t *testing.T) {
	s := ToPrimaryBold("a")

	if s != "a" {
		t.Errorf("Expected 'a', got %s", s)
	}
}

func TestTermstyleToSecondary(t *testing.T) {
	s := ToSecondary("a")

	if s != "a" {
		t.Errorf("Expected 'a', got %s", s)
	}
}
func TestTermstyleToSecondaryBold(t *testing.T) {
	s := ToSecondaryBold("a")

	if s != "a" {
		t.Errorf("Expected 'a', got %s", s)
	}
}
func TestTermstyleToBold(t *testing.T) {
	s := ToBold("a")

	if s != "a" {
		t.Errorf("Expected 'a', got %s", s)
	}
}
