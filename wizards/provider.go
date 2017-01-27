package wizards

import (
	"github.com/ismacaulay/fiz/utils"
	"gopkg.in/stretchr/testify.v1/mock"

	"errors"
	"fmt"
	"path/filepath"
)

type WizardInfo struct {
	Group string
	Name  string
	Path  string
}

type Provider interface {
	AllAvailableWizards() (map[string][]WizardInfo, error)
	GetWizardInfo(group, wizard string) (WizardInfo, error)
	FindWizardGroup(wizard string) (string, error)
}

type WizardProvider struct {
	fs utils.FileSystem
	dp utils.DirectoryProvider
}

func NewWizardProvider(fs utils.FileSystem, dp utils.DirectoryProvider) *WizardProvider {
	return &WizardProvider{fs, dp}
}

func (p *WizardProvider) AllAvailableWizards() (map[string][]WizardInfo, error) {
	wizards := make(map[string][]WizardInfo)

	wizardsDir := p.dp.WizardsDirectory()
	allFilesAndDirectories, err := p.fs.ListDirectory(wizardsDir)
	if err != nil {
		return wizards, err
	}

	for _, fileinfo := range allFilesAndDirectories {
		if fileinfo.IsDir {
			basepath := filepath.Join(wizardsDir, fileinfo.Name)
			allFilesInWizardDir, err := p.fs.ListDirectory(basepath)
			if err == nil {
				for _, f := range allFilesInWizardDir {
					if isWizardFile(f.Name) {
						wizards = appendWizard(wizards, fileinfo.Name, f.Name, basepath)
					}
				}
			}
		} else {
			if isWizardFile(fileinfo.Name) {
				wizards = appendWizard(wizards, "none", fileinfo.Name, wizardsDir)
			}
		}
	}

	return wizards, nil
}

func (p *WizardProvider) GetWizardInfo(group, wizard string) (WizardInfo, error) {
	allWizards, err := p.AllAvailableWizards()
	if err != nil {
		return WizardInfo{}, err
	}

	if wizards, ok := allWizards[group]; ok {
		for _, w := range wizards {
			if w.Name == wizard {
				return w, nil
			}
		}

		return WizardInfo{}, errors.New(fmt.Sprintf("Unknown wizard", wizard, " in group", group))
	}

	return WizardInfo{}, errors.New(fmt.Sprintf("Unknown group", group))
}

func (p *WizardProvider) FindWizardGroup(wizard string) (string, error) {
	allWizards, err := p.AllAvailableWizards()
	if err != nil {
		return "", err
	}

	for group, wizards := range allWizards {
		for _, w := range wizards {
			if w.Name == wizard {
				return group, nil
			}
		}
	}

	return "", errors.New(fmt.Sprintf("Unknown wizard", wizard))
}

func appendWizard(wizards map[string][]WizardInfo, wizardGroup, wizardFile, basepath string) map[string][]WizardInfo {
	if _, ok := wizards[wizardGroup]; !ok {
		wizards[wizardGroup] = make([]WizardInfo, 0)
	}

	wizardName := wizardFile[0 : len(wizardFile)-len(utils.WIZARD_EXT)]
	wizard := WizardInfo{wizardGroup, wizardName, filepath.Join(basepath, wizardFile)}

	wizards[wizardGroup] = append(wizards[wizardGroup], wizard)
	return wizards
}

func isWizardFile(fname string) bool {
	return filepath.Ext(fname) == utils.WIZARD_EXT
}

/************************************
 * Mock
 ************************************/
type MockProvider struct {
	mock.Mock
}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (m *MockProvider) Run() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockProvider) AllAvailableWizards() (map[string][]WizardInfo, error) {
	args := m.Called()
	return args.Get(0).(map[string][]WizardInfo), args.Error(1)
}

func (m *MockProvider) GetWizardInfo(group, wizard string) (WizardInfo, error) {
	args := m.Called(group, wizard)
	return args.Get(0).(WizardInfo), args.Error(1)
}

func (m *MockProvider) FindWizardGroup(wizard string) (string, error) {
	args := m.Called(wizard)
	return args.String(0), args.Error(1)
}
