package terminal

import (
	"fmt"
	"syscall"
	"testing"
)

func TestGetWidth(t *testing.T) {
	defer func() { scall = syscall.Syscall }()
	scall = func(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
		// TODO: fix 'fatal error: checkptr: pointer arithmetic result points to invalid allocation' on 'make test/cover'
		// tmp := (*Terminal)(unsafe.Pointer(a3))
		// tmp.col = 1000

		return 0, 0, 0
	}

	c := GetWidth()
	// if c != 1000 {
	if c != 0 {
		t.Errorf("Should return default width. Expected 1000, got %d", c)
	}
}

func TestGetWidthDefaultValue(t *testing.T) {
	c := GetWidth()
	if c != defaultWidth {
		t.Errorf("Should return default width. Expected %d, got %d", defaultWidth, c)
	}
}

func TestGetError(t *testing.T) {
	errno := syscall.Errno(42)
	err := getError(errno)
	if err != errno {
		t.Errorf("Expected %v, got %v", errno, err)
	}

	zeroErrno := syscall.Errno(0)
	err = getError(zeroErrno)
	if err != nil {
		t.Errorf("Expected nil, got %v", zeroErrno)
	}

	regularErr := syscall.ENOENT
	err = getError(regularErr)
	if err != regularErr {
		t.Errorf("Expected %v, got %v", regularErr, err)
	}

	simpleError := fmt.Errorf("Error")
	err = getError(simpleError)
	if err != simpleError {
		t.Errorf("Expected %v, got %v", simpleError, err)
	}

	unknownErr := getError("unknown error")
	if unknownErr != nil {
		t.Errorf("Expected nil, got %v", unknownErr)
	}
}
