package wizards

import (
    "fmt"
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
    fmt.Println("Running:", w.info.Path)
    return nil
}
