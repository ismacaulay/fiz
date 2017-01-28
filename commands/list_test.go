package commands

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"gopkg.in/stretchr/testify.v1/suite"
	"testing"

	"bytes"
	"errors"
	"github.com/ismacaulay/fiz/utils"
	"github.com/ismacaulay/fiz/wizards"
)

type ListCommandTestSuite struct {
	suite.Suite

	Provider  *wizards.MockProvider
	Generator *utils.MockTemplateGenerator
	Printer   *utils.MockPrinter

	Patient *ListCommand
}

func (td *ListCommandTestSuite) SetupTest() {
	td.beforeEachCase()
}

func (td *ListCommandTestSuite) beforeEachCase() {
	td.Provider = wizards.NewMockProvider()
	td.Generator = utils.NewMockTemplateGenerator()
	td.Printer = utils.NewMockPrinter()

	td.Patient = NewListCommand(td.Provider, td.Generator, td.Printer)
}

func (td *ListCommandTestSuite) TestNoWizardsFound() {
	assert := assert.New(td.Suite.T())

	wizards := make(map[string][]wizards.WizardInfo)
	header := "Header"
	no_wizards := "No Wizards"
	footer := "Footer"

	td.Provider.On("AllAvailableWizards").Return(wizards, nil)

	td.Generator.On("Execute", HEADER_TMPL, nil).Return(bytes.NewBufferString(header), nil)
	td.Generator.On("Execute", NO_WIZARDS_TMPL, nil).Return(bytes.NewBufferString(no_wizards), nil)
	td.Generator.On("Execute", FOOTER_TMPL, nil).Return(bytes.NewBufferString(footer), nil)
	td.Printer.On("Message", header)
	td.Printer.On("Message", no_wizards)
	td.Printer.On("Message", footer)

	err := td.Patient.Run()

	assert.NoError(err)
	td.Generator.AssertExpectations(td.Suite.T())
	td.Printer.AssertExpectations(td.Suite.T())
}

func (td *ListCommandTestSuite) TestPrintWizards() {
	assert := assert.New(td.Suite.T())

	w := make(map[string][]wizards.WizardInfo)
	w["none"] = make([]wizards.WizardInfo, 2)
	w["none"][0] = wizards.WizardInfo{"none", "hello", ""}
	w["none"][1] = wizards.WizardInfo{"none", "world", ""}
	w["group one"] = make([]wizards.WizardInfo, 3)
	w["group one"][0] = wizards.WizardInfo{"group one", "hello", ""}
	w["group one"][1] = wizards.WizardInfo{"group one", "group", ""}
	w["group one"][2] = wizards.WizardInfo{"group one", "one", ""}
	w["group two"] = make([]wizards.WizardInfo, 2)
	w["group two"][0] = wizards.WizardInfo{"group two", "a", ""}
	w["group two"][1] = wizards.WizardInfo{"group two", "b", ""}

	header := "Header"
	none := "None"
	group1 := "Group 1"
	group1_wizards := "Group 1 Wizards"
	group2 := "Group 2"
	group2_wizards := "Group 2 Wizards"
	footer := "Footer"

	td.Provider.On("AllAvailableWizards").Return(w, nil)

	td.Generator.On("Execute", HEADER_TMPL, nil).Return(bytes.NewBufferString(header), nil)
	td.Generator.On("Execute", NONE_GROUP_TMPL, w["none"]).Return(bytes.NewBufferString(none), nil)
	td.Generator.On("Execute", GROUP_TMPL, "group one").Return(bytes.NewBufferString(group1), nil)
	td.Generator.On("Execute", GROUP_WIZARD_TMPL, w["group one"]).Return(bytes.NewBufferString(group1_wizards), nil)
	td.Generator.On("Execute", GROUP_TMPL, "group two").Return(bytes.NewBufferString(group2), nil)
	td.Generator.On("Execute", GROUP_WIZARD_TMPL, w["group two"]).Return(bytes.NewBufferString(group2_wizards), nil)
	td.Generator.On("Execute", FOOTER_TMPL, nil).Return(bytes.NewBufferString(footer), nil)
	td.Printer.On("Message", header)
	td.Printer.On("Message", none)
	td.Printer.On("Message", group1)
	td.Printer.On("Message", group1_wizards)
	td.Printer.On("Message", group2)
	td.Printer.On("Message", group2_wizards)
	td.Printer.On("Message", footer)

	err := td.Patient.Run()

	assert.NoError(err)
	td.Generator.AssertExpectations(td.Suite.T())
	td.Printer.AssertExpectations(td.Suite.T())
}

