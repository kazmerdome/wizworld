package ingredient_test

import (
	"testing"

	"github.com/kazmerdome/wizworld/internal/actor/cli"
	"github.com/kazmerdome/wizworld/internal/ingredient"
	"github.com/kazmerdome/wizworld/mocks"
	"github.com/stretchr/testify/assert"
)

func TestIngredient(t *testing.T) {
	e := ingredient.NewIngredient(&mocks.WizardApiClient{})
	assert.Implements(t, (*cli.CobraModule)(nil), e)
	assert.Equal(t, len(e.GetCommands()), 1)
}
