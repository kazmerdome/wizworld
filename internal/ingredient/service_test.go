package ingredient_test

import (
	"context"
	"fmt"
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	wizardworldapi "github.com/kazmerdome/wizworld/internal/actor/wizard-world-api"
	"github.com/kazmerdome/wizworld/internal/ingredient"
	"github.com/kazmerdome/wizworld/mocks"
	"github.com/stretchr/testify/assert"
)

type serviceFixture struct {
	service ingredient.IngredientService
	mocks   struct {
		wizardApiClient *mocks.WizardApiClient
	}
}

func newServiceFixture() *serviceFixture {
	f := new(serviceFixture)
	f.mocks.wizardApiClient = &mocks.WizardApiClient{}
	f.service = ingredient.NewIngredientService(f.mocks.wizardApiClient)
	return f
}

// GetIngredients
func TestGetIngredients_FailsOn_ListIngredients(t *testing.T) {
	ctx := context.Background()
	f := newServiceFixture()
	f.mocks.wizardApiClient.On("ListIngredients", ctx).
		Return([]wizardworldapi.IngredientResponse{}, fmt.Errorf("some error"))
	ingredients, err := f.service.GetIngredients(ctx)

	assert.EqualError(t, err, "some error")
	assert.Empty(t, ingredients)
}

func TestGetIngredients_Success(t *testing.T) {
	ctx := context.Background()
	f := newServiceFixture()
	mockedResponse := []wizardworldapi.IngredientResponse{{Id: faker.UUID(), Name: faker.BeerHop()}}
	f.mocks.wizardApiClient.On("ListIngredients", ctx).
		Return(mockedResponse, nil)
	ingredients, err := f.service.GetIngredients(ctx)

	assert.NoError(t, err)
	assert.Len(t, ingredients, 1)
	assert.Equal(t, ingredients[0].Id, mockedResponse[0].Id)
	assert.Equal(t, ingredients[0].Name, mockedResponse[0].Name)
}
