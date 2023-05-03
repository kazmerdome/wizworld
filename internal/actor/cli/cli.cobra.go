package cli

import "github.com/spf13/cobra"

type cobraCli struct {
	rootCmd *cobra.Command
}

type CobraCliOptions struct {
	RootCommand          string
	RootShortDescription string
	RootLongDescription  string
}

func NewCobraCli(opts *CobraCliOptions) *cobraCli {
	rootCmd := &cobra.Command{
		Use:   opts.RootCommand,
		Short: opts.RootShortDescription,
		Long:  opts.RootLongDescription,
	}
	return &cobraCli{
		rootCmd: rootCmd,
	}
}

func (r *cobraCli) Execute() error {
	return r.rootCmd.Execute()
}

func (r *cobraCli) ExposeCommands(m CobraModule) {
	r.rootCmd.AddCommand(m.GetCommands()...)
}
