package wizards

type VariableJson struct {
	Name      string
	Type      string
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
