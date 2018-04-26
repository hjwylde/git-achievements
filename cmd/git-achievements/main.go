package main

import "fmt"
import "github.com/hjwylde/git-achievements/internal/app"
import "log"
import "os"
import "os/exec"
import "strings"

const unknownCommandTemplate = "Unknown command: %s\n\n"
const unknownOptionTemplate = "Unknown option: %s\n\n"
const usageTemplate = "Usage: git-achievements [-h|--help] [-v|--version] COMMAND\n"
const availableCommandsTemplate = "\nAvailable commands:\n  %s\n"

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		help()
		return
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		version()
		return
	}

	cmd := exec.Command("git", "rev-parse")
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		os.Exit(2)
	}

	subcommand, ok := app.CommandMap[os.Args[1]]
	if !ok {
		if strings.HasPrefix(os.Args[1], "-") {
			fmt.Printf(unknownOptionTemplate, os.Args[1])
		} else {
			fmt.Printf(unknownCommandTemplate, os.Args[1])
		}

		usage()
		os.Exit(2)
	}

	flagSet := subcommand.FlagSet
	flagSet.Parse(os.Args[2:])

	if err := subcommand.Run(flagSet); err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Printf(usageTemplate)
}

func help() {
	usage()

	var subcommands []string
	for subcommand := range app.CommandMap {
		subcommands = append(subcommands, subcommand)
	}

	fmt.Printf(availableCommandsTemplate, strings.Join(subcommands, "\n  "))
}

func version() {
	fmt.Printf("0.1.0\n")
}
