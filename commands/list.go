package commands

import (
	"github.com/ismacaulay/fiz/output"
	"github.com/ismacaulay/fiz/wizards"
)

type ListCommand struct {
	provider wizards.Provider
	printer  output.Printer
}

func NewListCommand(provider wizards.Provider, printer output.Printer) *ListCommand {
	return &ListCommand{provider, printer}
}

func (c *ListCommand) Run() error {
	wizards, err := c.provider.AllAvailableWizards()
	if err != nil {
		return err
	}

	msg := `Available Wizards:`

	for category, names := range wizards {
		wizardsString := ""
		for _, name := range names {
			wizardsString += `
		` + name
		}

		msg = msg + `
	` + category + ":" + wizardsString
	}
	c.printer.Message(msg)
	return nil
}
