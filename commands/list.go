package commands

import (
	"github.com/ismacaulay/fiz/output"
	"github.com/ismacaulay/fiz/wizards"

	"text/template"
	"bytes"
)

type ListCommand struct {
	provider wizards.Provider
	printer  output.Printer
}

func NewListCommand(provider wizards.Provider, printer output.Printer) *ListCommand {
	return &ListCommand{provider, printer}
}

const (
	LIST_WIZARDS_OUTPUT_TMPL = `Available Wizards:
{{- template "check_length" . -}}

{{- define "check_length" -}}
{{- $length := len . -}}
	{{- if eq $length 0 -}}
		{{- template "no_wizards_found" -}}
	{{- else -}}
		{{- template "print_categories" . -}}
	{{- end -}}
{{- end -}}

{{- define "no_wizards_found" }}
    No wizards found
    {{- template "new_line" -}}
{{- end -}}

{{- define "print_categories" -}}
{{- range $category, $wizards := . }}
    {{ $category }}
   	{{- template "print_wizards" $wizards -}}
   	{{- template "new_line" -}}
{{- end -}}
{{- end -}}

{{- define "print_wizards" -}}
{{- range $index, $wizard := . }}
        - {{ $wizard.Name }}
{{- end -}}
{{- end -}}

{{- define "new_line" -}}
{{"\n"}}
{{- end -}}
`
)

func (c *ListCommand) Run() error {
	wizards, _ := c.provider.AllAvailableWizards()

	buf := new(bytes.Buffer)
	output, err := template.New("list_wizards_output_tmpl").Parse(LIST_WIZARDS_OUTPUT_TMPL)
	if(err != nil) {
		return err
	}

	if err := output.Execute(buf, wizards); err != nil {
		return err
	}
	c.printer.Message(buf.String())
	return nil
}
