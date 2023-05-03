package ingredient_test

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/kazmerdome/wizworld/internal/ingredient"
	"github.com/kazmerdome/wizworld/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type commandFixture struct {
	command *ingredient.IngredientCommand
	rootCmd *cobra.Command
	ctx     context.Context
	mocks   struct {
		service *mocks.IngredientService
	}
	getOut func() string
}

func newCommandFixture() *commandFixture {
	f := new(commandFixture)
	f.rootCmd = &cobra.Command{}
	f.ctx = context.TODO()
	buf := new(bytes.Buffer)
	f.rootCmd.SetOut(buf)
	f.rootCmd.SetErr(buf)
	f.rootCmd.SetContext(f.ctx)
	f.getOut = func() string {
		return strings.TrimSpace(buf.String())
	}
	f.mocks.service = &mocks.IngredientService{}
	f.command = ingredient.NewIngredientCommand(f.mocks.service)
	return f
}

func TestGetIngredientsCommand_FailsOn_ServiceCall(t *testing.T) {
	f := newCommandFixture()
	f.rootCmd.AddCommand(f.command.GetCommands()...)
	f.rootCmd.SetArgs([]string{"ingredients"})
	f.mocks.service.On("GetIngredients", f.ctx).
		Return([]ingredient.Ingredient{}, fmt.Errorf("some error"))

	err := f.rootCmd.Execute()
	assert.Contains(t, f.getOut(), "some error")
	assert.EqualError(t, err, "some error")
}

func TestGetIngredientsCommand_Success(t *testing.T) {
	f := newCommandFixture()
	f.rootCmd.AddCommand(f.command.GetCommands()...)
	f.rootCmd.SetArgs([]string{"ingredients"})
	f.mocks.service.On("GetIngredients", f.ctx).
		Return([]ingredient.Ingredient{{Id: "id", Name: "name"}}, nil)

	err := f.rootCmd.Execute()
	assert.NoError(t, err, "some error")
}
