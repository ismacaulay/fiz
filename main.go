package main

import (
	"os"

	"github.com/ismacaulay/fiz/app"
	"github.com/ismacaulay/fiz/commands"
	"github.com/ismacaulay/fiz/output"
)

func main() {
	runner := commands.NewCommandRunner()
	printer := output.NewTextPrinter()

	app := app.NewApp(runner, printer)
	app.Run(os.Args[1:])
}
