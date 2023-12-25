package advantageair

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/axatol/go-advantage-air/internal/util"
)

// Client is the interface for interacting with the Advantage Air hub.
type Client interface {
	// GetSystemInfo returns the system info the hub responds with by default.
	GetSystemInfo(ctx context.Context) (string, error)
	// GetSystemData returns the current state of the system.
	GetSystemData(ctx context.Context) (*SystemData, error)
	// SetAircon sets the state of the aircon.
	SetAircon(ctx context.Context, change Change) error
}

// NewClient returns a new Client.
//
// The address should be the address of the Advantage Air hub, including protocol and port,
// e.g. http://192.168.1.2:2025.
func NewClient(address string, retries int) Client {
	if retries < 1 {
		retries = 0
	}

	rt := util.NewRetryRoundTripper(http.DefaultTransport, retries, func(res *http.Response) bool {
		return res.StatusCode == http.StatusOK
	})

	return &client{address, http.Client{Transport: rt}}
}

type client struct {
	address string
	client  http.Client
}

func (c *client) GetSystemInfo(ctx context.Context) (string, error) {
	uri := c.address + "/"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create system info request: %s", err)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get system info: %s", err)
	}

	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read system info: %s", err)
	}

	return string(raw), nil
}

func (c *client) GetSystemData(ctx context.Context) (*SystemData, error) {
	uri := c.address + "/getSystemData"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create system data request: %s", err)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get system data: %s", err)
	}

	defer res.Body.Close()
	var data SystemData
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode system data: %s", err)
	}

	return &data, nil
}

func (c *client) SetAircon(ctx context.Context, change Change) error {
	raw, err := json.Marshal(change)
	if err != nil {
		return fmt.Errorf("failed to marshal change: %s", err)
	}

	query := url.Values{"json": {string(raw)}}
	uri := c.address + "/setAircon?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return fmt.Errorf("failed to create set aircon request: %s", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to set aircon: %s", err)
	}

	defer resp.Body.Close()
	var data SetResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode set aircon response: %s", err)
	}

	if !data.Acknowledged {
		return fmt.Errorf("set aircon change was not acknowledged")
	}

	return nil
}
