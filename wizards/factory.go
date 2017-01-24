package wizards

import (
	"github.com/ismacaulay/fiz/io"
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
	input   io.Input
	printer io.Printer
}

func NewWizardFactory(
	v Validator, p Processor, g Generator,
	fs utils.FileSystem, input io.Input, printer io.Printer) *WizardFactory {
	return &WizardFactory{v, p, g, fs, input, printer}
}

func (f *WizardFactory) Create(info WizardInfo, outdir string) Wizard {
	return NewWizard(info, f.v, f.p, f.g, f.fs, f.input, f.printer, outdir)
}
