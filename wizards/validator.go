package wizards

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ismacaulay/fiz/utils"
)

type Validator interface {
	Validate(info WizardInfo, json WizardJson) error
}

type WizardValidator struct {
	t utils.TemplateGenerator
}

func NewWizardValidator(t utils.TemplateGenerator) *WizardValidator {
	return &WizardValidator{t}
}

func (v *WizardValidator) Validate(info WizardInfo, data WizardJson) error {
	basepath, _ := filepath.Split(info.Path)

	for _, t := range data.Templates {
		if err := v.validateTemplate(t.Name, basepath); err != nil {
			return err
		}
		if err := v.validateOutputPath(t.Output, data.Variables); err != nil {
			return err
		}
		if err := v.validateCondition(t.Condition, data.Variables); err != nil {
			return err
		}
	}

	for _, variable := range data.Variables {
		if err := v.validateCondition(variable.Condition, data.Variables); err != nil {
			return err
		}
	}
	return nil
}

func (v *WizardValidator) validateTemplate(name, basepath string) error {
	templatePath := filepath.Clean(filepath.Join(basepath, name))
	return v.t.Validate(templatePath)
}

func (v *WizardValidator) validateOutputPath(output string, variables []VariableJson) error {
	if len(output) == 0 {
		return errors.New("No output file specified")
	}

	processedOutput := output
	for {
		lIndex := strings.Index(processedOutput, "{")
		rIndex := strings.Index(processedOutput, "}")
		if lIndex > -1 && rIndex > -1 {
			if lIndex < rIndex {
				variable := processedOutput[lIndex+1 : rIndex]
				found := false
				for _, v := range variables {
					if v.Name == variable {
						if v.Type == "string" {
							processedOutput = processedOutput[rIndex+1:]
							found = true
							break
						}
						return errors.New("Variable needs to be a string")
					}
				}
				if !found {
					return errors.New(fmt.Sprint("Could not find variable: ", variable))
				}
			} else {
				return errors.New(fmt.Sprint(output, "is invalid"))
			}
		} else if lIndex > -1 {
			return errors.New(fmt.Sprint("Missing } in ", output))
		} else if rIndex > -1 {
			return errors.New(fmt.Sprint("Missing { in ", output))
		} else {
			return nil
		}
	}
}

func (v *WizardValidator) validateCondition(expression []string, variables []VariableJson) error {
	if len(expression) == 0 {
		return nil
	} else if len(expression)%2 == 0 {
		return errors.New("Not enough expression elements")
	}

	for index, element := range expression {
		if index%2 == 0 {
			found := false
			for _, variable := range variables {
				if variable.Name == element {
					found = true
					switch variable.Type {
					case "bool":
						continue
					default:
						return errors.New("Invalid expression")
					}

				}
			}
			if !found {
				return errors.New("Invalid expression")
			}
		} else {
			switch element {
			case "&&", "||":
				continue
			default:
				return errors.New("Invalid operator")
			}
		}
	}

	return nil
}
