package app

import (
	"github.com/ismacaulay/fiz/commands"
	"github.com/ismacaulay/fiz/wizards"
)

type Application struct {
	runner commands.Runner
}

func NewApp(external External) *Application {
	wizardFactory := wizards.NewWizardFactory(external.FileSystem(), external.Input(), external.Printer())
	wizardProvider := wizards.NewWizardProvider(external.FileSystem(), external.DirectoryProvider())
	wizardLoader := wizards.NewWizardLoader(wizardProvider, wizardFactory, external.FileSystem())

	cmdFactory := commands.NewCmdFactory(wizardProvider, wizardLoader, external.Printer())
	runner := commands.NewCommandRunner(external.Printer(), cmdFactory)

	return &Application{runner}
}

func (app *Application) Run(args []string) {
	app.runner.Run(args)
}
