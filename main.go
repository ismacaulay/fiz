package main

import (
	"os"

	"github.com/ismacaulay/fiz/app"
	"github.com/ismacaulay/fiz/commands"
	"github.com/ismacaulay/fiz/output"
)

func main() {
	printer := output.NewTextPrinter("0.0.1")

	runner := commands.NewCommandRunner(printer)

	app := app.NewApp(runner, printer)
	app.Run(os.Args[1:])
}
