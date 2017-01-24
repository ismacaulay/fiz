package wizards

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ismacaulay/fiz/io"
)

type TemplatePair struct {
	input, output string
}

type Processor interface {
	Process(info WizardInfo, outdir string, data WizardJson) (map[string]interface{}, []TemplatePair, error)
}

type WizardProcessor struct {
	input io.Input
}

func NewWizardProcessor(input io.Input) *WizardProcessor {
	return &WizardProcessor{input}
}

func (p *WizardProcessor) Process(info WizardInfo, outdir string, data WizardJson) (map[string]interface{}, []TemplatePair, error) {
	vars, err := p.getVariables(data)
	if err != nil {
		return nil, nil, err
	}

	basepath, _ := filepath.Split(info.Path)
	templates, err := p.getTemplates(basepath, outdir, data, vars)
	if err != nil {
		return nil, nil, err
	}

	return vars, templates, nil
}

func (p *WizardProcessor) getVariables(data WizardJson) (map[string]interface{}, error) {
	vars := make(map[string]interface{})
	var err error

	for _, v := range data.Variables {
		if len(v.Condition) == 0 {
			vars, err = p.getVariable(v, vars)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, v := range data.Variables {
		if len(v.Condition) > 0 {
			condition := evaluateCondition(v.Condition, vars)

			if condition {
				vars, err = p.getVariable(v, vars)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return vars, nil
}

func (p *WizardProcessor) getVariable(v VariableJson, vars map[string]interface{}) (map[string]interface{}, error) {
	switch v.Type {
	case "bool":
		value, err := p.input.GetBoolean(v.Name)
		if err != nil {
			return vars, err
		}
		vars[v.Name] = value
	default:
		value, err := p.input.GetString(v.Name)
		if err != nil {
			return vars, err
		}
		vars[v.Name] = value
	}

	return vars, nil
}

func (p *WizardProcessor) getTemplates(basepath, outdir string, data WizardJson, vars map[string]interface{}) ([]TemplatePair, error) {
	paths := make([]TemplatePair, 0)
	for _, t := range data.Templates {
		if evaluateCondition(t.Condition, vars) {
			fname, err := replaceVars(t.Output, vars)
			if err != nil {
				return nil, err
			}

			input := filepath.Clean(filepath.Join(basepath, t.Name))
			output := filepath.Clean(filepath.Join(outdir, fname))

			paths = append(paths, TemplatePair{input, output})
		}
	}

	return paths, nil
}

func replaceVars(s string, vars map[string]interface{}) (string, error) {
	lIndex := strings.Index(s, "{")
	rIndex := strings.LastIndex(s, "}")

	if lIndex != -1 && rIndex != -1 {
		variable := s[lIndex+1 : rIndex]
		value, ok := vars[variable]
		if !ok {
			return s, errors.New(fmt.Sprint("Could not find variable:", variable))
		}

		switch value.(type) {
		case string:
			ret := s[:lIndex] + value.(string) + s[rIndex+1:]
			return ret, nil
		default:
			return s, errors.New(fmt.Sprint("Could not use variable:", variable, "since it is not a string"))
		}
	}

	return s, nil
}

func evaluateCondition(conditions []string, vars map[string]interface{}) bool {
	if len(conditions) == 0 {
		return true
	}

	condition := vars[conditions[0]].(bool)
	nextOperator := "&&"
	for _, c := range conditions[1:] {
		switch c {
		case "&&", "||":
			nextOperator = c
		default:
			switch nextOperator {
			case "&&":
				condition = condition && vars[c].(bool)
			case "||":
				condition = condition || vars[c].(bool)
			}
		}
	}
	return condition
}
