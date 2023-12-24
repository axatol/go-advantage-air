package advantageair

type AirconState string

const (
	AirconStateOpen  AirconState = "on"
	AirconStateClose AirconState = "off"
)

type AirconMode string

const (
	AirconModeCool AirconMode = "cool"
	AirconModeHeat AirconMode = "heat"
	AirconModeFan  AirconMode = "fan"
	AirconModeDry  AirconMode = "dry"
)

type AirconFanSpeed string

const (
	AirconFanSpeedLow    AirconFanSpeed = "low"
	AirconFanSpeedMedium AirconFanSpeed = "medium"
	AirconFanSpeedHigh   AirconFanSpeed = "high"
	AirconFanSpeedAuto   AirconFanSpeed = "auto"
)

type AirconZoneState string

const (
	AirconZoneStateOpen  AirconZoneState = "open"
	AirconZoneStateClose AirconZoneState = "close"
)

type SetResponse struct {
	Acknowledged bool   `json:"ack"`
	Request      string `json:"request"`
}

type SystemData struct {
	Aircons map[string]Aircon `json:"aircons"`
	System  System            `json:"system"`
}

func (s *SystemData) GetAirconByID(id string) *Aircon {
	if s == nil {
		return nil
	}

	if aircon, ok := s.Aircons[id]; ok {
		return &aircon
	}

	return nil
}

func (s *SystemData) GetAirconByName(name string) *Aircon {
	if s == nil {
		return nil
	}

	for _, aircon := range s.Aircons {
		if aircon.Info.Name == name {
			return &aircon
		}
	}

	return nil
}

type System struct {
	ServiceVersion string            `json:"aaServiceRev"`
	DeviceNames    map[string]string `json:"deviceNames"`
	AppRevision    string            `json:"myAppRev"`
	Name           string            `json:"name"`
	NeedsUpdate    bool              `json:"needsUpdate"`
	TSPErrorCode   string            `json:"tspErrorCode"`
	TSPIP          string            `json:"tspIp"`
	TSPModel       string            `json:"tspModel"`
}

type Aircon struct {
	Info  AirconInfo            `json:"info"`
	Zones map[string]AirconZone `json:"zones"`
}

func (a *Aircon) GetZoneByID(id string) *AirconZone {
	if a == nil {
		return nil
	}

	if zone, ok := a.Zones[id]; ok {
		return &zone
	}

	return nil
}

func (a *Aircon) GetZoneByNumber(number int64) *AirconZone {
	if a == nil {
		return nil
	}

	for _, zone := range a.Zones {
		if zone.Number == number {
			return &zone
		}
	}

	return nil
}

func (a *Aircon) GetZoneByName(name string) *AirconZone {
	if a == nil {
		return nil
	}

	for _, zone := range a.Zones {
		if zone.Name == name {
			return &zone
		}
	}

	return nil
}

type AirconInfo struct {
	Error   string `json:"airconErrorCode"`
	Fan     string `json:"fan"`
	Mode    string `json:"mode"`
	MyZone  int64  `json:"myZone"`
	Name    string `json:"name"`
	SetTemp int64  `json:"setTemp"`
	State   string `json:"state"`
}

type AirconZone struct {
	Error               int64   `json:"error"`
	MeasuredTemperature float64 `json:"measuredTemp"`
	Name                string  `json:"name"`
	Number              int64   `json:"number"`
	SensorUID           string  `json:"SensorUid"`
	SetTemperature      int64   `json:"setTemp"`
	State               string  `json:"state"`
}
