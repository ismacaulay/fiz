package wizards

type VariableJson struct {
	Name      string
	Default   interface{}
	Condition []string
}

type TemplateJson struct {
	Name      string
	Condition []string
}

type WizardJson struct {
	Templates []TemplateJson
	Variables []VariableJson
}
