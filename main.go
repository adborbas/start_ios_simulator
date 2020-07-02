package main

import (
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
		fmt.Printf("Failed to start simulator %s", err)
		os.Exit(1)
	}

	os.Exit(0)
}
