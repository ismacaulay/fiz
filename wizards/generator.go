package wizards

import (
	"bytes"
	"text/template"

	"github.com/ismacaulay/fiz/utils"
)

type Generator interface {
	Generate(templates []TemplatePair, vars map[string]interface{}) error
}

type OutputGenerator struct {
	fs utils.FileSystem
}

func NewOutputGenerator(fs utils.FileSystem) *OutputGenerator {
	return &OutputGenerator{fs}
}

func (g *OutputGenerator) Generate(templates []TemplatePair, vars map[string]interface{}) error {
	for _, t := range templates {
		generator, err := template.ParseFiles(t.input)
		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)
		if err := generator.Execute(buf, vars); err != nil {
			return err
		}

		if err := g.fs.WriteFile(t.output, buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
