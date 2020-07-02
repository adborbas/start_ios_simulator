package main

import (
	"fmt"
	"os"

	"github.com/adborbas/start_ios_simulator/iossimulator"
)

func main() {
	// args := os.Args[1:]
	// name := args[0]
	// fmt.Printf("Helloka %s \n", name)

	simctl := iossimulator.Simctl{}
	runtime, err := simctl.Runtime("13.5")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(runtime)
}
