package utils


import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/stretchr/testify.v1/mock"
)

type Input interface {
	GetBoolean(message string) (bool, error)
	GetString(message string) (string, error)
}

type CliInput struct {
}

func NewCliInput() *CliInput {
	return &CliInput{}
}

func (i *CliInput) GetBoolean(message string) (bool, error) {
	text, err := getInput(message + " [yes/no]: ")
	if err != nil {
		return false, err
	}

	switch strings.ToLower(strings.TrimSpace(text)) {
	case "y", "yes":
		return true, nil
	case "n", "no", "":
		return false, nil
	default:
		return false, errors.New("Invalid input")
	}
}

func (i *CliInput) GetString(message string) (string, error) {
	text, err := getInput(message + ": ")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func getInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	return reader.ReadString('\n')
}

/************************************
 * Mock
 ************************************/
type MockInput struct {
	mock.Mock
}

func NewMockInput() *MockInput {
	return &MockInput{}
}

func (m *MockInput) GetBoolean(message string) (bool, error) {
	args := m.Called(message)
	return args.Bool(0), args.Error(1)
}

func (m *MockInput) GetString(message string) (string, error) {
	args := m.Called(message)
	return args.String(0), args.Error(1)
}
