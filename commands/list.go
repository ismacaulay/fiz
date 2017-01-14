package commands

import (
	"github.com/ismacaulay/fiz/defines"
	"github.com/ismacaulay/fiz/io"
	"github.com/ismacaulay/fiz/wizards"

	"bytes"
	"text/template"
)

type ListCommand struct {
	provider wizards.Provider
	printer  io.Printer
}

func NewListCommand(provider wizards.Provider, printer io.Printer) *ListCommand {
	return &ListCommand{provider, printer}
}

const (
	HEADER_TMPL = "Available Wizards:"

	GROUP_TMPL = `
    {{ . }}:`

	NONE_GROUP_TMPL = `
{{- range $index, $wizard := . }}
    - {{ $wizard.Name }}
{{- end -}}
`

	GROUP_WIZARD_TMPL = `
{{- range $index, $wizard := . }}
        - {{ $wizard.Name }}
{{- end -}}
`
)

func (c *ListCommand) Run() error {
	wizards, _ := c.provider.AllAvailableWizards()

	c.printTemplate(nil, HEADER_TMPL)

	if list, ok := wizards[defines.NONE_GROUP]; ok {
		if err := c.printTemplate(list, NONE_GROUP_TMPL); err != nil {
			return err
		}
	}

	delete(wizards, defines.NONE_GROUP)
	for group, list := range wizards {
		if err := c.printTemplate(group, GROUP_TMPL); err != nil {
			return err
		}

		if err := c.printTemplate(list, GROUP_WIZARD_TMPL); err != nil {
			c.printer.Error(err)

			return err
		}
	}

	c.printer.Message("\n\n")
	return nil
}

func (c *ListCommand) printTemplate(data interface{}, tmpl string) error {
	buf := new(bytes.Buffer)
	output, err := template.New("template").Parse(tmpl)
	if err != nil {
		return err
	}

	if err := output.Execute(buf, data); err != nil {
		return err
	}

	c.printer.Message(buf.String())
	return nil
}
