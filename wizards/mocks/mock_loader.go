package wizards_mocks

import (
	"gopkg.in/stretchr/testify.v1/mock"

	"github.com/ismacaulay/fiz/wizards"
)

type MockLoader struct {
	mock.Mock
}

func NewMockLoader() *MockLoader {
	return &MockLoader{}
}

func (m *MockLoader) Load(commands []string) (wizards.Wizard, error) {
	args := m.Called(commands)
	return args.Get(0).(wizards.Wizard), args.Error(1)
}
