package wizardworldapi

import (
	"context"
)

//go:generate make name=WizardApiClient mock

type WizardApiClient interface {
	ListElixirs(ctx context.Context, params ListElixirsRequest) ([]ElixirResponse, error)
	ListIngredients(ctx context.Context) ([]IngredientResponse, error)
}
