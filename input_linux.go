package main

import (
	"syscall"
	"unsafe"
)

type input syscall.Termios

func (i *input) init() {
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		syscall.TCGETS,
		uintptr(unsafe.Pointer(i)),
	)

	t := *i
	t.Lflag &^= syscall.ECHO
	t.Lflag &^= syscall.ICANON
	t.Cc[syscall.VMIN] = 0
	t.Cc[syscall.VTIME] = 0

	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		syscall.TCSETS,
		uintptr(unsafe.Pointer(&t)),
	)
}

func (i *input) restore() {
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		syscall.TCSETS,
		uintptr(unsafe.Pointer(i)),
	)
}
