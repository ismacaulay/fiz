package wizards

import (
	"github.com/ismacaulay/fiz/utils"
)

type Factory interface {
	Create(info WizardInfo, outdir string) Wizard
}

type WizardFactory struct {
	v       Validator
	p       Processor
	g       Generator
	fs      utils.FileSystem
	input   utils.Input
	printer utils.Printer
}

func NewWizardFactory(
	v Validator, p Processor, g Generator,
	fs utils.FileSystem, input utils.Input, printer utils.Printer) *WizardFactory {
	return &WizardFactory{v, p, g, fs, input, printer}
}

func (f *WizardFactory) Create(info WizardInfo, outdir string) Wizard {
	return NewWizard(info, f.v, f.p, f.g, f.fs, f.input, f.printer, outdir)
}
