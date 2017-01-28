package utils

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"gopkg.in/stretchr/testify.v1/suite"
	"testing"

	"io/ioutil"
	"os"
	"path/filepath"
)

type TemplateTestSuite struct {
	suite.Suite

	Patient *RealTemplateGenerator
}

func (td *TemplateTestSuite) SetupTest() {
	td.beforeEachCase()
}

func (td *TemplateTestSuite) beforeEachCase() {
	td.Patient = NewTemplateGenerator()
}

func (td *TemplateTestSuite) TestValidateFile() {
	assert := assert.New(td.Suite.T())

	cases := []struct {
		name        string
		content     string
		createFile  bool
		shouldError bool
	}{
		{
			"Valid",
			"{{ . }}",
			true,
			false,
		},
		{
			"No File",
			"",
			false,
			true,
		},
		{
			"Invalid template data",
			"{{ Invalid }}",
			true,
			true,
		},
	}

	for _, c := range cases {
		td.beforeEachCase()

		content := []byte(c.content)
		dir, err := ioutil.TempDir("", "TestValidateValidTemplateFile")
		assert.NoError(err)

		defer os.RemoveAll(dir)

		templateFile := filepath.Join(dir, "template.txt")
		if c.createFile {
			err = ioutil.WriteFile(templateFile, content, 0666)
			assert.NoError(err)
		}

		err = td.Patient.ValidateFile(templateFile)

		if c.shouldError {
			assert.Error(err, c.name)
		} else {
			assert.NoError(err, c.name)
		}
	}
}

func (td *TemplateTestSuite) TestExecute() {
	assert := assert.New(td.Suite.T())

	cases := []struct {
		name        string
		tmpl        string
		data        string
		shouldError bool
	}{
		{
			"Valid",
			"{{ . }}",
			"Hello World!",
			false,
		},
		{
			"Invalid template",
			"{{ Invalid }}",
			"Hello World",
			true,
		},
		{
			"Invalid Data",
			"{{ .Name }}",
			"Hello World",
			true,
		},
	}

	for _, c := range cases {
		td.beforeEachCase()

		buf, err := td.Patient.Execute(c.tmpl, c.data)

		if c.shouldError {
			assert.Error(err, c.name)
			assert.Nil(buf, c.name)
		} else {
			assert.NoError(err, c.name)
			assert.Equal(c.data, buf.String(), c.name)
		}
	}
}

func (td *TemplateTestSuite) TestExecuteFile() {
	assert := assert.New(td.Suite.T())

	cases := []struct {
		name        string
		content     string
		createFile  bool
		data        string
		shouldError bool
	}{
		{
			"Valid",
			"{{ . }}",
			true,
			"Hello World",
			false,
		},
		{
			"No File",
			"{{ . }}",
			false,
			"Hello World",
			true,
		},
		{
			"Invalid template data",
			"{{ .Invalid }}",
			true,
			"Hello World",
			true,
		},
	}

	for _, c := range cases {
		td.beforeEachCase()

		content := []byte(c.content)
		dir, err := ioutil.TempDir("", "TestValidateValidTemplateFile")
		assert.NoError(err)

		defer os.RemoveAll(dir)

		templateFile := filepath.Join(dir, "template.txt")
		if c.createFile {
			err = ioutil.WriteFile(templateFile, content, 0666)
			assert.NoError(err)
		}

		buf, err := td.Patient.ExecuteFile(templateFile, c.data)

		if c.shouldError {
			assert.Error(err, c.name)
			assert.Nil(buf, c.name)
		} else {
			assert.NoError(err, c.name)
			assert.Equal(c.data, buf.String(), c.name)
		}
	}
}

func TestTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateTestSuite))
}
