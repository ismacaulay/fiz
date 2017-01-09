package app

import (
	"github.com/ismacaulay/fiz/commands"
)

type Application struct {
	runner  commands.Runner
}

func NewApp(runner commands.Runner) *Application {
	return &Application{runner}
}

func (app *Application) Run(args []string) {
	app.runner.Run(args)
}
