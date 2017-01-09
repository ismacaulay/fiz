package commands

import (
    "github.com/ismacaulay/fiz/wizards"
)

type WizardCommand struct {
    loader wizards.Loader
    commands []string
}

func NewWizardCommand(loader wizards.Loader, commands []string) *WizardCommand {
    return &WizardCommand{loader, commands}
}

func (c *WizardCommand) Run() error {
    wizard, err := c.loader.Load(c.commands)
    if err != nil {
        return err
    }

    wizard.Run()
    return nil
}
