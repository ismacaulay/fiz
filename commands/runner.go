package commands

import (
	"fmt"
)

type Runner interface {
	Run(command string)
}

type CommandRunner struct {
}

func NewCommandRunner() *CommandRunner {
	return &CommandRunner{}
}

func (r *CommandRunner) Run(command string) {
	fmt.Println("running", command)
}
