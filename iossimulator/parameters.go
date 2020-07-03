package iossimulator

import "fmt"

// StartSimulatorParameters ...
type StartSimulatorParameters struct {
	Runtime, Device string
}

// ParseParameters ...
func ParseParameters(args []string) (*StartSimulatorParameters, error) {
	// return &StartSimulatorParameters{Runtime: "13.5",
	// Device: "iPhone 81"}, nil

	if len(args) != 4 {
		return nil, fmt.Errorf("required 4 arguments but found %d", len(args))
	}

	var parameters StartSimulatorParameters
	for i := 0; i < len(args); i += 2 {
		switch args[i] {
		case "-r":
			parameters.Runtime = args[i+1]
		case "-d":
			parameters.Device = args[i+1]
		default:
			return nil, fmt.Errorf("invalid flag %s", args[i])
		}
	}

	return &parameters, nil
}
