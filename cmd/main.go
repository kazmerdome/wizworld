package main

import (
	"os"

	"github.com/kazmerdome/wizworld/internal/actor/cli"
	wizardworldapi "github.com/kazmerdome/wizworld/internal/actor/wizard-world-api"
	"github.com/kazmerdome/wizworld/internal/elixir"
	"github.com/kazmerdome/wizworld/internal/ingredient"
)

func main() {
	// Init cli framework
	c := cli.NewCobraCli(&cli.CobraCliOptions{
		RootShortDescription: "Wizworld is a tool for interacting with the world of wizards",
		RootLongDescription:  `You can see the available commands listed below`,
	})

	// Init Wizard Api Client
	wizardClient := wizardworldapi.NewHttpClient(5, "https://wizard-world-api.herokuapp.com")

	// Add subcommands and init modules
	c.ExposeCommands(elixir.NewElixir(wizardClient))
	c.ExposeCommands(ingredient.NewIngredient(wizardClient))

	// Execute Cli
	err := c.Execute()
	if err != nil {
		os.Exit(1)
	}
}
