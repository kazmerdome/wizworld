package elixir

import (
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type ElixirCommand struct {
	service ElixirService
}

func NewElixirCommand(service ElixirService) *ElixirCommand {
	return &ElixirCommand{service: service}
}

func (r *ElixirCommand) addElixirCommand() *cobra.Command {
	var ingredientsFlagValue string
	var command = &cobra.Command{
		Use:     "elixirs",
		Short:   "Get a list of elixirs by ingredients",
		Example: "wizworld elixirs -i \"Neem oil, Jewelweed\" ",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get ingredient list
			ingredients := strings.Split(ingredientsFlagValue, ",")
			for i := 0; i < len(ingredients); i++ {
				ingredients[i] = strings.TrimSpace(ingredients[i])
			}

			// Call elixir service
			elixirs, err := r.service.GetElixirsByIngredients(cmd.Context(), ingredients)
			if err != nil {
				return err
			}

			// Handle not found/empty list
			if len(elixirs) < 1 {
				color.Red("no elixirs found with the requested ingredient(s): %s.", ingredients)
				color.Black(`If you are not sure about the correct name of the ingredient,
use "wizworld ingredients" command to get a list of ingredients.
Note that the names of the ingredients are case sensitive!
				`)
			}

			// Display results in table
			if len(elixirs) > 0 {
				headerFmt := color.New(color.FgHiMagenta, color.Underline).SprintfFunc()
				columnFmt := color.New(color.FgHiMagenta).SprintfFunc()
				elixirTbl := table.New("Name", "Difficulty", "Id")
				elixirTbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
				for _, elixir := range elixirs {
					elixirTbl.AddRow(elixir.Name, elixir.Difficulty, elixir.Id)
				}
				elixirTbl.WithWriter(log.Writer())
				elixirTbl.Print()
			}
			return nil
		},
	}

	command.Flags().StringVarP(&ingredientsFlagValue, "ingredients", "i", "", "Single or multiple ingredients, separated by ,")
	return command
}

func (r *ElixirCommand) GetCommands() []*cobra.Command {
	return []*cobra.Command{
		r.addElixirCommand(),
	}
}
