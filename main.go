package main

import (
	"os"

	"github.com/ismacaulay/fiz/app"
	"github.com/ismacaulay/fiz/commands"
	"github.com/ismacaulay/fiz/output"
	"github.com/ismacaulay/fiz/utils"
	"github.com/ismacaulay/fiz/wizards"
)

const VERSION = "0.0.1"

func main() {
	filesystem := utils.NewFileSystem()
	directoryProvider := utils.NewDirectoryProvider()

	wizardFactory := wizards.NewWizardFactory(filesystem)
	wizardProvider := wizards.NewWizardProvider(filesystem, directoryProvider)
	wizardLoader := wizards.NewWizardLoader(wizardProvider, wizardFactory)

	printer := output.NewTextPrinter(VERSION)

	cmdFactory := commands.NewCmdFactory(wizardProvider, wizardLoader, printer)
	runner := commands.NewCommandRunner(printer, cmdFactory)

	app := app.NewApp(runner)
	app.Run(os.Args[1:])
}
