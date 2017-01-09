package wizards

import (
    "github.com/ismacaulay/fiz/utils"

    "path/filepath"
    "errors"
    "fmt"
)

type WizardInfo struct {
    Group string
    Name string
    Path string
}

type Provider interface {
	AllAvailableWizards() (map[string][]WizardInfo, error)
    GetWizardInfo(group, name string) (WizardInfo, error)
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
                    if isWizardFile(f.Name){
                        wizards = appendWizard(wizards, fileinfo.Name, f.Name, basepath)
                    }
                }
            }
        } else {
            if isWizardFile(fileinfo.Name) {
                wizards = appendWizard(wizards, "default", fileinfo.Name, wizardsDir)
            }
        }
    }

	return wizards, nil
}

func (p *WizardProvider) GetWizardInfo(group, name string) (WizardInfo, error) {
    allWizards, err := p.AllAvailableWizards()
    if err != nil {
        return WizardInfo{}, err
    }

    if wizards, ok := allWizards[group]; ok {
        for _, w := range wizards {
            if w.Name == name {
                return w, nil
            }
        }

        return WizardInfo{}, errors.New(fmt.Sprintf("Unknown wizard", name, " in group", group))
    }

    return WizardInfo{}, errors.New(fmt.Sprintf("Unknown group", group))
}

func appendWizard(wizards map[string][]WizardInfo, wizardGroup, wizardFile, basepath string) map[string][]WizardInfo {
    if _, ok := wizards[wizardGroup]; !ok {
        wizards[wizardGroup] = make([]WizardInfo, 0)
    }

    wizardName := wizardFile[0:len(wizardFile) - len(".wizard")]
    wizard := WizardInfo{wizardGroup, wizardName, filepath.Join(basepath, wizardFile)}

    wizards[wizardGroup] = append(wizards[wizardGroup], wizard)
    return wizards
}

func isWizardFile(fname string) bool {
    return filepath.Ext(fname) == ".wizard"
}

