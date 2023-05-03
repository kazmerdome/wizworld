package ingredient

import (
	"context"

	wizardworldapi "github.com/kazmerdome/wizworld/internal/actor/wizard-world-api"
)

type IngredientService interface {
	GetIngredients(ctx context.Context) ([]Ingredient, error)
}

type ingredientService struct {
	wizardApiClient wizardworldapi.WizardApiClient
}

func NewIngredientService(wizardApiClient wizardworldapi.WizardApiClient) *ingredientService {
	return &ingredientService{wizardApiClient: wizardApiClient}
}

func (r *ingredientService) GetIngredients(ctx context.Context) ([]Ingredient, error) {
	ingredientList, err := r.wizardApiClient.ListIngredients(ctx)
	if err != nil {
		return nil, err
	}
	result := []Ingredient{}
	for _, ingredientItem := range ingredientList {
		result = append(result, Ingredient{
			Id:   ingredientItem.Id,
			Name: ingredientItem.Name,
		})
	}
	return result, nil
}
