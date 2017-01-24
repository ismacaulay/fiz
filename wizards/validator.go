package wizards

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
)

type Validator interface {
	Validate(info WizardInfo, json WizardJson) error
}

type pathValidator interface {
	Validate(name, basepath string) error
}

type stringValidator interface {
	Validate(output string, variables []VariableJson) error
}

type expressionValidator interface {
	Validate(expression []string, variables []VariableJson) error
}

type templateValidator struct {
}

type outputPathValidator struct {
}

type conditionValidator struct {
}

type WizardValidator struct {
	pv pathValidator
	sv stringValidator
	cv expressionValidator
}

func NewWizardValidator() *WizardValidator {
	return &WizardValidator{}
}

func (v *WizardValidator) Validate(info WizardInfo, data WizardJson) error {
	basepath, _ := filepath.Split(info.Path)

	for _, t := range data.Templates {
		if err := v.pv.Validate(t.Name, basepath); err != nil {
			return err
		}
		if err := v.sv.Validate(t.Output, data.Variables); err != nil {
			return err
		}
		if err := v.cv.Validate(t.Condition, data.Variables); err != nil {
			return err
		}
	}

	for _, variable := range data.Variables {
		if err := v.cv.Validate(variable.Condition, data.Variables); err != nil {
			return err
		}
	}
	return nil
}

func (v *templateValidator) Validate(name, basepath string) error {
	templatePath := filepath.Clean(filepath.Join(basepath, name))
	if _, err := template.ParseFiles(templatePath); err != nil {
		return err
	}

	return nil
}

func (v *outputPathValidator) Validate(output string, variables []VariableJson) error {
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

func (v *conditionValidator) Validate(expression []string, variables []VariableJson) error {
	if len(expression) == 0 {
		return nil
	} else if len(expression)%2 == 0 {
		return errors.New("Not enough expression elements")
	}

	for index, element := range expression {
		if index%2 == 0 {
			for _, variable := range variables {
				if variable.Name == element {
					switch variable.Type {
					case "bool":
						continue
					default:
						return errors.New("Invalid expression")
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
