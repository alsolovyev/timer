package info

import (
	"testing"
	"time"
)

func TestNewInfo(t *testing.T) {
	i := New(time.Minute)

	if i.Prefix != "" {
		t.Errorf("Expected empty prefix, got %s", i.Prefix)
	}

	if i.StartTime.IsZero() {
		t.Error("Expected non-zero start time")
	}

	if i.EndTime.IsZero() {
		t.Error("Expected non-zero end time")
	}
}

func TestInfoGetView(t *testing.T) {
	// With countdown disabled
	p := "Prefix"
	i := &Info{
		Prefix:        p,
		WithCountdown: false,
	}

	r := i.GetView()

	if r != p {
		t.Errorf("Expected %s, got %s", p, r)
	}

	// With countdow enabled
	i = &Info{
		StartTime:     time.Now(),
		EndTime:       time.Now().Add(10 * time.Minute),
		WithCountdown: true,
	}

	r = i.GetView()
	if r != "10m0s" {
		t.Errorf("Expected '10m0s', got %s", r)
	}

}

func TestInfoWithName(t *testing.T) {
	n := "Test Name"
	ep := n + " "
	i := New(time.Duration(1), WithName(n))

	if i.Prefix != ep {
		t.Errorf("Expected %s prefix, got %s", ep, i.Prefix)
	}
}

func TestInfoWithEmptyName(t *testing.T) {
	n := ""
	ep := ""
	i := New(time.Duration(1), WithName(n))

	if i.Prefix != ep {
		t.Errorf("Expected %s prefix, got %s", ep, i.Prefix)
	}
}

func TestInfoWithStartTime(t *testing.T) {
	i := New(time.Duration(1), WithStartTime())
	ep := time.Now().Format(TIME_FORMAT) + " "

	if i.Prefix != ep {
		t.Errorf("Expected %s prefix, got %s", ep, i.Prefix)
	}
}

func TestInfoWithCountdown(t *testing.T) {
	i := New(time.Duration(1), WithCountdown())

	if i.WithCountdown != true {
		t.Error("Expected WithCountdown to be truely")
	}
}

func TestInfoGetRemainTime(t *testing.T) {
	i := New(10*time.Hour, WithStartTime())
	r := i.GetRemainTime()

	if r != "10h0m0s" {
		t.Errorf("Expected '10h0m0s', got %s", r)
	}

	i = New(10*time.Minute, WithStartTime())
	r = i.GetRemainTime()

	if r != "10m0s" {
		t.Errorf("Expected '10m0s', got %s", r)
	}

	i = New(10*time.Second, WithStartTime())
	r = i.GetRemainTime()

	if r != "10s" {
		t.Errorf("Expected '10s', got %s", r)
	}

	i = New(900*time.Millisecond, WithStartTime())
	r = i.GetRemainTime()

	if r != "900ms" {
		t.Errorf("Expected '900ms', got %s", r)
	}

}
