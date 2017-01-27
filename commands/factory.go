package commands

import (
	"github.com/ismacaulay/fiz/utils"
	"github.com/ismacaulay/fiz/wizards"

	"gopkg.in/stretchr/testify.v1/mock"
)

type Factory interface {
	CreateListCmd() Command
	CreateWizardCmd(commands []string) Command
}

type CmdFactory struct {
	provider wizards.Provider
	loader   wizards.Loader
	printer  utils.Printer
}

func NewCmdFactory(provider wizards.Provider, loader wizards.Loader, printer utils.Printer) *CmdFactory {
	return &CmdFactory{provider, loader, printer}
}

func (f *CmdFactory) CreateListCmd() Command {
	return NewListCommand(f.provider, f.printer)
}

func (f *CmdFactory) CreateWizardCmd(commands []string) Command {
	return NewWizardCommand(f.loader, commands)
}

/************************************
 * Mock
 ************************************/
type MockFactory struct {
	mock.Mock
}

func NewMockFactory() *MockFactory {
	return &MockFactory{}
}

func (m *MockFactory) CreateListCmd() Command {
	args := m.Called()
	return args.Get(0).(Command)
}

func (m *MockFactory) CreateWizardCmd(cmds []string) Command {
	args := m.Called(cmds)
	return args.Get(0).(Command)
}
