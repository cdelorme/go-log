package log

import (
	"os"
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var procGetConsoleMode = kernel32.NewProc("GetConsoleMode")

func useColor() bool {
	var st uint32
	r, _, e := syscall.Syscall(procGetConsoleMode.Addr(), 2, uintptr(os.Stderr.Fd()), uintptr(unsafe.Pointer(&st)), 0)
	return r != 0 && e == 0
}
