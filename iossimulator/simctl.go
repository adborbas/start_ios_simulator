package iossimulator

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// Simctl ...
type Simctl struct {
}

type runtimesResponse struct {
	Runtimes []Runtime `json:"runtimes"`
}

// Runtime ...
func (sim Simctl) Runtime(version string) (*Runtime, error) {
	rawJSONString, err := sim.run("list", "--json", "runtimes", version)
	if err != nil {
		return nil, err
	}

	var response runtimesResponse
	if err = json.Unmarshal([]byte(rawJSONString), &response); err != nil {
		return nil, err
	}

	for _, runtime := range response.Runtimes {
		if runtime.Version == version {
			return &runtime, nil
		}
	}

	return nil, fmt.Errorf("runtime %s not found", version)
}

type devicesResponse struct {
	Devices map[string][]Device `json:"devices"`
}

// Device ...
func (sim Simctl) Device(name string, runtime Runtime) (*Device, error) {
	rawJSONString, err := sim.run("list", "--json", "devices", name)
	if err != nil {
		return nil, err
	}

	var response devicesResponse
	if err = json.Unmarshal([]byte(rawJSONString), &response); err != nil {
		return nil, err
	}

	for runtimeIdentifier, devices := range response.Devices {
		if runtimeIdentifier != runtime.Identifier {
			continue
		}
		for _, device := range devices {
			if device.Name == name {
				return &device, nil
			}
		}
	}
	return nil, fmt.Errorf("device %s not found", name)
}

// Boot ...
func (sim Simctl) Boot(device Device) error {
	if strings.ToLower(device.State) != "booted" {
		_, err := sim.run("boot", device.Identifier)
		if err != nil {
			return err
		}
	}
	_, err := exec.Command("open", "-a", "Simulator").CombinedOutput()
	return err
}

// StartSimulator ...
func (sim Simctl) StartSimulator(parameters StartSimulatorParameters) error {
	runtime, err := sim.Runtime(parameters.Version)
	if err != nil {
		return err
	}

	foundDevice, err := sim.Device(parameters.Device, *runtime)
	if err != nil {
		return err
	}

	fmt.Printf("Starting simulator %s, %s \n", foundDevice.Name, runtime.Version)
	if err := sim.Boot(*foundDevice); err != nil {
		return fmt.Errorf("could not start simulator %s", err)
	}

	return nil
}

func (sim Simctl) run(args ...string) ([]byte, error) {
	cmd := exec.Command("xcrun", append([]string{"simctl"}, args...)...)
	return cmd.CombinedOutput()
}
