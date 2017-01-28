package wizards

import (
	"errors"
	"gopkg.in/stretchr/testify.v1/assert"
	"gopkg.in/stretchr/testify.v1/suite"
	"testing"

	"github.com/ismacaulay/fiz/utils"
)

type WizardProcessorTestSuite struct {
	suite.Suite

	Input *utils.MockInput

	Patient *WizardProcessor
}

func (td *WizardProcessorTestSuite) beforeEachCase() {
	td.Input = utils.NewMockInput()
	td.Patient = NewWizardProcessor(td.Input)
}

func (td *WizardProcessorTestSuite) TestReplaceVars() {
	assert := assert.New(td.Suite.T())
	var cases = []struct {
		name      string
		s         string
		variables map[string]interface{}
		result    string
		err       error
	}{
		{
			"Valid - Single Instance",
			"{Name}",
			map[string]interface{}{
				"Name": "HelloWorld",
			},
			"HelloWorld",
			nil,
		},
		{
			"Valid - Multiple Instance",
			"{Name}/{Name}/{Name}",
			map[string]interface{}{
				"Name": "HelloWorld",
			},
			"HelloWorld/HelloWorld/HelloWorld",
			nil,
		},
		{
			"Valid - Multiple Instance",
			"{Name1}{Name2}{Name3}",
			map[string]interface{}{
				"Name1": "Hello",
				"Name2": "World",
				"Name3": "!",
			},
			"HelloWorld!",
			nil,
		},
		{
			"Invlid - Missing variable",
			"{Name}{Name2}{Name3}",
			map[string]interface{}{
				"Name1": "Hello",
				"Name2": "World",
				"Name3": "!",
			},
			"HelloWorld!",
			errors.New("Could not find variable: Name"),
		},
		{
			"Invlid - Invalid variable",
			"{Name1}{Name2}{Name3}",
			map[string]interface{}{
				"Name1": "Hello",
				"Name2": true,
				"Name3": "!",
			},
			"HelloWorld!",
			errors.New("Could not use variable: Name2 since it is not a string"),
		},
	}

	for _, c := range cases {
		td.beforeEachCase()

		result, err := td.Patient.replaceVars(c.s, c.variables)

		if c.err == nil {
			assert.NoError(err, c.name)
			assert.Equal(c.result, result, c.name)
		} else {
			assert.Equal(err, c.err, c.name)
		}
	}
}

func (td *WizardProcessorTestSuite) TestEvaluateCondition() {
	assert := assert.New(td.Suite.T())
	var cases = []struct {
		name       string
		conditions []string
		variables  map[string]interface{}
		result     bool
	}{
		{
			"Empty Conditions",
			[]string{},
			nil,
			true,
		},
		{
			"Single Condition - true",
			[]string{"Boolean"},
			map[string]interface{}{
				"Boolean": true,
			},
			true,
		},
		{
			"Single Condition - false",
			[]string{"Boolean"},
			map[string]interface{}{
				"Boolean": false,
			},
			false,
		},
		{
			"Multiple Condition - And true",
			[]string{"Boolean1", "&&", "Boolean2"},
			map[string]interface{}{
				"Boolean1": true,
				"Boolean2": true,
			},
			true,
		},
		{
			"Multiple Condition - And false",
			[]string{"Boolean1", "&&", "Boolean2"},
			map[string]interface{}{
				"Boolean1": true,
				"Boolean2": false,
			},
			false,
		},
		{
			"Multiple Condition - Or true",
			[]string{"Boolean1", "||", "Boolean2"},
			map[string]interface{}{
				"Boolean1": true,
				"Boolean2": false,
			},
			true,
		},
		{
			"Multiple Condition - Or false",
			[]string{"Boolean1", "||", "Boolean2"},
			map[string]interface{}{
				"Boolean1": false,
				"Boolean2": false,
			},
			false,
		},
		{
			"Multiple Condition - Mixture",
			[]string{"Boolean1", "&&", "Boolean2", "&&", "Boolean3", "||", "Boolean4"},
			map[string]interface{}{
				"Boolean1": true,
				"Boolean2": true,
				"Boolean3": false,
				"Boolean4": true,
			},
			true,
		},
	}

	for _, c := range cases {
		td.beforeEachCase()

		result := td.Patient.evaluateCondition(c.conditions, c.variables)

		assert.Equal(c.result, result, c.name)
	}
}

func TestWizardProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(WizardProcessorTestSuite))
}
