package utils

import (
	"path/filepath"
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
