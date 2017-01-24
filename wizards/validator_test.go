package wizards_test

import (
	"gopkg.in/stretchr/testify.v1/suite"
	"testing"

	"github.com/ismacaulay/fiz/wizards"
)

/*
 * WizardValidator tests
 */

/*
 * TemplateValidator tests
 */
type TemplateValidatorTestSuite struct {
	suite.Suite

	Patient *wizards.TemplateValidator
}

func (td *TemplateValidatorTestSuite) SetupTest() {
	td.Patient = &wizards.TemplateValidator{}
}

func (td *TemplateValidatorTestSuite) TestSomething() {
}

func TestTemplateValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateValidatorTestSuite))
}

/*
 * OutputPathValidator tests
 */
type OutputPathValidatorTestSuite struct {
	suite.Suite

	Patient *wizards.OutputPathValidator
}

func (td *OutputPathValidatorTestSuite) SetupTest() {
	td.Patient = &wizards.OutputPathValidator{}
}

func (td *OutputPathValidatorTestSuite) TestSomething() {
}

func TestOutputPathValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(OutputPathValidatorTestSuite))
}

/*
 * ConditionValidator tests
 */
type ConditionValidatorTestSuite struct {
	suite.Suite

	Patient *wizards.ConditionValidator
}

func (td *ConditionValidatorTestSuite) SetupTest() {
	td.Patient = &wizards.ConditionValidator{}
}

func (td *ConditionValidatorTestSuite) TestSomething() {
}

func TestConditionValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(ConditionValidatorTestSuite))
}
