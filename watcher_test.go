package downtime

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tonobo/mtr/pkg/icmp"

	"github.com/RaphaelPour/downtime/mocks"
)

func TestNewWatcher(t *testing.T) {
	interval := time.Second * 1
	timeout := time.Second * 2
	verbosity := true

	w := NewWatcher(
		net.IP{},
		WithInterval(interval),
		WithTimeout(timeout),
		WithVerbosity(verbosity),
		WithIsUpCheck(IsUpCheck),
		WithIsDownCheck(IsDownCheck),
	)

	require.Equal(t, interval, w.interval)
	require.Equal(t, timeout, w.timeout)
	require.Equal(t, verbosity, w.verbose)
	require.True(t, w.isUpCheck(icmp.ICMPReturn{Success: true}))
	require.True(t, w.isDownCheck(icmp.ICMPReturn{Success: false}))
}

func TestRun(t *testing.T) {
	trueCheck := func(_ icmp.ICMPReturn) bool { return true }
	target := net.IP{}
	icmpMock := mocks.NewICMP(t)
	icmpMock.On(
		"SendICMP",
		mock.AnythingOfType("string"),
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("time.Duration"),
		mock.AnythingOfType("int"),
	).Return(icmp.ICMPReturn{}, nil)
	DefaultICMPInterface = icmpMock

	w := NewWatcher(
		target,
		WithIsUpCheck(trueCheck),
		WithIsDownCheck(trueCheck),
		WithVerbosity(true),
	)

	span := w.Run()
	require.Positive(t, span.Duration())
}
