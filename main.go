package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path"
	"strings"
	"time"
)

const (
	timeUsage   = "Time intervals as a comma-separated list."
	timeDefault = 5 * time.Minute
)

var timeFlag interval

func init() {
	flag.Var(&timeFlag, "time", timeUsage)
	flag.Usage = usage
	flag.Parse()
	if len(timeFlag) == 0 {
		timeFlag = append(timeFlag, timeDefault)
	}
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	tick := time.Tick(50 * time.Millisecond)

	var stdin input
	stdin.init()
	defer stdin.restore()

	var stdout output
	stdout.init()
	defer stdout.restore()

	for i, timer := range timeFlag {
		remaining := timer
		end := time.Now().Add(remaining)
		paused := true

		for t := <-tick; t.Before(end); t = <-tick {
			select {
			case <-interrupt:
				return
			default:
				buf, err := io.ReadAll(os.Stdin)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}
				for _, key := range string(buf) {
					if key == ' ' {
						paused = !paused
					}
				}
			}
			if paused {
				end = t.Add(remaining)
			} else {
				remaining = end.Sub(t)
			}
			render(i+1, len(timeFlag), timer, remaining, paused)
		}
	}
}

func usage() {
	var msg strings.Builder
	fmt.Fprintf(&msg, "\nUsage: %s [flags]\n\n", path.Base(os.Args[0]))
	fmt.Fprintf(&msg, "\t-time\n\t\t%s\n\t\tDefault value: %s\n\n",
		timeUsage, timeDefault)
	fmt.Print(msg.String())
	os.Exit(1)
}
