package commands

import (
	"github.com/ismacaulay/fiz/io"
	"github.com/ismacaulay/fiz/wizards"
)

type Factory interface {
	CreateListCmd() Command
	CreateWizardCmd(commands []string) Command
}

type CmdFactory struct {
	provider wizards.Provider
	loader   wizards.Loader
	printer  io.Printer
}

func NewCmdFactory(provider wizards.Provider, loader wizards.Loader, printer io.Printer) *CmdFactory {
	return &CmdFactory{provider, loader, printer}
}

func (f *CmdFactory) CreateListCmd() Command {
	return NewListCommand(f.provider, f.printer)
}

func (f *CmdFactory) CreateWizardCmd(commands []string) Command {
	return NewWizardCommand(f.loader, commands)
}
