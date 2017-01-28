package commands

import (
	"github.com/ismacaulay/fiz/utils"
)

type Runner interface {
	Run(commands []string)
}

type CommandRunner struct {
	printer utils.Printer
	factory Factory
}

func NewCommandRunner(printer utils.Printer, factory Factory) *CommandRunner {
	return &CommandRunner{printer, factory}
}

func (r *CommandRunner) Run(commands []string) {
	if len(commands) == 0 {
		r.printer.Help()
		return
	}

	command := commands[0]
	switch command {
	case "list", "-l":
		cmd := r.factory.CreateListCmd()
		cmd.Run()
	case "config", "--config":
		cmd := r.factory.CreateConfigCmd()
		cmd.Run()
	case "help", "-h", "--help":
		r.printer.Help()
	case "version", "--version":
		r.printer.Version()
	default:
		cmd := r.factory.CreateWizardCmd(commands)
		if err := cmd.Run(); err != nil {
			r.printer.Error(err)
		}
	}
}
