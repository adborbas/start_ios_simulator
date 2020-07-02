package iossimulator

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// Simctl ...
type Simctl struct {
}

// AvailableDevices ...
// func (sim Simctl) AvailableDevices() ([]string, error) {
// 	rawList, err := sim.run("list")
// 	if err != nil {
// 		return []string{}, err
// 	}

// 	return strings.Split(rawList, "\n"), nil
// }

type runtimesResponse struct {
	runtimes []Runtime `json:"runtimes"`
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

	for _, runtime := range response.runtimes {
		if runtime.version == version {
			return &runtime, nil
		}
	}

	return nil, fmt.Errorf("%s not found", version)
}

// type devicesResponse struct {
// 	devices []Device
// }

func (sim Simctl) run(args ...string) (string, error) {
	cmd := exec.Command("xcrun", append([]string{"simctl"}, args...)...)
	outBytes, err := cmd.CombinedOutput()
	return string(outBytes), err
}
