package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/stretchr/testify.v1/mock"
)

type FileInfo struct {
	Name  string
	IsDir bool
}

type FileSystem interface {
	ListDirectory(path string) ([]FileInfo, error)
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
	GetCwd() (string, error)
}

type RealFileSystem struct {
}

func NewFileSystem() *RealFileSystem {
	return &RealFileSystem{}
}

func (fs *RealFileSystem) ListDirectory(path string) ([]FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	ret := make([]FileInfo, 0)
	for _, file := range files {
		ret = append(ret, FileInfo{file.Name(), file.IsDir()})
	}

	return ret, nil
}

func (fs *RealFileSystem) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func (fs *RealFileSystem) WriteFile(path string, data []byte) error {
	dir, _ := filepath.Split(path)
	if len(dir) > 0 {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(path, data, 0644)
}

func (fs *RealFileSystem) GetCwd() (string, error) {
	return os.Getwd()
}

/************************************
 * Mock
 ************************************/
type MockFileSystem struct {
	mock.Mock
}

func NewMockFileSystem() *MockFileSystem {
	return &MockFileSystem{}
}

func (m *MockFileSystem) ListDirectory(path string) ([]FileInfo, error) {
	args := m.Called(path)
	return args.Get(0).([]FileInfo), args.Error(1)
}

func (m *MockFileSystem) ReadFile(path string) ([]byte, error) {
	args := m.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockFileSystem) WriteFile(path string, data []byte) error {
	args := m.Called(path, data)
	return args.Error(0)
}

func (m *MockFileSystem) GetCwd() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}
