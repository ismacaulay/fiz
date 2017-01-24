package wizards

import (
	"encoding/json"
	"fmt"

	"github.com/ismacaulay/fiz/io"
	"github.com/ismacaulay/fiz/utils"
)

type Wizard interface {
	Run() error
}

type RealWizard struct {
	info      WizardInfo
	validator Validator
	processor Processor
	generator Generator
	fs        utils.FileSystem
	input     io.Input
	printer   io.Printer
	outdir    string
}

func NewWizard(
	info WizardInfo, v Validator, p Processor, g Generator,
	fs utils.FileSystem,
	input io.Input, printer io.Printer,
	outdir string) *RealWizard {
	return &RealWizard{info, v, p, g, fs, input, printer, outdir}
}

type TemplatePathPair struct {
	input  string
	output string
}

func (w *RealWizard) Run() error {
	w.printer.Message(fmt.Sprintf("Running wizard: %s\n\n", w.info.Path))

	data, err := w.fs.ReadFile(w.info.Path)
	if err != nil {
		return err
	}

	var wizardJson WizardJson
	if err = json.Unmarshal(data, &wizardJson); err != nil {
		return err
	}

	if err = w.validator.Validate(w.info, wizardJson); err != nil {
		return err
	}

	vars, templates, err := w.processor.Process(w.info, w.outdir, wizardJson)
	if err != nil {
		return err
	}

	w.printer.Message("\nGenerating files:\n")
	for _, t := range templates {
		w.printer.Message(fmt.Sprintf("\t %s\n", t.output))
	}

	generate, err := w.input.GetBoolean("Are you sure?")
	if err != nil {
		return err
	}

	if generate {
		w.printer.Message("Generating.\n")
		if err := w.generator.Generate(templates, vars); err != nil {
			return err
		}
	}
	return nil
}
