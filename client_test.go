package advantageair_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	advantageair "github.com/axatol/go-advantage-air"
	"github.com/stretchr/testify/assert"
)

func mockServer(t *testing.T) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	mux.Handle("/getSystemData", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, err := os.ReadFile("mockdata/getSystemData.json")
		if err != nil {
			t.Fatalf("failed to read mock data: %s", err)
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(raw)
		assert.NoError(t, err)
	}))
	mux.Handle("/setAircon", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, err := os.ReadFile("mockdata/setAircon.json")
		if err != nil {
			t.Fatalf("failed to read mock data: %s", err)
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(raw)
		assert.NoError(t, err)
	}))
	mux.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Advantage Air v15.1378"))
		assert.NoError(t, err)
	}))
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)
	return server
}

func TestGetSystemData(t *testing.T) {
	server := mockServer(t)
	client := advantageair.NewClient(server.URL)
	ctx := context.Background()
	_, err := client.GetSystemData(ctx)
	assert.NoError(t, err)
}

func TestSetAircon(t *testing.T) {
	server := mockServer(t)
	client := advantageair.NewClient(server.URL)
	ctx := context.Background()
	err := client.SetAircon(ctx, advantageair.NewChange())
	assert.NoError(t, err)
}
