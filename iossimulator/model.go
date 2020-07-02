package iossimulator

// Runtime ...
type Runtime struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Identifier string `json:"identifier"`
}

// Device ...
type Device struct {
	State      string `json:"state"`
	Name       string `json:"name"`
	Identifier string `json:"udid"`
}
