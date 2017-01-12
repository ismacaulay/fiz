package wizards

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Wizard interface {
	Run() error
}

type RealWizard struct {
	info WizardInfo
}

func NewWizard(info WizardInfo) *RealWizard {
	return &RealWizard{info}
}

func (w *RealWizard) Run() error {
	raw, err := ioutil.ReadFile(w.info.Path)
	if err != nil {
		return err
	}

	var wizardData WizardJson
	err = json.Unmarshal(raw, &wizardData)
	if err != nil {
		return err
	}

	fmt.Println("I READ: ", wizardData)

	return nil
}
