package commands

import (
	"github.com/ismacaulay/fiz/utils"
	"github.com/ismacaulay/fiz/wizards"
)

type ListCommand struct {
	provider  wizards.Provider
	generator utils.TemplateGenerator
	printer   utils.Printer
}

func NewListCommand(provider wizards.Provider, generator utils.TemplateGenerator, printer utils.Printer) *ListCommand {
	return &ListCommand{provider, generator, printer}
}

const (
	LIST_HEADER_TMPL = "Available Wizards:"

	LIST_GROUP_TMPL = `
    {{ . }}:`

	LIST_NONE_GROUP_TMPL = `
{{- range $index, $wizard := . }}
    - {{ $wizard.Name }}
{{- end -}}
`

	LIST_GROUP_WIZARD_TMPL = `
{{- range $index, $wizard := . }}
        - {{ $wizard.Name }}
{{- end -}}
`
	LIST_NO_WIZARDS_TMPL = `
    No wizards available`

	LIST_FOOTER_TMPL = "\n"
)

func (c *ListCommand) Run() error {
	wizards, _ := c.provider.AllAvailableWizards()

	c.printTemplate(LIST_HEADER_TMPL, nil)

	if len(wizards) == 0 {
		c.printTemplate(LIST_NO_WIZARDS_TMPL, nil)
	} else {
		if list, ok := wizards[utils.NONE_GROUP]; ok {
			if err := c.printTemplate(LIST_NONE_GROUP_TMPL, list); err != nil {
				return err
			}
		}

		delete(wizards, utils.NONE_GROUP)
		for group, list := range wizards {
			if err := c.printTemplate(LIST_GROUP_TMPL, group); err != nil {
				return err
			}

			if err := c.printTemplate(LIST_GROUP_WIZARD_TMPL, list); err != nil {
				return err
			}
		}
	}

	c.printTemplate(LIST_FOOTER_TMPL, nil)
	return nil
}

func (c *ListCommand) printTemplate(tmpl string, data interface{}) error {
	buf, err := c.generator.Execute(tmpl, data)
	if err != nil {
		return err
	}

	c.printer.Message(buf.String())
	return nil
}
