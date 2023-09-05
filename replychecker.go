package downtime

import (
	"github.com/tonobo/mtr/pkg/icmp"
)

type ReplyChecker func(icmp.ICMPReturn) bool

func IsDownCheck(reply icmp.ICMPReturn) bool {
	return !reply.Success
}

func IsUpCheck(reply icmp.ICMPReturn) bool {
	return reply.Success
}
