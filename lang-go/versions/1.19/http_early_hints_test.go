package go119_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"net/textproto"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpEarlyHints(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	earlyHintHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Link", "</style.css>; rel=preload; as=style")
		w.WriteHeader(http.StatusEarlyHints)
		wg.Wait()

		w.Header().Add("Link", "</script.js>; rel=preload; as=scripy")
		w.WriteHeader(http.StatusEarlyHints)

		w.Write([]byte("hello"))
	})
	ts := httptest.NewServer(earlyHintHandler)

	counter := 0
	trace := &httptrace.ClientTrace{
		Got1xxResponse: func(code int, header textproto.MIMEHeader) error {
			switch counter {
			case 0:
				assert.Equal(t, http.StatusEarlyHints, code)
				assert.EqualValues(t, []string{"</style.css>; rel=preload; as=style"}, header["Link"])
				wg.Done()
			case 1:
				assert.Equal(t, http.StatusEarlyHints, code)
				assert.EqualValues(t, []string{"</style.css>; rel=preload; as=style", "</script.js>; rel=preload; as=scripy"}, header["Link"])
			default:
				t.Error("Unexpected 1xx response")
			}
			counter++
			return nil
		},
	}
	req, _ := http.NewRequestWithContext(httptrace.WithClientTrace(context.Background(), trace), http.MethodGet, ts.URL, nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, "hello", string(body))
}
