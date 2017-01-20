package wizards_mocks

import (
	"gopkg.in/stretchr/testify.v1/mock"
)

type MockWizard struct {
	mock.Mock
}

func NewMockWizard() *MockWizard {
	return &MockWizard{}
}

func (m *MockWizard) Run() error {
	args := m.Called()
	return args.Error(0)
}
