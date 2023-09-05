package downtime

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

type WatcherOption func(*Watcher)

func WithVerbosity(verbose bool) WatcherOption {
	return func(w *Watcher) {
		w.verbose = verbose
	}
}

func WithInterval(interval time.Duration) WatcherOption {
	return func(w *Watcher) {
		w.interval = interval
	}
}

func WithTimeout(timeout time.Duration) WatcherOption {
	return func(w *Watcher) {
		w.timeout = timeout
	}
}

func WithIsUpCheck(check ReplyChecker) WatcherOption {
	return func(w *Watcher) {
		w.isUpCheck = check
	}
}

func WithIsDownCheck(check ReplyChecker) WatcherOption {
	return func(w *Watcher) {
		w.isDownCheck = check
	}
}

type Watcher struct {
	target   net.IP
	interval time.Duration
	timeout  time.Duration
	verbose  bool

	isUpCheck   ReplyChecker
	isDownCheck ReplyChecker
}

func NewWatcher(target net.IP, options ...WatcherOption) Watcher {
	w := Watcher{
		target:      target,
		interval:    DefaultInterval,
		timeout:     DefaultTimeout,
		isUpCheck:   DefaultIsUpCheck,
		isDownCheck: DefaultIsDownCheck,
	}

	for _, option := range options {
		option(&w)
	}

	return w
}

func (w Watcher) Run() TimeSpan {
	return TimeSpan{
		Start: w.RunWithCheck(w.isDownCheck),
		End:   w.RunWithCheck(w.isUpCheck),
	}
}

func (w Watcher) RunWithCheck(check ReplyChecker) time.Time {
	ticker := time.NewTicker(w.interval)
	// just use a random sequence
	seq := rand.Int()
	for {
		<-ticker.C
		addr := &net.IPAddr{IP: w.target}
		hop, err := DefaultICMPInterface.SendICMP("", addr, "", 55, os.Getpid(), w.timeout, seq)
		timestamp := time.Now()
		if w.verbose {
			fmt.Printf("hop=%v err=%s\n", hop, err)
		}
		seq += 1

		if check(hop) {
			return timestamp
		}
	}
}
