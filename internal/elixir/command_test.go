package elixir_test

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/kazmerdome/wizworld/internal/elixir"
	"github.com/kazmerdome/wizworld/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type commandFixture struct {
	command *elixir.ElixirCommand
	rootCmd *cobra.Command
	ctx     context.Context
	mocks   struct {
		service *mocks.ElixirService
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
	f.mocks.service = &mocks.ElixirService{}
	f.command = elixir.NewElixirCommand(f.mocks.service)
	return f
}

func TestGetElixirCommand_FailsOn_ServiceCall(t *testing.T) {
	f := newCommandFixture()
	f.rootCmd.AddCommand(f.command.GetCommands()...)
	f.rootCmd.SetArgs([]string{"elixirs", "-i", "test"})
	f.mocks.service.On("GetElixirsByIngredients", f.ctx, []string{"test"}).
		Return([]elixir.Elixir{}, fmt.Errorf("some error"))

	err := f.rootCmd.Execute()
	assert.Contains(t, f.getOut(), "some error")
	assert.EqualError(t, err, "some error")
}

func TestGetElixirCommand_When_NoElixirsFound(t *testing.T) {
	f := newCommandFixture()
	f.rootCmd.AddCommand(f.command.GetCommands()...)
	f.rootCmd.SetArgs([]string{"elixirs", "-i", "test"})
	f.mocks.service.On("GetElixirsByIngredients", f.ctx, []string{"test"}).
		Return([]elixir.Elixir{}, nil)

	err := f.rootCmd.Execute()
	assert.NoError(t, err)
}

func TestGetElixirCommand_When_ElixirFound(t *testing.T) {
	f := newCommandFixture()
	f.rootCmd.AddCommand(f.command.GetCommands()...)
	f.rootCmd.SetArgs([]string{"elixirs", "-i", "test"})
	f.mocks.service.On("GetElixirsByIngredients", f.ctx, []string{"test"}).
		Return([]elixir.Elixir{{
			Id:         faker.UUID(),
			Name:       faker.BeerHop(),
			Difficulty: faker.CarModel(),
		}}, nil)

	err := f.rootCmd.Execute()
	assert.NoError(t, err)
}
