package wizards

type VariableJson struct {
	Name string
	Type string
}

type WizardJson struct {
	Templates []string
	Variables []VariableJson
}
