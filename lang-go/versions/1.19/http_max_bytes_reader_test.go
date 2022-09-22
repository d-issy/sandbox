package go119_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpMaxBytesReader(t *testing.T) {
	const MAX_BYTES_LIMIT = 10
	maxBytesErrorHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, MAX_BYTES_LIMIT)
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("could not read POST payload"))
			return
		}
		w.Write([]byte(fmt.Sprintf("recevied payout: %q\n", string(payload))))
	})
	ts := httptest.NewServer(maxBytesErrorHandler)

	for _, tt := range []struct {
		reqBody       string
		resStatusCode int
	}{
		{strings.Repeat("x", MAX_BYTES_LIMIT), http.StatusOK},
		{strings.Repeat("x", MAX_BYTES_LIMIT+1), http.StatusBadRequest},
	} {
		reqBody := bytes.NewBufferString(tt.reqBody)
		req, _ := http.NewRequest(http.MethodPost, ts.URL, reqBody)
		res, err := http.DefaultClient.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, tt.resStatusCode, res.StatusCode)
	}
}
