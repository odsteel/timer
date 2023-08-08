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
	stepsUsage = "Time intervals per round as a comma-separated list."
	repsUsage  = "Number of rounds."
)

const (
	stepsDefault = 5 * time.Minute
	repsDefault  = 1
)

var (
	stepsFlag interval
	repsFlag  uint
)

func init() {
	flag.Var(&stepsFlag, "steps", stepsUsage)
	flag.UintVar(&repsFlag, "reps", repsDefault, repsUsage)
	flag.Usage = usage
	flag.Parse()
	if len(stepsFlag) == 0 {
		stepsFlag = append(stepsFlag, stepsDefault)
	}
}

func main() {
}

func usage() {
	var msg strings.Builder
	fmt.Fprintf(&msg, "\nUsage: %s [flags]\n\n", path.Base(os.Args[0]))
	fmt.Fprintf(&msg, "\t-steps\n\t\t%s\n\t\tDefault value: %s\n\n",
		stepsUsage, stepsDefault)
	fmt.Fprintf(&msg, "\t-steps\n\t\t%s\n\t\tDefault value: %d\n\n",
		repsUsage, repsDefault)
	fmt.Print(msg.String())
	os.Exit(1)
}
