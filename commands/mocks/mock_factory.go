package commands_mocks

import (
	"gopkg.in/stretchr/testify.v1/mock"

	"github.com/ismacaulay/fiz/commands"
)

type MockFactory struct {
	mock.Mock
}

func NewMockFactory() *MockFactory {
	return &MockFactory{}
}

func (m *MockFactory) CreateListCmd() commands.Command {
	args := m.Called()
	return args.Get(0).(commands.Command)
}

func (m *MockFactory) CreateWizardCmd(cmds []string) commands.Command {
	args := m.Called(cmds)
	return args.Get(0).(commands.Command)
}
