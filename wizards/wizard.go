package wizards

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
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
	outdir  string
}

func NewWizard(info WizardInfo, fs utils.FileSystem, input io.Input, printer io.Printer, outdir string) *RealWizard {
	return &RealWizard{info, fs, input, printer, outdir}
}

type TemplatePathPair struct {
	input  string
	output string
}

func (w *RealWizard) Run() error {
	w.printer.Message(fmt.Sprintf("Running wizard: %s\n\n", w.info.Path))
	basepath, _ := filepath.Split(w.info.Path)

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

	vars, err := w.getVariables(wizardJson)
	if err != nil {
		return err
	}

	templates, err := w.getTemplates(basepath, w.outdir, wizardJson, vars)
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
		if err := w.generateFiles(templates, vars); err != nil {
			return err
		}
	}
	return nil
}

func (w *RealWizard) validateWizardJson(data WizardJson) error {
	basepath, _ := filepath.Split(w.info.Path)

	for _, t := range data.Templates {
		if err := validateTemplateFile(t.Name, basepath); err != nil {
			return err
		}
		if err := validateTemplateOutput(t.Output, data.Variables); err != nil {
			return err
		}
		if err := validateCondition(t.Condition, data.Variables); err != nil {
			return err
		}
	}

	for _, v := range data.Variables {
		if err := validateCondition(v.Condition, data.Variables); err != nil {
			return err
		}
	}
	return nil
}

func validateTemplateFile(name, basepath string) error {
	templatePath := filepath.Clean(filepath.Join(basepath, name))
	if _, err := template.ParseFiles(templatePath); err != nil {
		return err
	}

	return nil
}

func validateTemplateOutput(output string, variables []VariableJson) error {
	if len(output) == 0 {
		return errors.New("No output file specified")
	}

	if lIndex := strings.Index(output, "{"); lIndex != -1 {
		if rIndex := strings.LastIndex(output, "}"); rIndex != -1 {
			if lIndex < rIndex {
				variable := output[lIndex+1 : rIndex]
				for _, v := range variables {
					if v.Name == variable {
						return nil
					}
				}

				return errors.New(fmt.Sprint("Could not find variable:", variable))
			} else {
				return errors.New(fmt.Sprint(output, "is invalid"))
			}
		} else {
			return errors.New(fmt.Sprint("Missing } in", output))
		}
	} else if strings.LastIndex(output, "}") != -1 {
		return errors.New(fmt.Sprint("Missing { in", output))
	}

	return nil
}

func validateCondition(condition []string, variables []VariableJson) error {
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
			condition := evaluateCondition(v.Condition, vars)

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

func (w *RealWizard) getTemplates(basepath, outdir string, data WizardJson, vars map[string]interface{}) ([]TemplatePathPair, error) {
	paths := make([]TemplatePathPair, 0)
	for _, t := range data.Templates {
		if evaluateCondition(t.Condition, vars) {
			fname, err := replaceVars(t.Output, vars)
			if err != nil {
				return nil, err
			}

			input := filepath.Clean(filepath.Join(basepath, t.Name))
			output := filepath.Clean(filepath.Join(outdir, fname))

			paths = append(paths, TemplatePathPair{input, output})
		}
	}

	return paths, nil
}

func replaceVars(s string, vars map[string]interface{}) (string, error) {
	lIndex := strings.Index(s, "{")
	rIndex := strings.LastIndex(s, "}")

	if lIndex != -1 && rIndex != -1 {
		variable := s[lIndex+1 : rIndex]
		value, ok := vars[variable]
		if !ok {
			return s, errors.New(fmt.Sprint("Could not find variable:", variable))
		}

		switch value.(type) {
		case string:
			ret := s[:lIndex] + value.(string) + s[rIndex+1:]
			return ret, nil
		default:
			return s, errors.New(fmt.Sprint("Could not use variable:", variable, "since it is not a string"))
		}
	}

	return s, nil
}

func evaluateCondition(conditions []string, vars map[string]interface{}) bool {
	if len(conditions) == 0 {
		return true
	}

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

func (w *RealWizard) generateFiles(templates []TemplatePathPair, vars map[string]interface{}) error {

	for _, t := range templates {
		generator, err := template.ParseFiles(t.input)
		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)
		if err := generator.Execute(buf, vars); err != nil {
			return err
		}

		if err := w.fs.WriteFile(t.output, buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
