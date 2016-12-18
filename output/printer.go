package output

import (
	"fmt"
)

type Printer interface {
	Help()
}

type TextPrinter struct {
}

func NewTextPrinter() *TextPrinter {
	return &TextPrinter{}
}

func (p *TextPrinter) Help() {
	helpText := `Name: 
	fiz - a file wizard

Usage: 
	fiz <command>

Commands:
	list	list all wizards
	help	print this help message
`

	fmt.Println(helpText)
}
