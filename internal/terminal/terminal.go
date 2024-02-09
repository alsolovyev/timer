package terminal

import (
	"syscall"
	"unsafe"
)

const TIOCGWINSZ = 0x40087468

type Terminal struct {
	row  uint16
	col  uint16
}

// Col gets terminal width
func (w Terminal) Col() int {
	return int(w.col)
}

// Row gets terminal height
func (w Terminal) Row() int {
	return int(w.row)
}


func GetSize() (t Terminal, err error) {
	_, _, ec := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(TIOCGWINSZ),
		uintptr(unsafe.Pointer(&t)))

  err = getError(ec)

	if TIOCGWINSZ == 0 && err != nil {
		t = Terminal{80, 25}
	}

	return t, err
}

func getError(ec interface{}) (err error) {
	switch v := ec.(type) {

	case syscall.Errno: // Some implementation return syscall.Errno number
		if v != 0 {
			err = syscall.Errno(v)
		}

	case error: // Some implementation return error
		err = ec.(error)
	default:
		err = nil
	}
	return
}
