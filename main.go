// Package main is the entry point for the `miniproject` brute-force simulator.
// It demonstrates initializing the logger and running an attack using the
// `attack.Engine` with a configured `attack.AttackConfig`. Run the program
// with an argument of `info` or `debug` to set the logging level.
//
// Usage:
//
//	go run . info
//	go run . debug
package main

import (
	"fmt"
	"miniproject/attack"
	"miniproject/utils"
	"os"
)

func main() {
	// main initializes logging, builds an attack configuration, and runs the
	// attack engine. It expects a single command-line argument to set the
	// log level: `info` or `debug`.
	fmt.Println("initiating logger")

	if len(os.Args) < 2 {
		os.Exit(2)
	}

	switch os.Args[1] {
	case "info":
		utils.Init(utils.Info)
	case "debug":
		utils.Init(utils.Debug)
	default:
		fmt.Println("Arguments thappu")
		os.Exit(3)
	}

	utils.Infof("Application started")

	targetPassword := "Abb0"
	targetHash := utils.ConvertToHash(&targetPassword)

	config := attack.AttackConfig{
		Mask:        "A???",
		Charset:     "abcdefghijklmnopqrstuvwxyz0123456789@#$%^&*()-_=+[]{}|;:',.<>/?`~ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		WorkerCount: 20,
		BufferSize:  200,
	}

	stats := &attack.AttackStats{}

	engine := &attack.Engine{
		Config:     config,
		TargetHash: targetHash,
		Stats:      stats,
	}

	result := engine.Run()

	utils.Infof("PASSWORD FOUND: %s", result.Password)
	utils.Infof("Attempts: %d", result.Attempts)
	utils.Infof("Duration: %s", result.Duration)
}
