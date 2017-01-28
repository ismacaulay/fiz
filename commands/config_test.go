package commands

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"gopkg.in/stretchr/testify.v1/suite"
	"testing"

	"bytes"
	"errors"
	"github.com/ismacaulay/fiz/utils"
)

type ConfigCommandTestSuite struct {
	suite.Suite

	Directory *utils.MockDirectoryProvider
	Generator *utils.MockTemplateGenerator
	Printer   *utils.MockPrinter

	Patient *ConfigCommand
}

func (td *ConfigCommandTestSuite) SetupTest() {
	td.beforeEachCase()
}

func (td *ConfigCommandTestSuite) beforeEachCase() {
	td.Directory = utils.NewMockDirectoryProvider()
	td.Generator = utils.NewMockTemplateGenerator()
	td.Printer = utils.NewMockPrinter()

	td.Patient = NewConfigCommand(td.Directory, td.Generator, td.Printer)
}

func (td *ConfigCommandTestSuite) TestPrintConfig() {
	assert := assert.New(td.Suite.T())

	wizards_dir := "C:\\WizrdsDir\\wizards"
	configuration := map[string]string{
		"Wizards Directory": wizards_dir,
	}

	header := "Header"
	config := "Config"
	footer := "Footer"

	td.Directory.On("WizardsDirectory").Return(wizards_dir)
	td.Generator.On("Execute", CONFIG_HEADER_TMPL, nil).Return(bytes.NewBufferString(header), nil)
	td.Generator.On("Execute", CONFIG_LIST_TMPL, configuration).Return(bytes.NewBufferString(config), nil)
	td.Generator.On("Execute", CONFIG_FOOTER_TMPL, nil).Return(bytes.NewBufferString(footer), nil)
	td.Printer.On("Message", header)
	td.Printer.On("Message", config)
	td.Printer.On("Message", footer)

	err := td.Patient.Run()

	assert.NoError(err)
	td.Directory.AssertExpectations(td.Suite.T())
	td.Generator.AssertExpectations(td.Suite.T())
	td.Printer.AssertExpectations(td.Suite.T())
}

func (td *ConfigCommandTestSuite) TestErrorWhenHeaderTemplateErrors() {
	assert := assert.New(td.Suite.T())

	expectedErr := errors.New("This is an error")

	wizards_dir := "C:\\WizrdsDir\\wizards"
	header := "Header"

	td.Directory.On("WizardsDirectory").Return(wizards_dir)
	td.Generator.On("Execute", CONFIG_HEADER_TMPL, nil).Return(bytes.NewBufferString(header), expectedErr)

	err := td.Patient.Run()

	assert.Error(err)
	assert.Equal(expectedErr, err)
}

func (td *ConfigCommandTestSuite) TestErrorWhenConfigTemplateErrors() {
	assert := assert.New(td.Suite.T())

	expectedErr := errors.New("This is an error")

	wizards_dir := "C:\\WizrdsDir\\wizards"
	configuration := map[string]string{
		"Wizards Directory": wizards_dir,
	}

	header := "Header"
	config := "Config"

	td.Directory.On("WizardsDirectory").Return(wizards_dir)
	td.Generator.On("Execute", CONFIG_HEADER_TMPL, nil).Return(bytes.NewBufferString(header), nil)
	td.Generator.On("Execute", CONFIG_LIST_TMPL, configuration).Return(bytes.NewBufferString(config), expectedErr)
	td.Printer.On("Message", header)
	td.Printer.On("Message", config)

	err := td.Patient.Run()

	assert.Error(err)
	assert.Equal(expectedErr, err)
}

func (td *ConfigCommandTestSuite) TestErrorWhenFooterTemplateErrors() {
	assert := assert.New(td.Suite.T())

	expectedErr := errors.New("This is an error")

	wizards_dir := "C:\\WizrdsDir\\wizards"
	configuration := map[string]string{
		"Wizards Directory": wizards_dir,
	}

	header := "Header"
	config := "Config"
	footer := "Footer"

	td.Directory.On("WizardsDirectory").Return(wizards_dir)
	td.Generator.On("Execute", CONFIG_HEADER_TMPL, nil).Return(bytes.NewBufferString(header), nil)
	td.Generator.On("Execute", CONFIG_LIST_TMPL, configuration).Return(bytes.NewBufferString(config), nil)
	td.Generator.On("Execute", CONFIG_FOOTER_TMPL, nil).Return(bytes.NewBufferString(footer), expectedErr)
	td.Printer.On("Message", header)
	td.Printer.On("Message", config)
	td.Printer.On("Message", footer)

	err := td.Patient.Run()

	assert.Error(err)
	assert.Equal(expectedErr, err)
}

func TestConfigCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigCommandTestSuite))
}
