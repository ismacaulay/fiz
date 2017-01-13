package wizards

import (
	"github.com/ismacaulay/fiz/utils"
)

type Factory interface {
	Create(info WizardInfo) Wizard
}

type WizardFactory struct {
	fs utils.FileSystem
}

func NewWizardFactory(fs utils.FileSystem) *WizardFactory {
	return &WizardFactory{fs}
}

func (f *WizardFactory) Create(info WizardInfo) Wizard {
	return NewWizard(info, f.fs)
}
