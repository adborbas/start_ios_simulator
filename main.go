package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/adborbas/start_ios_simulator/iossimulator"
)

func main() {
	args := os.Args[1:]

	parameters, err := iossimulator.ParseParameters(args)
	if err != nil {
		fmt.Printf("Could not parse arguments %s \n", err)
		os.Exit(1)
	}

	simctl := iossimulator.Simctl{}
	if err := simctl.StartSimulator(*parameters); err != nil {

		var runtimeerr *iossimulator.RuntimeNotFoundError
		if errors.As(err, &runtimeerr) {
			fmt.Printf("Invalid runtime %s \n", runtimeerr.Version)
			printAvailableRuntimes(simctl)
			os.Exit(0)
		}

		var deviceerr *iossimulator.DeviceNotFoundError
		if errors.As(err, &deviceerr) {
			fmt.Printf("Invalid device %s \n", deviceerr.Name)
			printAvailableDevices(simctl)
			os.Exit(0)
		}

		fmt.Printf("Failed to start simulator %s", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func printAvailableRuntimes(simctl iossimulator.Simctl) {
	runtimes, err := simctl.AvailableRuntimes()
	if err != nil {
		return
	}

	fmt.Println()
	fmt.Println("Available runtimes: ")
	for _, runtime := range runtimes {
		fmt.Println(runtime.Version)
	}
}

func printAvailableDevices(simctl iossimulator.Simctl) {
	devices, err := simctl.AvailableDevices()
	if err != nil {
		return
	}

	fmt.Println()
	fmt.Println("Available devices: ")
	for _, device := range devices {
		fmt.Printf("%s, %s \n", device.Runtime.Version, device.Name)
	}
}
