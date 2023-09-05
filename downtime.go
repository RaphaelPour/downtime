package downtime

import (
	"time"
)

type TimeSpan struct {
	Start time.Time
	End   time.Time
}

func (span TimeSpan) Duration() time.Duration {
	return span.End.Sub(span.Start)
}

var (
	DefaultICMPInterface ICMP = new(DefaultICMP)
	DefaultIsDownCheck        = IsDownCheck
	DefaultIsUpCheck          = IsUpCheck
	DefaultTimeout            = time.Millisecond * 50
	DefaultInterval           = time.Millisecond * 100
)
