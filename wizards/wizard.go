package wizards

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/ismacaulay/fiz/utils"
)

type Wizard interface {
	Run() error
}

type RealWizard struct {
	info WizardInfo
	fs   utils.FileSystem
}

func NewWizard(info WizardInfo, fs utils.FileSystem) *RealWizard {
	return &RealWizard{info, fs}
}

func (w *RealWizard) Run() error {
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

	fmt.Println("Wizard valid.... lets go.")
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
				if variable.Name == element && variable.Default != nil {
					switch variable.Default.(type) {
					case bool:
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
