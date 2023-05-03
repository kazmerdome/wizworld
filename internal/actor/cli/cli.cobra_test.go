package cli_test

import (
	"testing"

	"github.com/kazmerdome/wizworld/internal/actor/cli"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type mockedModule struct{}

func (r mockedModule) GetCommands() []*cobra.Command {
	return []*cobra.Command{}
}

func TestCobraCli(t *testing.T) {
	cc := cli.NewCobraCli(&cli.CobraCliOptions{})
	assert.Implements(t, (*cli.Cli)(nil), cc)

	err := cc.Execute()
	assert.NoError(t, err)

	mm := mockedModule{}
	cc.ExposeCommands(mm)
	assert.Len(t, mm.GetCommands(), 0)
}
