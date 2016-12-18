package app

import (
	"github.com/ismacaulay/fiz/commands"
	"github.com/ismacaulay/fiz/output"
)

type Application struct {
	runner  commands.Runner
	printer output.Printer
}

func NewApp(runner commands.Runner, printer output.Printer) *Application {
	return &Application{runner, printer}
}

func (app *Application) Run(args []string) {
	switch len(args) {
	case 1:
		app.runner.Run(args[0])
	default:
		app.printer.Help()
	}
}
