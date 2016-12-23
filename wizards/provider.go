package wizards

type Provider interface {
	AllAvailableWizards() (map[string][]string, error)
}

type WizardProvider struct {
}

func NewWizardProvider() *WizardProvider {
	return &WizardProvider{}
}

func (p *WizardProvider) AllAvailableWizards() (map[string][]string, error) {
	tmp := make(map[string][]string)

	tmp["hello"] = []string{"hello", "world"}
	tmp["default"] = []string{"foo", "bar", "foobar"}
	tmp["project"] = []string{"wizzy"}

	return tmp, nil
}
