package advantageair

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client interface {
	GetSystemData(ctx context.Context) (*SystemData, error)
	SetAircon(ctx context.Context, change Change) error
}

func NewClient(address string) Client { return &client{address} }

type client struct{ address string }

func (c *client) GetSystemData(ctx context.Context) (*SystemData, error) {
	uri := c.address + "/getSystemData"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create system data request: %s", err)
	}

	res, err := http.DefaultClient.Do(req)
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

	resp, err := http.DefaultClient.Do(req)
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
