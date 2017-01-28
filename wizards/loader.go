package wizards

import (
	"errors"
	"strings"

	"github.com/ismacaulay/fiz/utils"

	"gopkg.in/stretchr/testify.v1/mock"
)

type Loader interface {
	Load(commands []string) (Wizard, error)
}

type WizardLoader struct {
	provider Provider
	factory  Factory
	fs       utils.FileSystem
}

func NewWizardLoader(p Provider, f Factory, fs utils.FileSystem) *WizardLoader {
	return &WizardLoader{p, f, fs}
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

	outDir, err := l.fs.GetCwd()
	if err != nil {
		return nil, err
	}

	return l.factory.Create(info, outDir), nil
}

func invalidCommandError(commands []string) error {
	msg := "Invalid command: "
	msg += strings.Join(commands[:], " ")
	return errors.New(msg)
}

/************************************
 * Mock
 ************************************/
type MockLoader struct {
	mock.Mock
}

func NewMockLoader() *MockLoader {
	return &MockLoader{}
}

func (m *MockLoader) Load(commands []string) (Wizard, error) {
	args := m.Called(commands)
	return args.Get(0).(Wizard), args.Error(1)
}
