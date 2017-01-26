package utils

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"path/filepath"
	"testing"

	"github.com/ismacaulay/fiz/defines"
)

func TestWizardsDirectoryCorrect(t *testing.T) {
	assert := assert.New(t)

	expectedPath := filepath.Join(defines.APP_DATA_DIR, "wizards")

	patient := NewDirectoryProvider()

	assert.Equal(expectedPath, patient.WizardsDirectory())
}
