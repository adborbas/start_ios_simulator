package iossimulator

import "fmt"

// RuntimeNotFoundError ...
type RuntimeNotFoundError struct {
	Version string
}

func (e *RuntimeNotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Version)
}

// DeviceNotFoundError ...
type DeviceNotFoundError struct {
	Name string
}

func (e *DeviceNotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Name)
}
