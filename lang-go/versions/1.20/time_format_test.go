package go120_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeFormat(t *testing.T) {
	tt := time.Date(2023, 2, 4, 18, 33, 30, 0, time.UTC)
	assert.Equal(t, "2023-02-04 18:33:30", tt.Format(time.DateTime))
	assert.Equal(t, "2023-02-04", tt.Format(time.DateOnly))
	assert.Equal(t, "18:33:30", tt.Format(time.TimeOnly))
}
