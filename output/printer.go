package output

import (
	"fmt"
)

type Printer interface {
	Help()
	Version()
	Message(msg string)
	Error(err error)
	Commands()
}

type TextPrinter struct {
	version string
}

func NewTextPrinter(version string) *TextPrinter {
	return &TextPrinter{version}
}

func (p *TextPrinter) Help() {
	helpText := `Name:
	fiz - a file wizard

Version:
	` + p.version + `

Usage:
	fiz <command>

`
	p.Message(helpText)
	p.Commands()
}

func (p *TextPrinter) Version() {
	fmt.Println("fiz", p.version)
}

func (p *TextPrinter) Message(msg string) {
	fmt.Print(msg)
}

func (p *TextPrinter) Error(err error) {
	fmt.Print(err)
}

func (p *TextPrinter) Commands() {
	commands := `Commands:
	list, -l		list all wizards
	<wizard>		run wizard
	<group> <wizard>	run wizard in group
	version, --version	print the version
	help, -h, --help	print this help message
`
	p.Message(commands)
}
