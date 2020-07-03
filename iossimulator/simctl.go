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

// AvailableRuntimes ...
func (sim Simctl) AvailableRuntimes() ([]Runtime, error) {
	return sim.runtimes("iOS")
}

// Runtime ...
func (sim Simctl) Runtime(version string) (*Runtime, error) {
	runtimes, err := sim.runtimes(version)
	if err != nil {
		return nil, err
	}

	for _, runtime := range runtimes {
		if runtime.Version == version {
			return &runtime, nil
		}
	}

	return nil, &RuntimeNotFoundError{version}
}

// AvailableDevices ...
func (sim Simctl) AvailableDevices() ([]Device, error) {
	devices, err := sim.devices("iPhone")
	if err != nil {
		return nil, err
	}

	return devices, err
}

// Device ...
func (sim Simctl) Device(name string, runtime Runtime) (*Device, error) {

	devices, err := sim.devices(name)
	if err != nil {
		return nil, err
	}

	for _, device := range devices {
		if device.Runtime.Identifier == runtime.Identifier && device.Name == name {
			return &device, nil
		}
	}

	return nil, &DeviceNotFoundError{name}
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
	runtime, err := sim.Runtime(parameters.Runtime)
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

type runtimesResponse struct {
	Runtimes []Runtime `json:"runtimes"`
}

func (sim Simctl) runtimes(version string) ([]Runtime, error) {
	args := []string{"list", "--json", "runtimes"}
	if len(version) != 0 {
		args = append(args, version)
	}
	rawJSON, err := sim.run(args...)
	if err != nil {
		return nil, err
	}

	var response runtimesResponse
	if err = json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}

	return response.Runtimes, nil
}

type devicesResponse struct {
	Devices map[string][]Device `json:"devices"`
}

func (sim Simctl) devices(name string) ([]Device, error) {
	args := []string{"list", "--json", "devices"}
	if len(name) != 0 {
		args = append(args, name)
	}
	rawJSONString, err := sim.run(args...)
	if err != nil {
		return nil, err
	}

	var response devicesResponse
	if err = json.Unmarshal([]byte(rawJSONString), &response); err != nil {
		return nil, err
	}

	runtimes, err := sim.runtimes("iOS")
	if err != nil {
		return nil, err
	}

	runtimeDevices := []Device{}
	for runtimeIdentifier, devices := range response.Devices {
		devices := devices
		if strings.Contains(runtimeIdentifier, "iOS") {
			for index := range devices {
				if rnt := sim.findRuntime(runtimes, runtimeIdentifier); rnt != nil {
					devices[index].Runtime = *rnt
				}
			}
			runtimeDevices = append(runtimeDevices, devices...)
			break
		}
	}

	return runtimeDevices, nil
}

func (sim Simctl) findRuntime(runtimes []Runtime, identifier string) *Runtime {
	for _, runtime := range runtimes {
		if runtime.Identifier == identifier {
			return &runtime
		}
	}

	return nil
}
