package app

import (
	"github.com/ismacaulay/fiz/io"
	"github.com/ismacaulay/fiz/utils"
)

type External interface {
	FileSystem() utils.FileSystem
	DirectoryProvider() utils.DirectoryProvider
	TemplateGenerator() utils.TemplateGenerator
	Input() io.Input
	Printer() io.Printer
}

type RealExternal struct {
	fileSystem        utils.FileSystem
	directoryProvider utils.DirectoryProvider
	templateGenerator utils.TemplateGenerator
	input             io.Input
	printer           io.Printer
}

func NewExternal(version string) *RealExternal {
	filesystem := utils.NewFileSystem()
	directoryProvider := utils.NewDirectoryProvider()
	templateGenerator := utils.NewTemplateGenerator()
	input := io.NewCliInput()
	printer := io.NewTextPrinter(version)

	return &RealExternal{filesystem, directoryProvider, templateGenerator, input, printer}
}

func (e *RealExternal) FileSystem() utils.FileSystem {
	return e.fileSystem
}

func (e *RealExternal) DirectoryProvider() utils.DirectoryProvider {
	return e.directoryProvider
}

func (e *RealExternal) TemplateGenerator() utils.TemplateGenerator {
	return e.templateGenerator
}

func (e *RealExternal) Input() io.Input {
	return e.input
}

func (e *RealExternal) Printer() io.Printer {
	return e.printer
}
