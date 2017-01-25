package utils

import (
	"text/template"

	"gopkg.in/stretchr/testify.v1/mock"
)

type TemplateGenerator interface {
	Validate(input string) error
	Execute(input, output string, data interface{}) error
}

type RealTemplateGenerator struct {
}

func NewTemplateGenerator() *RealTemplateGenerator {
	return &RealTemplateGenerator{}
}

func (t *RealTemplateGenerator) Validate(input string) error {
	_, err := template.ParseFiles(input)
	return err
}

func (t *RealTemplateGenerator) Execute(input, output string, data interface{}) error {
	return nil
}

/************************************
 * Mock
 ************************************/
type MockTemplateGenerator struct {
	mock.Mock
}

func NewMockTemplateGenerator() *MockTemplateGenerator {
	return &MockTemplateGenerator{}
}

func (m *MockTemplateGenerator) Validate(input string) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockTemplateGenerator) Execute(input, output string, data interface{}) error {
	args := m.Called(input, output, data)
	return args.Error(0)
}
