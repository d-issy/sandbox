package go119_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeDurationTests = []struct {
	duration time.Duration
	want     time.Duration
}{
	{1 * time.Minute, time.Minute},
	{-1 * time.Minute, time.Minute},
}

func TestTimeDurationAbs(t *testing.T) {
	for _, tt := range timeDurationTests {
		got := tt.duration.Abs()
		assert.Equal(t, got, tt.want)
	}
}
