package palette

import (
	"regexp"
	"testing"
)

func TestPaletteHasPrimary(t *testing.T) {
	if Primary == "" {
		t.Errorf("Primary color is not defined in the Palette")
	}

	if !isHexColor(Primary) {
		t.Errorf("Field 'Primary' does not start with '#'")
	}
}

func TestPaletteHasSecondary(t *testing.T) {
	if Secondary == "" {
		t.Errorf("Secondary color is not defined in the Palette")
	}

	if !isHexColor(Secondary) {
		t.Errorf("Field 'Secondary' does not start with '#'")
	}
}

func isHexColor(c string) bool {
	p := regexp.MustCompile(`^#`)
	return p.MatchString(Secondary)
}
