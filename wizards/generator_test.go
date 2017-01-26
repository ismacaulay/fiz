package wizards

import (
	"bytes"
	"errors"
	"gopkg.in/stretchr/testify.v1/assert"
	"gopkg.in/stretchr/testify.v1/suite"
	"testing"

	"github.com/ismacaulay/fiz/utils"
)

type OutputGeneratorTestSuite struct {
	suite.Suite

	Generator  *utils.MockTemplateGenerator
	FileSystem *utils.MockFileSystem

	Patient *OutputGenerator
}

func (td *OutputGeneratorTestSuite) beforeEachCase() {
	td.Generator = utils.NewMockTemplateGenerator()
	td.FileSystem = utils.NewMockFileSystem()

	td.Patient = NewOutputGenerator(td.Generator, td.FileSystem)
}

func (td *OutputGeneratorTestSuite) TestExecuteFileOnEachInput() {
	assert := assert.New(td.Suite.T())

	var cases = []struct {
		name string
		err  error
	}{
		{
			"No Error",
			nil,
		},
		{
			"Error",
			errors.New("Invalid input"),
		},
	}

	buf := new(bytes.Buffer)
	templates := []TemplatePair{
		TemplatePair{"input.cpp", "output.cpp"},
		TemplatePair{"input.h", "output.h"},
	}
	vars := make(map[string]interface{})
	vars["hello"] = "world"
	vars["world"] = true

	for _, c := range cases {
		td.beforeEachCase()

		td.Generator.On("ExecuteFile", "input.cpp", vars).Return(buf, c.err)
		td.Generator.On("ExecuteFile", "input.h", vars).Return(buf, c.err)
		td.FileSystem.On("WriteFile", "output.cpp", buf.Bytes()).Return(nil)
		td.FileSystem.On("WriteFile", "output.h", buf.Bytes()).Return(nil)

		err := td.Patient.Generate(templates, vars)
		if err != nil {
			assert.Equal(err, c.err, c.name)
		}
	}
}

func (td *OutputGeneratorTestSuite) TestWriteBufferToEachOutput() {
	assert := assert.New(td.Suite.T())

	var cases = []struct {
		name string
		data string
		err  error
	}{
		{
			"No Error",
			"This is some data",
			nil,
		},
		{
			"Error",
			"This is some different data",
			errors.New("Fails to write"),
		},
	}

	templates := []TemplatePair{
		TemplatePair{"input.cpp", "output.cpp"},
		TemplatePair{"input.h", "output.h"},
	}
	vars := make(map[string]interface{})
	vars["hello"] = "world"
	vars["world"] = true

	for _, c := range cases {
		td.beforeEachCase()

		buf := bytes.NewBufferString(c.data)

		td.Generator.On("ExecuteFile", "input.cpp", vars).Return(buf, nil)
		td.Generator.On("ExecuteFile", "input.h", vars).Return(buf, nil)
		td.FileSystem.On("WriteFile", "output.cpp", buf.Bytes()).Return(c.err)
		td.FileSystem.On("WriteFile", "output.h", buf.Bytes()).Return(c.err)

		err := td.Patient.Generate(templates, vars)
		if err != nil {
			assert.Equal(err, c.err, c.name)
		}
	}
}

func TestOutputGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(OutputGeneratorTestSuite))
}
