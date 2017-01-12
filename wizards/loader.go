package wizards

import (
	"errors"
)

type Loader interface {
	Load(commands []string) (Wizard, error)
}

type WizardLoader struct {
	provider Provider
}

func NewWizardLoader(p Provider) *WizardLoader {
	return &WizardLoader{p}
}

func (l *WizardLoader) Load(commands []string) (Wizard, error) {
	var group, wizard string

	i := len(commands)
	switch {
	case i == 0, i > 2:
		return nil, invalidCommandError(commands)
	case i == 1:
		wizard = commands[0]
		var err error
		group, err = l.provider.FindWizardGroup(wizard)
		if err != nil {
			return nil, invalidCommandError(commands)
		}
	case i == 2:
		group = commands[0]
		wizard = commands[1]
	}

	info, err := l.provider.GetWizardInfo(group, wizard)
	if err != nil {
		return nil, invalidCommandError(commands)
	}

	return NewWizard(info), nil
}

func invalidCommandError(commands []string) error {
	msg := "Invalid command: "
	for _, c := range commands {
		msg += c + " "
	}
	msg += "\n\n"
	return errors.New(msg)
}
