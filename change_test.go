package advantageair_test

import (
	"encoding/json"
	"testing"

	advantageair "github.com/axatol/go-advantage-air"
	"github.com/stretchr/testify/assert"
)

func assertJSONEq(t *testing.T, expected, actual interface{}) {
	t.Helper()
	expectedJSON, err := json.Marshal(expected)
	assert.NoError(t, err)
	actualJSON, err := json.Marshal(actual)
	assert.NoError(t, err)
	assert.JSONEq(t, string(expectedJSON), string(actualJSON))
}

func TestChangeSetAirconState(t *testing.T) {
	expected := map[string]any{"aircons": map[string]any{"ac1": map[string]any{"info": map[string]any{"state": "on"}}}}
	actual := advantageair.NewChange().SetAirconState("ac1", "on")
	assertJSONEq(t, expected, actual)
}

func TestChangeSetAirconMode(t *testing.T) {
	expected := map[string]any{"aircons": map[string]any{"ac1": map[string]any{"info": map[string]any{"mode": "cool"}}}}
	actual := advantageair.NewChange().SetAirconMode("ac1", "cool")
	assertJSONEq(t, expected, actual)
}

func TestChangeSetAirconFan(t *testing.T) {
	expected := map[string]any{"aircons": map[string]any{"ac1": map[string]any{"info": map[string]any{"fan": "high"}}}}
	actual := advantageair.NewChange().SetAirconFan("ac1", "high")
	assertJSONEq(t, expected, actual)
}

func TestChangeSetAirconZoneTemperature(t *testing.T) {
	expected := map[string]any{"aircons": map[string]any{"ac1": map[string]any{"zones": map[string]any{"z1": map[string]any{"setTemp": 24}}}}}
	actual := advantageair.NewChange().SetAirconZoneTemperature("ac1", "z1", 24)
	assertJSONEq(t, expected, actual)
}

func TestChangeSetAirconZoneState(t *testing.T) {
	expected := map[string]any{"aircons": map[string]any{"ac1": map[string]any{"zones": map[string]any{"z1": map[string]any{"state": "open"}}}}}
	actual := advantageair.NewChange().SetAirconZoneState("ac1", "z1", "open")
	assertJSONEq(t, expected, actual)
}

func TestChangeComposite(t *testing.T) {
	expected := map[string]any{
		"aircons": map[string]any{
			"ac1": map[string]any{
				"info": map[string]any{
					"state": "on",
					"mode":  "cool",
					"fan":   "high",
				},
				"zones": map[string]any{
					"z1": map[string]any{
						"setTemp": 24,
						"state":   "open",
					},
				},
			},
		},
	}
	actual := advantageair.NewChange().
		SetAirconState("ac1", "on").
		SetAirconMode("ac1", "cool").
		SetAirconFan("ac1", "high").
		SetAirconZoneTemperature("ac1", "z1", 24).
		SetAirconZoneState("ac1", "z1", "open")
	assertJSONEq(t, expected, actual)
}
