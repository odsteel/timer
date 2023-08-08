package main

import (
	"fmt"
	"strings"
	"time"
)

type interval []time.Duration

func (i *interval) String() string {
	return fmt.Sprint(*i)
}

func (i *interval) Set(value string) error {
	for _, v := range strings.Split(value, ",") {
		step, err := time.ParseDuration(strings.TrimSpace(v))
		if err != nil {
			return err
		}
		*i = append(*i, step)
	}
	return nil
}
