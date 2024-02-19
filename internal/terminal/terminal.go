package terminal

import (
	"syscall"
	"unsafe"
)

const (
	defaultWidth = 80
	tiocgwinsz   = 0x40087468
)

type Terminal struct {
	row uint16
	col uint16
}

// GetWidth retrieves the width of the terminal.
func GetWidth() int {
	var t Terminal
	_, _, ec := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(tiocgwinsz),
		uintptr(unsafe.Pointer(&t)))

	if err := getError(ec); err != nil {
		return defaultWidth
	}

	return int(t.col)
}

func getError(ec interface{}) error {
	switch v := ec.(type) {
	case syscall.Errno:
		if v != 0 {
			return syscall.Errno(v)
		}
		return nil
	case error:
		return ec.(error)
	default:
		return nil
	}
}
