package commands

import (
	"github.com/ismacaulay/fiz/output"
	"github.com/ismacaulay/fiz/wizards"
)

type Cmd int

const (
	List Cmd = iota
)

type Factory interface {
	Create(t Cmd) Command
}

type CmdFactory struct {
	provider wizards.Provider
	printer  output.Printer
}

func NewCmdFactory(provider wizards.Provider, printer output.Printer) *CmdFactory {
	return &CmdFactory{provider, printer}
}

func (f *CmdFactory) Create(t Cmd) Command {
	switch t {
	case List:
		return NewListCommand(f.provider, f.printer)
	}
	return nil
}