func (td *ListCommandTestSuite) TestErrorWhenExecuteErrorsOnNoneTemplate() {
	assert := assert.New(td.Suite.T())

	w := make(map[string][]wizards.WizardInfo)
	w["none"] = make([]wizards.WizardInfo, 2)

	expectedError := errors.New("This is an error")
	header := "Header"

	td.Provider.On("AllAvailableWizards").Return(w, nil)
	td.Generator.On("Execute", HEADER_TMPL, nil).Return(bytes.NewBufferString(header), nil)
	td.Generator.On("Execute", NONE_GROUP_TMPL, w["none"]).Return(bytes.NewBufferString(""), expectedError)
	td.Printer.On("Message", header)

	err := td.Patient.Run()

	assert.Error(err)
	assert.Equal(expectedError, err)
}

func (td *ListCommandTestSuite) TestErrorWhenExecuteErrorsOnGroupTemplate() {
	assert := assert.New(td.Suite.T())

	w := make(map[string][]wizards.WizardInfo)
	w["none"] = make([]wizards.WizardInfo, 2)
	w["none"][0] = wizards.WizardInfo{"none", "hello", ""}
	w["none"][1] = wizards.WizardInfo{"none", "world", ""}
	w["group one"] = make([]wizards.WizardInfo, 3)

	expectedError := errors.New("This is an error")
	header := "Header"
	none := "None"
	group1 := "Group 1"

	td.Provider.On("AllAvailableWizards").Return(w, nil)

	td.Generator.On("Execute", HEADER_TMPL, nil).Return(bytes.NewBufferString(header), nil)
	td.Generator.On("Execute", NONE_GROUP_TMPL, w["none"]).Return(bytes.NewBufferString(none), nil)
	td.Generator.On("Execute", GROUP_TMPL, "group one").Return(bytes.NewBufferString(group1), expectedError)
	td.Printer.On("Message", header)
	td.Printer.On("Message", none)
	td.Printer.On("Message", group1)

	err := td.Patient.Run()

	assert.Error(err)
	assert.Equal(expectedError, err)
}

func (td *ListCommandTestSuite) TestErrorWhenExecuteErrorsOnGroupWizardsTemplate() {
	assert := assert.New(td.Suite.T())

	w := make(map[string][]wizards.WizardInfo)
	w["none"] = make([]wizards.WizardInfo, 2)
	w["none"][0] = wizards.WizardInfo{"none", "hello", ""}
	w["none"][1] = wizards.WizardInfo{"none", "world", ""}
	w["group one"] = make([]wizards.WizardInfo, 3)

	expectedError := errors.New("This is an error")
	header := "Header"
	none := "None"
	group1 := "Group 1"

	td.Provider.On("AllAvailableWizards").Return(w, nil)

	td.Generator.On("Execute", HEADER_TMPL, nil).Return(bytes.NewBufferString(header), nil)
	td.Generator.On("Execute", NONE_GROUP_TMPL, w["none"]).Return(bytes.NewBufferString(none), nil)
	td.Generator.On("Execute", GROUP_TMPL, "group one").Return(bytes.NewBufferString(group1), nil)
	td.Generator.On("Execute", GROUP_WIZARD_TMPL, w["group one"]).Return(bytes.NewBufferString(""), expectedError)
	td.Printer.On("Message", header)
	td.Printer.On("Message", none)
	td.Printer.On("Message", group1)

	err := td.Patient.Run()

	assert.Error(err)
	assert.Equal(expectedError, err)
}

func TestListCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ListCommandTestSuite))
}
