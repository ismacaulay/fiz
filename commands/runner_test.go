package commands_test

import (
	"errors"
	"gopkg.in/stretchr/testify.v1/suite"
	"testing"

	"github.com/ismacaulay/fiz/commands"
	"github.com/ismacaulay/fiz/commands/mocks"
	"github.com/ismacaulay/fiz/io/mocks"
)

type CommandRunnerTestSuite struct {
	suite.Suite

	Printer *io_mocks.MockPrinter
	Factory *commands_mocks.MockFactory

	Patient *commands.CommandRunner
}

func (td *CommandRunnerTestSuite) SetupTest() {
	td.Printer = io_mocks.NewMockPrinter()
	td.Factory = commands_mocks.NewMockFactory()

	td.Patient = commands.NewCommandRunner(td.Printer, td.Factory)
}

func (td *CommandRunnerTestSuite) TestPrintHelpIfCommandsEmpty() {
	td.Printer.On("Help")

	cmds := []string{}
	td.Patient.Run(cmds)

	td.Printer.AssertExpectations(td.Suite.T())
}

func (td *CommandRunnerTestSuite) TestCreateAndRunListCommand() {
	var cases = []struct {
		arg string
	}{
		{"list"},
		{"-l"},
	}

	command := commands_mocks.NewMockCommand()
	td.Factory.On("CreateListCmd").Return(command).Twice()
	command.On("Run").Return(nil).Twice()

	for _, c := range cases {
		args := make([]string, 1)
		args[0] = c.arg

		td.Patient.Run(args)
	}

	td.Factory.AssertExpectations(td.Suite.T())
	command.AssertExpectations(td.Suite.T())

}

func (td *CommandRunnerTestSuite) TestHelpCommandPrintsHelp() {
	var cases = []struct {
		arg string
	}{
		{"help"},
		{"-h"},
		{"--help"},
	}

	td.Printer.On("Help").Times(3)

	for _, c := range cases {
		args := make([]string, 1)
		args[0] = c.arg

		td.Patient.Run(args)
	}

	td.Printer.AssertExpectations(td.Suite.T())

}

func (td *CommandRunnerTestSuite) TestVersionCommandPrintsVersion() {
	var cases = []struct {
		arg string
	}{
		{"version"},
		{"--version"},
	}

	td.Printer.On("Version").Twice()

	for _, c := range cases {
		args := make([]string, 1)
		args[0] = c.arg

		td.Patient.Run(args)
	}

	td.Printer.AssertExpectations(td.Suite.T())
}

func (td *CommandRunnerTestSuite) TestCreateAndRunWizardCommand() {
	var cases = []struct {
		args []string
	}{
		{[]string{"hello", "world"}},
		{[]string{"world", "-t", "--hello"}},
	}

	command := commands_mocks.NewMockCommand()
	command.On("Run").Return(nil).Twice()

	for _, c := range cases {
		td.Factory.On("CreateWizardCmd", c.args).Return(command)

		td.Patient.Run(c.args)

		td.Factory.AssertExpectations(td.Suite.T())
	}

	command.AssertExpectations(td.Suite.T())
}

func (td *CommandRunnerTestSuite) TestPrintErrorWhenWizardCommandReturnsError() {
	args := []string{"hello", "world"}
	err := errors.New("Error")
	command := commands_mocks.NewMockCommand()

	td.Factory.On("CreateWizardCmd", args).Return(command)
	command.On("Run").Return(err)
	td.Printer.On("Error", err)

	td.Patient.Run(args)

	td.Printer.AssertExpectations(td.Suite.T())
}

func TestCommandRunnerTestSuite(t *testing.T) {
	suite.Run(t, new(CommandRunnerTestSuite))
}
