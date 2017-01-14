package wizards

import (
	"github.com/ismacaulay/fiz/io"
	"github.com/ismacaulay/fiz/utils"
)

type Factory interface {
	Create(info WizardInfo) Wizard
}

type WizardFactory struct {
	fs      utils.FileSystem
	input   io.Input
	printer io.Printer
}

func NewWizardFactory(fs utils.FileSystem, input io.Input, printer io.Printer) *WizardFactory {
	return &WizardFactory{fs, input, printer}
}

func (f *WizardFactory) Create(info WizardInfo) Wizard {
	return NewWizard(info, f.fs, f.input, f.printer)
}
