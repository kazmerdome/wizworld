package elixir

import (
	wizardworldapi "github.com/kazmerdome/wizworld/internal/actor/wizard-world-api"
	"github.com/spf13/cobra"
)

//go:generate make name=ElixirService mock

type elixir struct {
	command *ElixirCommand
}

func NewElixir(wizardApiClient wizardworldapi.WizardApiClient) *elixir {
	service := NewElixirService(wizardApiClient)
	command := NewElixirCommand(service)
	return &elixir{command: command}
}

// implements cli.CobraModule
func (r *elixir) GetCommands() []*cobra.Command {
	return r.command.GetCommands()
}
