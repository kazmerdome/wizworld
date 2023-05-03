package cli

import "github.com/spf13/cobra"

type CobraModule interface {
	GetCommands() []*cobra.Command
}

type Cli interface {
	Execute() error
	ExposeCommands(CobraModule)
}
