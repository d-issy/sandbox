package go120_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestResponseController(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		requestHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rc := http.NewResponseController(w)
			rc.SetWriteDeadline(time.Now().Add(3 * time.Second))
			io.Copy(w, r.Body)
		})
		ts := httptest.NewServer(requestHandler)
		reqBody := bytes.NewBufferString("hello")
		req, _ := http.NewRequest(http.MethodGet, ts.URL, reqBody)
		_, err := http.DefaultClient.Do(req)
		assert.NoError(t, err, "十分に時間があるのでエラーにならない")
	})
	t.Run("timeout error", func(t *testing.T) {
		requestHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rc := http.NewResponseController(w)
			rc.SetWriteDeadline(time.Now().Add(30 * time.Microsecond))
			time.Sleep(1 * time.Second)
			io.Copy(w, r.Body)
		})
		ts := httptest.NewServer(requestHandler)
		reqBody := bytes.NewBufferString("hello")
		req, _ := http.NewRequest(http.MethodGet, ts.URL, reqBody)
		_, err := http.DefaultClient.Do(req)
		assert.NotNil(t, err, "timeoutでerrorになるはず")
	})
}
