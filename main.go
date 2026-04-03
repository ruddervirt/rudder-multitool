package main

import (
	"fmt"
	"os"
	"strings"

	bcryptcheck "rudder-multitool/universal/bcrypt_check"
	changetime "rudder-multitool/universal/change_time"
)

const version = "0.1.0"

type commandFunc func(args []string) int

type command struct {
	name    string
	usage   string
	handler commandFunc
}

var commandRegistry = map[string]command{}

func registerCommand(name string, usage string, handler commandFunc) {
	commandRegistry[name] = command{name: name, usage: usage, handler: handler}
}

func registerBaseCommands() {
	registerCommand("version", "version", func(args []string) int {
		fmt.Println(version)
		return 0
	})

	registerCommand("bcrypt-check", "bcrypt-check <bcrypt_hash> <plaintext_password>", func(args []string) int {
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s bcrypt-check <bcrypt_hash> <plaintext_password>\n", os.Args[0])
			return 1
		}
		hash := args[0]
		password := args[1]
		if err := bcryptcheck.CheckBcrypt(hash, password); err != nil {
			fmt.Println("NO_MATCH")
			return 1
		}
		fmt.Println("MATCH")
		return 0
	})

	registerCommand("change-time", "change-time [--local|--utc] <ISO8601 timestamp>", func(args []string) int {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "Usage: %s change-time [--local|--utc] <ISO8601 timestamp>\n", os.Args[0])
			return 1
		}

		mode := changetime.ZoneAuto
		if strings.HasPrefix(args[0], "--") {
			if len(args) < 2 {
				fmt.Fprintf(os.Stderr, "Usage: %s change-time [--local|--utc] <ISO8601 timestamp>\n", os.Args[0])
				return 1
			}
			switch args[0] {
			case "--local":
				mode = changetime.ZoneLocal
			case "--utc":
				mode = changetime.ZoneUTC
			default:
				fmt.Fprintf(os.Stderr, "Usage: %s change-time [--local|--utc] <ISO8601 timestamp>\n", os.Args[0])
				return 1
			}
			args = args[1:]
		}

		if err := changetime.SetTimeWithZone(args[0], mode); err != nil {
			fmt.Println(err)
			return 1
		}
		fmt.Println("OK")
		return 0
	})
}

func main() {
	registerBaseCommands()
	registerPlatformCommands()

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	cmd := os.Args[1]
	entry, ok := commandRegistry[cmd]
	if !ok {
		os.Exit(1)
	}

	os.Exit(entry.handler(os.Args[2:]))
}
