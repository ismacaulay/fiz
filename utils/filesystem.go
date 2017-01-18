package utils

import (
	"io/ioutil"
	"os"
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
	return ioutil.WriteFile(path, data, 0644)
}

func (fs *RealFileSystem) GetCwd() (string, error) {
	return os.Getwd()
}
