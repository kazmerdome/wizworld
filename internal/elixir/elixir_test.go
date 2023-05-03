package elixir_test

import (
	"testing"

	"github.com/kazmerdome/wizworld/internal/actor/cli"
	"github.com/kazmerdome/wizworld/internal/elixir"
	"github.com/kazmerdome/wizworld/mocks"
	"github.com/stretchr/testify/assert"
)

func TestElixir(t *testing.T) {
	e := elixir.NewElixir(&mocks.WizardApiClient{})
	assert.Implements(t, (*cli.CobraModule)(nil), e)
	assert.Equal(t, len(e.GetCommands()), 1)
}
