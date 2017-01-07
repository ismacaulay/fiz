package wizards

import (
    "github.com/ismacaulay/fiz/utils"

    "path/filepath"
)

type Provider interface {
	AllAvailableWizards() (map[string][]string, error)
}

type WizardProvider struct {
    fs utils.FileSystem
    dp utils.DirectoryProvider
}

func NewWizardProvider(fs utils.FileSystem, dp utils.DirectoryProvider) *WizardProvider {
	return &WizardProvider{fs, dp}
}

func (p *WizardProvider) AllAvailableWizards() (map[string][]string, error) {
	wizards := make(map[string][]string)

    wizardsDir := p.dp.WizardsDirectory()
    allFilesAndDirectories, err := p.fs.ListDirectory(wizardsDir)
    if err != nil {
        return wizards, err
    }

    for _, fileinfo := range allFilesAndDirectories {
        if fileinfo.IsDir {
            allFilesInWizardDir, err := p.fs.ListDirectory(filepath.Join(wizardsDir, fileinfo.Name))
            if err == nil {
                for _, f := range allFilesInWizardDir {
                    if isWizardFile(f.Name){
                        wizards = appendWizard(wizards, fileinfo.Name, f.Name)
                    }
                }
            }
        } else {
            if isWizardFile(fileinfo.Name) {
                wizards = appendWizard(wizards, "default", fileinfo.Name)
            }
        }
    }

	return wizards, nil
}

func appendWizard(wizards map[string][]string, wizardGroup, wizardFile string) map[string][]string {
    if _, ok := wizards[wizardGroup]; !ok {
        wizards[wizardGroup] = make([]string, 0)
    }

    wizards[wizardGroup] = append(wizards["default"], wizardFile[0:len(wizardFile) - len(".wizard")])
    return wizards
}

func isWizardFile(fname string) bool {
    return filepath.Ext(fname) == ".wizard"
}
