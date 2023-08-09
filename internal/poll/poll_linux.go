package poll

import (
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

func poll(input chan<- Input) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	termios := initTermios()
	defer restoreTermios(termios)

	tick := time.Tick(time.Second / 30)

	for {
		select {
		case <-interrupt:
			close(input)
			return
		default:
			buf, _ := io.ReadAll(os.Stdin)
			input <- Input{
				Data: buf,
				Time: <-tick,
			}
		}
	}
}

func initTermios() (prev syscall.Termios) {
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		syscall.TCGETS,
		uintptr(unsafe.Pointer(&prev)),
	)

	t2 := prev
	t2.Lflag &^= syscall.ECHO
	t2.Lflag &^= syscall.ICANON
	t2.Cc[syscall.VMIN] = 0
	t2.Cc[syscall.VTIME] = 0

	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		syscall.TCSETS,
		uintptr(unsafe.Pointer(&t2)),
	)

	return prev
}

func restoreTermios(prev syscall.Termios) {
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		syscall.TCSETS,
		uintptr(unsafe.Pointer(&prev)),
	)
}
