package downtime

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSendICMP(t *testing.T) {
	t.Skip("test needs root")
	hop, err := DefaultICMPInterface.SendICMP(
		"",
		&net.IPAddr{IP: net.ParseIP("127.0.0.1")},
		"",
		55,
		os.Getpid(),
		time.Second,
		0,
	)

	require.NoError(t, err)
	require.True(t, hop.Success)
}
