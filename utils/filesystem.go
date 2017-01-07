package utils

import (
    "io/ioutil"
)

type FileInfo struct {
    Name string
    IsDir bool
}

type FileSystem interface {
    ListDirectory(path string) ([]FileInfo, error)
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
