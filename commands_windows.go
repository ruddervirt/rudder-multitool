//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"

	startupdisabled "rudder-multitool/win/startup_disabled"
)

func registerPlatformCommands() {
	registerCommand("startup-disabled", "startup-disabled <StartupEntryName>", func(args []string) int {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "Usage: %s startup-disabled <StartupEntryName>\n", os.Args[0])
			return 1
		}
		name := args[0]
		disabled, err := startupdisabled.IsStartupDisabled(name)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		if disabled {
			fmt.Println("DISABLED")
			return 0
		}
		fmt.Println("ENABLED")
		return 0
	})
}
