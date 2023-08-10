package main

import (
	"flag"
	"fmt"
	"os"
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
}

func usage() {
	var msg strings.Builder
	fmt.Fprintf(&msg, "\nUsage: %s [flags]\n\n", path.Base(os.Args[0]))
	fmt.Fprintf(&msg, "\t-time\n\t\t%s\n\t\tDefault value: %s\n\n",
		timeUsage, timeDefault)
	fmt.Print(msg.String())
	os.Exit(1)
}
