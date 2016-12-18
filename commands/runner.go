package commands

import (
	"fmt"

	"github.com/ismacaulay/fiz/output"
)

type Runner interface {
	Run(command string)
}

type CommandRunner struct {
	printer output.Printer
}

func NewCommandRunner(printer output.Printer) *CommandRunner {
	return &CommandRunner{printer}
}

func (r *CommandRunner) Run(command string) {
	switch command {
	case "list", "-l":
		fmt.Println("running", command)
	case "help", "-h", "--help":
		r.printer.Help()
	case "version", "--version":
		r.printer.Version()
	default:
		r.printer.Message(fmt.Sprint("Invalid command", command))
		r.printer.Help()
	}
}
