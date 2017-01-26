package utils

import (
	"bytes"
	"text/template"

	"gopkg.in/stretchr/testify.v1/mock"
)

type TemplateGenerator interface {
	ValidateFile(path string) error
	Execute(tmpl string, data interface{}) (bytes.Buffer, error)
	ExecuteFile(path string, data interface{}) (bytes.Buffer, error)
}

type RealTemplateGenerator struct {
}

func NewTemplateGenerator() *RealTemplateGenerator {
	return &RealTemplateGenerator{}
}

func (t *RealTemplateGenerator) ValidateFile(input string) error {
	_, err := template.ParseFiles(input)
	return err
}

func (t *RealTemplateGenerator) Execute(tmpl string, data interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	return buf, nil
}

func (t *RealTemplateGenerator) ExecuteFile(path string, data interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	return buf, nil
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

func (m *MockTemplateGenerator) ValidateFile(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

func (m *MockTemplateGenerator) Execute(tmpl string, data interface{}) (bytes.Buffer, error) {
	args := m.Called(tmpl, data)
	return args.Get(0).(bytes.Buffer), args.Error(1)
}

func (m *MockTemplateGenerator) ExecuteFile(path string, data interface{}) (bytes.Buffer, error) {
	args := m.Called(path, data)
	return args.Get(0).(bytes.Buffer), args.Error(1)
}
