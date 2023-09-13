package main

import (
	"fmt"

	flag "github.com/spf13/pflag"

	"github.com/RaphaelPour/downtime"
)

var (
	interval = flag.Duration("duration", downtime.DefaultInterval, "ICMP request interval")
	timeout  = flag.Duration("timeout", downtime.DefaultTimeout, "ICMP reply timeout")
	target   = flag.IP("target", nil, "target IP")
	verbose  = flag.Bool("verbose", false, "print icmp replies")
)

func main() {
	flag.Parse()

	if *target == nil {
		fmt.Println("target IP missing")
		return
	}

	w := downtime.NewWatcher(
		*target,
		downtime.WithInterval(*interval),
		downtime.WithTimeout(*timeout),
		downtime.WithVerbosity(*verbose),
	)

	downtime := w.Run()

	fmt.Printf("   start: %s\n", downtime.Start)
	fmt.Printf("     end: %s\n", downtime.End)
	fmt.Printf("duration: %s\n", downtime.Duration())
}
