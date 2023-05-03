package ingredient

import (
	"log"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type IngredientCommand struct {
	service IngredientService
}

func NewIngredientCommand(service IngredientService) *IngredientCommand {
	return &IngredientCommand{service: service}
}

func (r *IngredientCommand) listIngredientsCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:     "ingredients",
		Short:   "Get a list of ingredients",
		Example: "wizworld ingredients",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Call elixir service
			ingredients, err := r.service.GetIngredients(cmd.Context())
			if err != nil {
				return err
			}

			// Display results in table
			headerFmt := color.New(color.FgHiMagenta, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgHiMagenta).SprintfFunc()
			ingredientsTbl := table.New("Name", "Id")
			ingredientsTbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, ingredient := range ingredients {
				ingredientsTbl.AddRow(ingredient.Name, ingredient.Id)
			}
			// elixirTbl.Print()
			ingredientsTbl.WithWriter(log.Writer())
			ingredientsTbl.Print()

			return nil
		},
	}
	return command
}

func (r *IngredientCommand) GetCommands() []*cobra.Command {
	return []*cobra.Command{
		r.listIngredientsCommand(),
	}
}
