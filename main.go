package main

import (
	"os"

	"github.com/ismacaulay/fiz/app"
)

const VERSION = "0.0.1"

func main() {
	external := app.NewExternal(VERSION)
	app := app.NewApp(external)
	app.Run(os.Args[1:])
}
