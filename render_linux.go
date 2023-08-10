package main

import (
	"fmt"
	"time"
)

const (
	bgRed   = 22
	bgGreen = 22
	bgBlue  = 22
)

const (
	fgRed   = 220
	fgGreen = 220
	fgBlue  = 220
)

const (
	hlRed   = 220
	hlGreen = 22
	hlBlue  = 220
)

func render(lap, total int, timer, remaining time.Duration, paused bool) {
	timer = timer.Round(time.Second)
	remaining = remaining.Round(time.Second)

	fmt.Printf("\x1b[48;2;%d;%d;%dm\x1b[H\x1b[J", bgRed, bgGreen, bgBlue)

	fmt.Printf("\x1b[38;2;%d;%d;%dm", fgRed, fgGreen, fgBlue)
	fmt.Printf("\n  [%d/%d] ◴%s:    ", lap, total, timer)

	fmt.Printf("\x1b[38;2;%d;%d;%dm", hlRed, hlGreen, hlBlue)
	if paused {
		fmt.Printf("▶ %s", remaining)
	} else {
		fmt.Printf("◼ %s", remaining)
	}

	fmt.Print("\n\n  ")
	progress := int((cols - 4) / float64(timer) * float64(remaining))
	for i := 0; i < progress; i++ {
		fmt.Print("█")
	}
	for i := 0; i < (cols - 4 - progress); i++ {
		fmt.Print("░")
	}

	fmt.Print("\x1b[m")
}
