package utils

import (
	"path/filepath"

	"github.com/ismacaulay/fiz/defines"
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
	return filepath.Join(defines.APP_DATA_DIR, "wizards")
}
