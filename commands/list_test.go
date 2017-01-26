package commands

import (
	"gopkg.in/stretchr/testify.v1/suite"
	"testing"

	"github.com/ismacaulay/fiz/io"
	"github.com/ismacaulay/fiz/wizards"
)

type ListCommandTestSuite struct {
	suite.Suite

	Provider *wizards.MockProvider
	Printer  *io.MockPrinter

	Patient *ListCommand
}

func (td *ListCommandTestSuite) SetupTest() {
	td.Provider = wizards.NewMockProvider()
	td.Printer = io.NewMockPrinter()

	td.Patient = NewListCommand(td.Provider, td.Printer)
}

func (td *ListCommandTestSuite) TestPrintWizards() {
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

	header := "Available Wizards:"
	noneGroupWizards := `
    - hello
    - world`

	groupOne := `
    group one:`
	groupOneWizards := `
        - hello
        - group
        - one`

	groupTwo := `
    group two:`
	groupTwoWizards := `
        - a
        - b`

	newLines := "\n\n"

	td.Provider.On("AllAvailableWizards").Return(w, nil)
	td.Printer.On("Message", header)
	td.Printer.On("Message", noneGroupWizards)
	td.Printer.On("Message", groupOne)
	td.Printer.On("Message", groupOneWizards)
	td.Printer.On("Message", groupTwo)
	td.Printer.On("Message", groupTwoWizards)
	td.Printer.On("Message", newLines)

	td.Patient.Run()

	td.Printer.AssertExpectations(td.Suite.T())
}

func TestListCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ListCommandTestSuite))
}
