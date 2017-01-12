package wizards

type VariableJson struct {
	Name      string
	Type      string
	Default   interface{}
	Condition []string
}

type WizardJson struct {
	Templates []string
	Variables []VariableJson
}
