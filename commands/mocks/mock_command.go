package commands_mocks

import (
	"gopkg.in/stretchr/testify.v1/mock"
)

type MockCommand struct {
	mock.Mock
}

func NewMockCommand() *MockCommand {
	return &MockCommand{}
}

func (m *MockCommand) Run() error {
	args := m.Called()
	return args.Error(0)
}
