package downtime

import (
	"net"
	"time"

	"github.com/tonobo/mtr/pkg/icmp"
)

//go:generate mockery --name ICMP
type ICMP interface {
	SendICMP(localAddr string, dst net.Addr, target string, ttl, id int, timeout time.Duration, seq int) (icmp.ICMPReturn, error)
}

type DefaultICMP struct{}

func (d DefaultICMP) SendICMP(
	localAddr string,
	dst net.Addr,
	target string,
	ttl, id int,
	timeout time.Duration,
	seq int,
) (icmp.ICMPReturn, error) {
	return icmp.SendICMP(localAddr, dst, target, ttl, id, timeout, seq)
}
