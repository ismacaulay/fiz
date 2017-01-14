package app

import (
	"github.com/ismacaulay/fiz/io"
	"github.com/ismacaulay/fiz/utils"
)

type External interface {
	FileSystem() utils.FileSystem
	DirectoryProvider() utils.DirectoryProvider
	Input() io.Input
	Printer() io.Printer
}

type RealExternal struct {
	fileSystem        utils.FileSystem
	directoryProvider utils.DirectoryProvider
	input             io.Input
	printer           io.Printer
}

func NewExternal(version string) *RealExternal {
	filesystem := utils.NewFileSystem()
	directoryProvider := utils.NewDirectoryProvider()
	input := io.NewCliInput()
	printer := io.NewTextPrinter(version)

	return &RealExternal{filesystem, directoryProvider, input, printer}
}

func (e *RealExternal) FileSystem() utils.FileSystem {
	return e.fileSystem
}

func (e *RealExternal) DirectoryProvider() utils.DirectoryProvider {
	return e.directoryProvider
}

func (e *RealExternal) Input() io.Input {
	return e.input
}

func (e *RealExternal) Printer() io.Printer {
	return e.printer
}
