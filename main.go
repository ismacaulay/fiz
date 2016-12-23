package main

import (
	"os"

	"github.com/ismacaulay/fiz/app"
	"github.com/ismacaulay/fiz/commands"
	"github.com/ismacaulay/fiz/output"
	"github.com/ismacaulay/fiz/wizards"
)

func main() {
	provider := wizards.NewWizardProvider()
	printer := output.NewTextPrinter("0.0.1")
	cmdFactory := commands.NewCmdFactory(provider, printer)

	runner := commands.NewCommandRunner(printer, cmdFactory)

	app := app.NewApp(runner, printer)
	app.Run(os.Args[1:])
}
