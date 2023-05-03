package ingredient

import (
	wizardworldapi "github.com/kazmerdome/wizworld/internal/actor/wizard-world-api"
	"github.com/spf13/cobra"
)

//go:generate make name=IngredientService mock

type ingredient struct {
	command *IngredientCommand
}

func NewIngredient(wizardApiClient wizardworldapi.WizardApiClient) *ingredient {
	service := NewIngredientService(wizardApiClient)
	command := NewIngredientCommand(service)
	return &ingredient{command: command}
}

// implements cli.CobraModule
func (r *ingredient) GetCommands() []*cobra.Command {
	return r.command.GetCommands()
}
