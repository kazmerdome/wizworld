package elixir_test

import (
	"context"
	"fmt"
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	wizardworldapi "github.com/kazmerdome/wizworld/internal/actor/wizard-world-api"
	"github.com/kazmerdome/wizworld/internal/elixir"
	"github.com/kazmerdome/wizworld/mocks"
	"github.com/stretchr/testify/assert"
)

type serviceFixture struct {
	service elixir.ElixirService
	mocks   struct {
		wizardApiClient *mocks.WizardApiClient
	}
}

func newServiceFixture() *serviceFixture {
	f := new(serviceFixture)
	f.mocks.wizardApiClient = &mocks.WizardApiClient{}
	f.service = elixir.NewElixirService(f.mocks.wizardApiClient)
	return f
}

// GetElixirsByIngredients
//

func TestGetElixirsByIngredients_FailsOn_ListElixirs(t *testing.T) {
	ingredient := []string{faker.BeerHop()}
	ctx := context.Background()
	f := newServiceFixture()
	f.mocks.wizardApiClient.On("ListElixirs", ctx, wizardworldapi.ListElixirsRequest{Ingredient: ingredient[0]}).
		Return([]wizardworldapi.ElixirResponse{}, fmt.Errorf("some error"))
	elixirs, err := f.service.GetElixirsByIngredients(ctx, ingredient)

	assert.EqualError(t, err, "some error")
	assert.Empty(t, elixirs)
}

func TestGetElixirsByIngredients_When_ParameterIsASingleIngredient(t *testing.T) {
	ingredient := []string{faker.BeerHop()}
	ctx := context.Background()
	f := newServiceFixture()
	mockedResponse := []wizardworldapi.ElixirResponse{
		{Id: faker.UUID(), Name: faker.BeerHop(), Difficulty: faker.Sentence(10)},
		{Id: faker.UUID(), Name: faker.BeerHop(), Difficulty: faker.Sentence(10)},
	}
	f.mocks.wizardApiClient.On("ListElixirs", ctx, wizardworldapi.ListElixirsRequest{Ingredient: ingredient[0]}).
		Return(mockedResponse, nil)
	elixirs, err := f.service.GetElixirsByIngredients(ctx, ingredient)

	assert.NoError(t, err)
	assert.Equal(t, len(elixirs), 2)
	for i, elixir := range elixirs {
		assert.Equal(t, elixir.Id, mockedResponse[i].Id)
		assert.Equal(t, elixir.Name, mockedResponse[i].Name)
		assert.Equal(t, elixir.Difficulty, mockedResponse[i].Difficulty)
	}
}

func TestGetElixirsByIngredients_When_MultiIngredients(t *testing.T) {
	ingredients := []string{faker.BeerHop(), faker.BeerHop()}
	ctx := context.Background()
	f := newServiceFixture()
	mockedResponse := []wizardworldapi.ElixirResponse{
		{
			Id:              faker.UUID(),
			Name:            faker.BeerHop(),
			Characteristics: faker.Sentence(10),
			Ingredients: []wizardworldapi.IngredientResponse{
				{
					Id:   faker.UUID(),
					Name: ingredients[0],
				},
			},
		},
		{
			Id:              faker.UUID(),
			Name:            faker.BeerHop(),
			Characteristics: faker.Sentence(10),
			Ingredients: []wizardworldapi.IngredientResponse{
				{
					Id:   faker.UUID(),
					Name: ingredients[0],
				},
				{
					Id:   faker.UUID(),
					Name: ingredients[1],
				},
			},
		},
	}
	f.mocks.wizardApiClient.On("ListElixirs", ctx, wizardworldapi.ListElixirsRequest{Ingredient: ingredients[0]}).
		Return(mockedResponse, nil)
	elixirs, err := f.service.GetElixirsByIngredients(ctx, ingredients)

	assert.NoError(t, err)
	assert.Equal(t, len(elixirs), 1)
	assert.Equal(t, elixirs[0].Id, mockedResponse[1].Id)
	assert.Equal(t, elixirs[0].Name, mockedResponse[1].Name)
	assert.Equal(t, elixirs[0].Difficulty, mockedResponse[1].Difficulty)
}
