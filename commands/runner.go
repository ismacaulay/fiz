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
	factory Factory
}

func NewCommandRunner(printer output.Printer, factory Factory) *CommandRunner {
	return &CommandRunner{printer, factory}
}

func (r *CommandRunner) Run(command string) {
	switch command {
	case "list", "-l":
		cmd := r.factory.Create(List)
		cmd.Run()
	case "help", "-h", "--help":
		r.printer.Help()
	case "version", "--version":
		r.printer.Version()
	default:
		r.printer.Message(fmt.Sprint("Invalid command", command))
		r.printer.Help()
	}
}
