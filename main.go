package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/tonobo/mtr/pkg/icmp"
)

type ReplyChecker func(icmp.ICMPReturn) bool

var (
	interval = flag.Duration("duration", time.Millisecond*100, "ICMP request interval")
	timeout  = flag.Duration("timeout", time.Millisecond*50, "ICMP reply timeout")
	target   = flag.IP("target", nil, "target IP")
	verbose  = flag.Bool("verbose", false, "print icmp replies")
)

func main() {
	flag.Parse()

	if *target == nil {
		fmt.Println("target IP missing")
		return
	}

	start := checkForDowntime(
		*target,
		*interval,
		*timeout,
		func(reply icmp.ICMPReturn) bool {
			return !reply.Success
		},
		*verbose,
	)
	end := checkForDowntime(
		*target,
		*interval,
		*timeout,
		func(reply icmp.ICMPReturn) bool {
			return reply.Success
		},
		*verbose,
	)
	fmt.Printf("   start: %s\n", start)
	fmt.Printf("     end: %s\n", end)
	fmt.Printf("duration: %s\n", end.Sub(start))
}

func checkForDowntime(target net.IP, interval, timeout time.Duration, check ReplyChecker, verbose bool) time.Time {
	ticker := time.NewTicker(interval)
	// just use a random sequence
	seq := rand.Int()
	for {
		select {
		case <-ticker.C:
			addr := &net.IPAddr{IP: target}
			hop, err := icmp.SendICMP("", addr, "", 55, os.Getpid(), timeout, seq)
			if verbose {
				fmt.Printf("hop=%v err=%s\n", hop, err)
			}
			timestamp := time.Now()
			seq += 1

			if check(hop) {
				return timestamp
			}
		}
	}
}
