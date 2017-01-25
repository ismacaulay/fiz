package commands_test

import (
	"errors"
	"gopkg.in/stretchr/testify.v1/assert"
	"gopkg.in/stretchr/testify.v1/suite"
	"testing"

	"github.com/ismacaulay/fiz/commands"
	"github.com/ismacaulay/fiz/wizards"
)

type WizardCommandTestSuite struct {
	suite.Suite

	Loader *wizards.MockLoader

	Patient *commands.WizardCommand
}

func (td *WizardCommandTestSuite) SetupTest() {
	td.Loader = wizards.NewMockLoader()
}

func (td *WizardCommandTestSuite) TestRunReturnsErrorWhenLoadingFails() {
	assert := assert.New(td.Suite.T())

	args := make([]string, 0)
	td.Patient = commands.NewWizardCommand(td.Loader, args)

	td.Loader.On("Load", args).Return(wizards.NewMockWizard(), errors.New("Error"))

	err := td.Patient.Run()
	assert.Error(err)

	td.Loader.AssertExpectations(td.Suite.T())
}

func (td *WizardCommandTestSuite) TestRunWizardWhenLoadingSuccessful() {
	assert := assert.New(td.Suite.T())

	args := make([]string, 0)
	td.Patient = commands.NewWizardCommand(td.Loader, args)

	wizard := wizards.NewMockWizard()
	wizard.On("Run").Return(nil)
	td.Loader.On("Load", args).Return(wizard, nil)

	assert.Nil(td.Patient.Run())

	td.Loader.AssertExpectations(td.Suite.T())
	wizard.AssertExpectations(td.Suite.T())
}

func TestWizardCommandTestSuite(t *testing.T) {
	suite.Run(t, new(WizardCommandTestSuite))
}
