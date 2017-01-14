package wizards

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/ismacaulay/fiz/io"
	"github.com/ismacaulay/fiz/utils"
)

type Wizard interface {
	Run() error
}

type RealWizard struct {
	info    WizardInfo
	fs      utils.FileSystem
	input   io.Input
	printer io.Printer
}

func NewWizard(info WizardInfo, fs utils.FileSystem, input io.Input, printer io.Printer) *RealWizard {
	return &RealWizard{info, fs, input, printer}
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

	if err = w.validateWizardJson(wizardJson); err != nil {
		return err
	}

	_, err = w.getVariables(wizardJson)
	if err != nil {
		return err
	}

	w.printer.Message("\nGenerating files:\n")
	w.printer.Message("todo: output files, and data\n")

	generate, err := w.input.GetBoolean("Are you sure?")
	if err != nil {
		return err
	}

	if generate {
		w.printer.Message("Generating.\n")
	}
	return nil
}

func (w *RealWizard) validateWizardJson(data WizardJson) error {
	basepath, _ := filepath.Split(w.info.Path)

	for _, t := range data.Templates {
		if err := w.validateTemplateFile(t.Name, basepath); err != nil {
			return err
		}
		if err := w.validateCondition(t.Condition, data.Variables); err != nil {
			return err
		}
	}

	for _, v := range data.Variables {
		if err := w.validateCondition(v.Condition, data.Variables); err != nil {
			return err
		}
	}
	return nil
}

func (w *RealWizard) validateTemplateFile(name, basepath string) error {
	templatePath := filepath.Clean(filepath.Join(basepath, name))
	if _, err := template.ParseFiles(templatePath); err != nil {
		return err
	}

	return nil
}

func (w *RealWizard) validateCondition(condition []string, variables []VariableJson) error {
	if len(condition) == 0 {
		return nil
	} else if len(condition)%2 == 0 {
		return errors.New("Not enough condition elements")
	}

	for index, element := range condition {
		if index%2 == 0 {
			for _, variable := range variables {
				if variable.Name == element {
					switch variable.Type {
					case "bool":
						continue
					default:
						return errors.New("Invalid condition")
					}

				}
			}
		} else {
			switch element {
			case "&&", "||":
				continue
			default:
				return errors.New("Invalid element")
			}
		}
	}

	return nil
}

func (w *RealWizard) getVariables(data WizardJson) (map[string]interface{}, error) {
	vars := make(map[string]interface{})
	var err error

	for _, v := range data.Variables {
		if len(v.Condition) == 0 {
			vars, err = w.getVariable(v, vars)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, v := range data.Variables {
		if len(v.Condition) > 0 {
			condition := w.evaluateCondition(v.Condition, vars)

			if condition {
				vars, err = w.getVariable(v, vars)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return vars, nil
}

func (w *RealWizard) getVariable(v VariableJson, vars map[string]interface{}) (map[string]interface{}, error) {
	switch v.Type {
	case "bool":
		value, err := w.input.GetBoolean(v.Name)
		if err != nil {
			return vars, err
		}
		vars[v.Name] = value
	default:
		value, err := w.input.GetString(v.Name)
		if err != nil {
			return vars, err
		}
		vars[v.Name] = value
	}

	return vars, nil
}

func (w *RealWizard) evaluateCondition(conditions []string, vars map[string]interface{}) bool {
	condition := vars[conditions[0]].(bool)
	nextOperator := "&&"
	for _, c := range conditions[1:] {
		switch c {
		case "&&", "||":
			nextOperator = c
		default:
			switch nextOperator {
			case "&&":
				condition = condition && vars[c].(bool)
			case "||":
				condition = condition || vars[c].(bool)
			}
		}
	}
	return condition
}
