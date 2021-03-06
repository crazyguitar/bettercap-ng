package network

type Station struct {
	*Endpoint
	Frequency  int    `json:"frequency"`
	RSSI       int8   `json:"rssi"`
	Sent       uint64 `json:"sent"`
	Received   uint64 `json:"received"`
	Encryption string `json:"encryption"`
}

func NewStation(essid, bssid string, frequency int, rssi int8) *Station {
	return &Station{
		Endpoint:  NewEndpointNoResolve(MonitorModeAddress, bssid, essid, 0),
		Frequency: frequency,
		RSSI:      rssi,
	}
}

func (s Station) BSSID() string {
	return s.HwAddress
}

func (s *Station) ESSID() string {
	return s.Hostname
}
