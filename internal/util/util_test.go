package util_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/axatol/go-advantage-air/internal/util"
	"github.com/stretchr/testify/assert"
)

func mockRetryServer(t *testing.T, retries int) (*httptest.Server, func() int) {
	t.Helper()
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if calls < retries-1 {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			assert.NoError(t, err)
		} else {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
			assert.NoError(t, err)
		}
		calls += 1
	}))
	t.Cleanup(server.Close)
	return server, func() int { return calls }
}

func TestSetRecursively(t *testing.T) {
	expected := map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": "value"}}}
	actual := util.SetRecursively(map[string]any{}, "value", "foo", "bar", "baz")
	assert.Equal(t, expected, actual)
}

func FuzzSetRecursively(f *testing.F) {
	f.Add("value,foo,bar,baz")
	f.Add("value,lorem,ipsum,dolor,amet")
	f.Add("value,a,b,c,d,e,f,g,h,i")
	f.Add("value,1,a,2,b,3,c")
	f.Fuzz(func(t *testing.T, input string) {
		keysAndValue := strings.Split(input, ",")
		if len(keysAndValue) < 2 {
			t.Skip() // we need at least one key and a value
		}
		value := keysAndValue[0]
		keys := keysAndValue[1:]

		expectedBuilder := strings.Builder{}
		for _, key := range keys {
			expectedBuilder.WriteString("map[")
			expectedBuilder.WriteString(key)
			expectedBuilder.WriteString(":")
		}
		expectedBuilder.WriteString(value)
		for range keys {
			expectedBuilder.WriteString("]")
		}

		expected := expectedBuilder.String()
		actual := fmt.Sprintf("%+v", util.SetRecursively(map[string]any{}, value, keys...))
		assert.Equal(t, expected, actual)
	})
}

func TestRetryRoundTripper(t *testing.T) {

	tests := []struct {
		name      string
		retries   int
		calls     int
		validator func(*http.Response) bool
	}{
		{"negative validated retry", -1, 1, util.DefaultRetryValidator},
		{"no validated retry", 0, 1, util.DefaultRetryValidator},
		{"one validated retry", 1, 1, util.DefaultRetryValidator},
		{"two validated retries", 2, 2, util.DefaultRetryValidator},
		{"three validated retries", 3, 3, util.DefaultRetryValidator},
		{"negative nonvalidated retry", -1, 1, nil},
		{"no nonvalidated retry", 0, 1, nil},
		{"one nonvalidated retry", 1, 1, nil},
		{"two nonvalidated retries", 2, 2, nil},
		{"three nonvalidated retries", 3, 3, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server, calls := mockRetryServer(t, test.retries)
			retryRoundTripper := util.NewRetryRoundTripper(http.DefaultTransport, test.retries, test.validator)
			client := http.Client{Transport: retryRoundTripper}
			req, err := http.NewRequest(http.MethodGet, server.URL, nil)
			assert.NoError(t, err)
			res, err := client.Do(req)
			assert.Equal(t, test.calls, calls())
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, res.StatusCode)
		})
	}
}
