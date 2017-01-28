package utils

import (
	"path/filepath"

	"gopkg.in/stretchr/testify.v1/mock"
)

type DirectoryProvider interface {
	WizardsDirectory() string
}

type RealDirectoryProvider struct {
}

func NewDirectoryProvider() *RealDirectoryProvider {
	return &RealDirectoryProvider{}
}

func (dp *RealDirectoryProvider) WizardsDirectory() string {
	return filepath.Join(APP_DATA_DIR, "wizards")
}

/************************************
 * Mock
 ************************************/
type MockDirectoryProvider struct {
	mock.Mock
}

func NewMockDirectoryProvider() *MockDirectoryProvider {
	return &MockDirectoryProvider{}
}

func (m *MockDirectoryProvider) WizardsDirectory() string {
	args := m.Called()
	return args.String(0)
}
