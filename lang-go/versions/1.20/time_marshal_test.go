package go120_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeMarshal(t *testing.T) {
	t.Parallel()
	t.Run("utc", func(t *testing.T) {
		d, _ := json.Marshal(time.Date(2023, 2, 4, 18, 33, 30, 0, time.UTC))
		assert.Equal(t, `"2023-02-04T18:33:30Z"`, string(d))
	})
	t.Run("jst", func(t *testing.T) {
		jst, _ := time.LoadLocation("Asia/Tokyo")
		d, _ := json.Marshal(time.Date(2023, 2, 4, 18, 33, 30, 0, jst))
		assert.Equal(t, `"2023-02-04T18:33:30+09:00"`, string(d))
	})
}
