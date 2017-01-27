package utils

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"path/filepath"
	"testing"
)

func TestWizardsDirectoryCorrect(t *testing.T) {
	assert := assert.New(t)

	expectedPath := filepath.Join(APP_DATA_DIR, "wizards")

	patient := NewDirectoryProvider()

	assert.Equal(expectedPath, patient.WizardsDirectory())
}
