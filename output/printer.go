package output

import (
	"fmt"
)

type Printer interface {
	Help()
	Version()
	Message(msg string)
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

Commands:
	list, -l			list all wizards
	version, --version	print the version
	help, -h, --help	print this help message
`

	fmt.Println(helpText)
}

func (p *TextPrinter) Version() {
	fmt.Println("fiz", p.version)
}

func (p *TextPrinter) Message(msg string) {
	fmt.Println(msg)
}
