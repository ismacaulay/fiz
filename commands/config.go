package commands

import (
	"github.com/ismacaulay/fiz/utils"
)

type ConfigCommand struct {
	directory utils.DirectoryProvider
	generator utils.TemplateGenerator
	printer   utils.Printer
}

func NewConfigCommand(
	directory utils.DirectoryProvider,
	generator utils.TemplateGenerator,
	printer utils.Printer) *ConfigCommand {

	return &ConfigCommand{directory, generator, printer}
}

const (
	CONFIG_HEADER_TMPL = "Configuration:"

	CONFIG_LIST_TMPL = `
{{- range $name, $value := . }}
    {{ $name }}: {{$value}}
{{- end -}}
`
	CONFIG_FOOTER_TMPL = "\n"
)

func (c *ConfigCommand) Run() error {
	configuration := make(map[string]string)

	configuration["Wizards Directory"] = c.directory.WizardsDirectory()

	if err := c.printTemplate(CONFIG_HEADER_TMPL, nil); err != nil {
		return err
	}

	if err := c.printTemplate(CONFIG_LIST_TMPL, configuration); err != nil {
		return err
	}
	if err := c.printTemplate(CONFIG_FOOTER_TMPL, nil); err != nil {
		return err
	}

	return nil
}

func (c *ConfigCommand) printTemplate(tmpl string, data interface{}) error {
	buf, err := c.generator.Execute(tmpl, data)
	if err != nil {
		return err
	}

	c.printer.Message(buf.String())
	return nil
}
