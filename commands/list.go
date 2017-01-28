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
	NO_WIZARDS_TMPL = `
    No wizards available`

	FOOTER_TMPL = "\n"
)

func (c *ListCommand) Run() error {
	wizards, _ := c.provider.AllAvailableWizards()

	c.printTemplate(HEADER_TMPL, nil)

	if len(wizards) == 0 {
		c.printTemplate(NO_WIZARDS_TMPL, nil)
	} else {
		if list, ok := wizards[utils.NONE_GROUP]; ok {
			if err := c.printTemplate(NONE_GROUP_TMPL, list); err != nil {
				return err
			}
		}

		delete(wizards, utils.NONE_GROUP)
		for group, list := range wizards {
			if err := c.printTemplate(GROUP_TMPL, group); err != nil {
				return err
			}

			if err := c.printTemplate(GROUP_WIZARD_TMPL, list); err != nil {
				return err
			}
		}
	}

	c.printTemplate(FOOTER_TMPL, nil)
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
