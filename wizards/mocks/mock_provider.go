package wizards_mocks

import (
	"gopkg.in/stretchr/testify.v1/mock"

	"github.com/ismacaulay/fiz/wizards"
)

type MockProvider struct {
	mock.Mock
}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (m *MockProvider) Run() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockProvider) AllAvailableWizards() (map[string][]wizards.WizardInfo, error) {
	args := m.Called()
	return args.Get(0).(map[string][]wizards.WizardInfo), args.Error(1)
}

func (m *MockProvider) GetWizardInfo(group, wizard string) (wizards.WizardInfo, error) {
	args := m.Called(group, wizard)
	return args.Get(0).(wizards.WizardInfo), args.Error(1)
}

func (m *MockProvider) FindWizardGroup(wizard string) (string, error) {
	args := m.Called(wizard)
	return args.String(0), args.Error(1)
}
