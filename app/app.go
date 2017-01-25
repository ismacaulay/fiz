package app

import (
	"github.com/ismacaulay/fiz/commands"
	"github.com/ismacaulay/fiz/wizards"
)

type Application struct {
	runner commands.Runner
}

func NewApp(external External) *Application {
	validator := wizards.NewWizardValidator(external.TemplateGenerator())
	processor := wizards.NewWizardProcessor(external.Input())
	generator := wizards.NewOutputGenerator(external.FileSystem())
	factory := wizards.NewWizardFactory(
		validator, processor, generator,
		external.FileSystem(), external.Input(), external.Printer())
	provider := wizards.NewWizardProvider(external.FileSystem(), external.DirectoryProvider())
	loader := wizards.NewWizardLoader(provider, factory, external.FileSystem())

	cmdFactory := commands.NewCmdFactory(provider, loader, external.Printer())
	runner := commands.NewCommandRunner(external.Printer(), cmdFactory)

	return &Application{runner}
}

func (app *Application) Run(args []string) {
	app.runner.Run(args)
}
