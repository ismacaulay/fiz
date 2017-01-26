package wizards

import (
	"github.com/ismacaulay/fiz/utils"
)

type Generator interface {
	Generate(templates []TemplatePair, vars map[string]interface{}) error
}

type OutputGenerator struct {
	generator utils.TemplateGenerator
	fs        utils.FileSystem
}

func NewOutputGenerator(generator utils.TemplateGenerator, fs utils.FileSystem) *OutputGenerator {
	return &OutputGenerator{generator, fs}
}

func (g *OutputGenerator) Generate(templates []TemplatePair, vars map[string]interface{}) error {
	for _, t := range templates {
		buf, err := g.generator.ExecuteFile(t.input, vars)
		if err != nil {
			return err
		}

		if err := g.fs.WriteFile(t.output, buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
