package wizards

import (
	"errors"
	"gopkg.in/stretchr/testify.v1/assert"
	"gopkg.in/stretchr/testify.v1/suite"
	"path/filepath"
	"testing"

	"github.com/ismacaulay/fiz/utils"
)

/*
 * WizardValidator tests
 */
type WizardValidatorTestSuite struct {
	suite.Suite

	Template *utils.MockTemplateGenerator

	Patient *WizardValidator
}

func (td *WizardValidatorTestSuite) SetupTest() {
	td.Template = utils.NewMockTemplateGenerator()
	td.Patient = NewWizardValidator(td.Template)
}

func (td *WizardValidatorTestSuite) TestValidate() {
}

func (td *WizardValidatorTestSuite) TestValidateTemplate() {
	assert := assert.New(td.Suite.T())

	var cases = []struct {
		name               string
		filename, basepath string
		errMsg             string
	}{
		{"Valid", "hello.world", "/opt/helloworld", ""},
		{"Invalid", "hello.world", "C:\\user\\helloworld", "Invalid"},
	}

	for _, c := range cases {
		td.Template.On("Validate", filepath.Clean(filepath.Join(c.basepath, c.filename))).Return(errors.New(c.errMsg))

		err := td.Patient.validateTemplate(c.filename, c.basepath)

		if err != nil {
			assert.EqualError(err, c.errMsg, c.name)
		}
		td.Template.AssertExpectations(td.Suite.T())
	}
}

func (td *WizardValidatorTestSuite) TestValidateOutputPath() {
	assert := assert.New(td.Suite.T())
	var cases = []struct {
		name      string
		output    string
		variables []VariableJson
		errMsg    string
	}{
		{"Empty output", "", []VariableJson{}, "No output file specified"},
		{"Valid - output without variable", "/opt/Hello/World.cpp", []VariableJson{}, ""},
		{
			"Valid - output with 1 string var",
			"/opt/Hello/{ClassName}.cpp",
			[]VariableJson{VariableJson{"ClassName", "string", nil}},
			"",
		},
		{
			"Valid - output with duplicate string var",
			"/opt/Hello/{ClassName}/Test{ClassName}.cpp",
			[]VariableJson{VariableJson{"ClassName", "string", nil}},
			"",
		},
		{
			"Valid - output with duplicate string var",
			"/opt/Hello/{ClassName}/Test{ClassName2}.cpp",
			[]VariableJson{
				VariableJson{"ClassName", "string", nil},
				VariableJson{"Boolean", "bool", nil},
				VariableJson{"ClassName2", "string", nil},
			},
			"",
		},
		{
			"Invalid - output with 1 bool var",
			"/opt/Hello/{ClassName}.cpp",
			[]VariableJson{VariableJson{"ClassName", "bool", nil}},
			"Variable needs to be a string",
		},
		{
			"Invalid - missing var",
			"/opt/Hello/{ClassName}.cpp",
			[]VariableJson{VariableJson{"Name", "bool", nil}},
			"Could not find variable: ClassName",
		},
		{
			"Invalid - missing }",
			"/opt/Hello/{ClassName.cpp",
			[]VariableJson{VariableJson{"ClassName", "bool", nil}},
			"Missing } in /opt/Hello/{ClassName.cpp",
		},
		{
			"Invalid - missing {",
			"/opt/Hello/ClassName}.cpp",
			[]VariableJson{VariableJson{"ClassName", "bool", nil}},
			"Missing { in /opt/Hello/ClassName}.cpp",
		},
		{
			"Invalid - missing {",
			"/opt/Hello/{ClassName}/TestClassName}.cpp",
			[]VariableJson{VariableJson{"ClassName", "string", nil}},
			"Missing { in /opt/Hello/{ClassName}/TestClassName}.cpp",
		},
	}

	for _, c := range cases {
		err := td.Patient.validateOutputPath(c.output, c.variables)
		if err != nil {
			assert.EqualError(err, c.errMsg, c.name)
		}
	}
}

func (td *WizardValidatorTestSuite) TestValidateCondition() {
	assert := assert.New(td.Suite.T())
	var cases = []struct {
		name       string
		expression []string
		variables  []VariableJson
		errMsg     string
	}{
		{
			"Valid - Empty expression",
			[]string{},
			[]VariableJson{}, "",
		},
		{
			"Invalid - Even length expression",
			[]string{"Hello", "World"},
			[]VariableJson{},
			"Not enough expression elements",
		},
		{
			"Invalid - Even length expression 2",
			[]string{"Hello", "World", "This", "Error"},
			[]VariableJson{},
			"Not enough expression elements",
		},
		{
			"Valid - Single boolean",
			[]string{"Boolean"},
			[]VariableJson{VariableJson{"Boolean", "bool", nil}},
			"",
		},
		{
			"Invalid - Single boolean without var",
			[]string{"Boolean"},
			[]VariableJson{VariableJson{"String", "bool", nil}},
			"Invalid expression",
		},
		{
			"Invalid - Non boolean",
			[]string{"String"},
			[]VariableJson{VariableJson{"String", "string", nil}},
			"Invalid expression",
		},
		{
			"Valid - 2 booleans - Or",
			[]string{"Boolean1", "||", "Boolean2"},
			[]VariableJson{
				VariableJson{"Boolean1", "bool", nil},
				VariableJson{"Boolean2", "bool", nil},
			},
			"",
		},
		{
			"Valid - 2 booleans - And",
			[]string{"Boolean1", "&&", "Boolean2"},
			[]VariableJson{
				VariableJson{"Boolean1", "bool", nil},
				VariableJson{"Boolean2", "bool", nil},
			},
			"",
		},
		{
			"Valid - 3 booleans",
			[]string{"Boolean1", "&&", "Boolean2", "||", "Boolean3"},
			[]VariableJson{
				VariableJson{"Boolean1", "bool", nil},
				VariableJson{"Boolean2", "bool", nil},
				VariableJson{"Boolean3", "bool", nil},
			},
			"",
		},
		{
			"Invalid - 1 bool, 1 non bool",
			[]string{"Boolean1", "||", "String"},
			[]VariableJson{
				VariableJson{"Boolean1", "bool", nil},
				VariableJson{"String", "string", nil},
			},
			"Invalid expression",
		},
		{
			"Invalid - 2 bool - invalid operator",
			[]string{"Boolean1", "--", "String"},
			[]VariableJson{
				VariableJson{"Boolean1", "bool", nil},
				VariableJson{"String", "string", nil},
			},
			"Invalid operator",
		},
	}

	for _, c := range cases {
		err := td.Patient.validateCondition(c.expression, c.variables)
		if err != nil {
			assert.EqualError(err, c.errMsg, c.name)
		}
	}
}

func TestWizardValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(WizardValidatorTestSuite))
}
