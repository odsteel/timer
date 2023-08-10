package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

type output struct {
	rows, cols, x, y uint16
}

const (
	rows = 6
	cols = 40
)

func (o *output) init() {
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(o)),
	)
	fmt.Printf("\x1b[8;%d;%dt\x1b[H\x1b[J\x1b[?25l", rows, cols)
}

func (o *output) restore() {
	fmt.Printf("\x1b[8;%d;%dt\x1b[H\x1b[J\x1b[?25h", o.rows, o.cols)
}
