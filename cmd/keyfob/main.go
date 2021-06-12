package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.aoach.tech/keyfob/pkg/keyfob"
)


var (
	// Build information (injected by goreleaser).
	version = "dev"
)

// main function.
func main() {
	command := NewRootCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func NewRootCommand() *cobra.Command {
	keyfob := keyfob.New()

	cmd := &cobra.Command{
		Use: "keyfob",
		Short: "Run commands with groups of env vars securely! 🔒🔒🔒",
		Long: "Run commands with groups of env vars. Secrets stored securely in the OS keychain. 🔒",
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			keyfob.Out = cmd.OutOrStdout()
			keyfob.Err = cmd.ErrOrStderr()

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(keyfob.Out, cmd.UsageString())
		},
	}

	return cmd	
}
