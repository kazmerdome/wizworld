package elixir

import (
	"context"

	wizardworldapi "github.com/kazmerdome/wizworld/internal/actor/wizard-world-api"
)

type ElixirService interface {
	GetElixirsByIngredients(ctx context.Context, ingredients []string) ([]Elixir, error)
}

type elixirService struct {
	wizardApiClient wizardworldapi.WizardApiClient
}

func NewElixirService(wizardApiClient wizardworldapi.WizardApiClient) *elixirService {
	return &elixirService{wizardApiClient: wizardApiClient}
}

func (r *elixirService) GetElixirsByIngredients(ctx context.Context, ingredients []string) ([]Elixir, error) {
	ingredient := ""
	if len(ingredients) > 0 {
		ingredient = ingredients[0]
	}

	elixirList, err := r.wizardApiClient.ListElixirs(ctx, wizardworldapi.ListElixirsRequest{Ingredient: ingredient})
	if err != nil {
		return nil, err
	}

	result := []Elixir{}
	for _, elixirItem := range elixirList {
		matched := 0
		for _, elixirIngredient := range elixirItem.Ingredients {
			for _, ingredientParam := range ingredients {
				if ingredientParam == elixirIngredient.Name {
					matched++
				}
			}
		}
		if len(ingredients) < 2 || matched == len(ingredients) {
			result = append(result, Elixir{
				Id:         elixirItem.Id,
				Name:       elixirItem.Name,
				Difficulty: elixirItem.Difficulty,
			})
		}
	}

	return result, nil
}
