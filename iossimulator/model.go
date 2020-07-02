package iossimulator

// Runtime ...
type Runtime struct {
	name       string `json:"name"`
	version    string `json:"version"`
	identifier string `json:"identifier"`
}

type Device struct {
	name       string `json:"name"`
	identifier string `json:"udid"`
}
