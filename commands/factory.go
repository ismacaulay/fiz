package commands

import (
	"github.com/ismacaulay/fiz/utils"
	"github.com/ismacaulay/fiz/wizards"

	"gopkg.in/stretchr/testify.v1/mock"
)

type Factory interface {
	CreateListCmd() Command
	CreateWizardCmd(commands []string) Command
	CreateConfigCmd() Command
}

type CmdFactory struct {
	provider  wizards.Provider
	loader    wizards.Loader
	directory utils.DirectoryProvider
	generator utils.TemplateGenerator
	printer   utils.Printer
}

func NewCmdFactory(provider wizards.Provider, loader wizards.Loader,
	directory utils.DirectoryProvider, generator utils.TemplateGenerator, printer utils.Printer) *CmdFactory {
	return &CmdFactory{provider, loader, directory, generator, printer}
}

func (f *CmdFactory) CreateListCmd() Command {
	return NewListCommand(f.provider, f.generator, f.printer)
}

func (f *CmdFactory) CreateWizardCmd(commands []string) Command {
	return NewWizardCommand(f.loader, commands)
}

func (f *CmdFactory) CreateConfigCmd() Command {
	return NewConfigCommand(f.directory, f.generator, f.printer)
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

func (m *MockFactory) CreateConfigCmd() Command {
	args := m.Called()
	return args.Get(0).(Command)
}
