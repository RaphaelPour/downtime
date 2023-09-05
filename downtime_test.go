package downtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeSpan(t *testing.T) {
	now := time.Now()
	span := TimeSpan{
		Start: now,
		End:   now.Add(time.Minute * 1),
	}

	require.Equal(t, 1.0, span.Duration().Minutes())
}
