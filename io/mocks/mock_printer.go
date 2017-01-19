package io_mocks

import (
	"gopkg.in/stretchr/testify.v1/mock"
)

type MockPrinter struct {
	mock.Mock
}

func NewMockPrinter() *MockPrinter {
	return &MockPrinter{}
}

func (m *MockPrinter) Help() {
	m.Called()
}

func (m *MockPrinter) Version() {
	m.Called()
}

func (m *MockPrinter) Message(msg string) {
	m.Called(msg)
}

func (m *MockPrinter) Error(err error) {
	m.Called(err)
}

func (m *MockPrinter) Commands() {
	m.Called()
}
