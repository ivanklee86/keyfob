package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aoach.tech/keyfob/pkg/keyfob"
)

var (
	// Build information (injected by goreleaser).
	version = "dev"
)

const (
	defaultConfigFilename = "keyfob"
	envPrefix             = "KEYFOB"
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
		Use:     "keyfob",
		Short:   "Run commands with groups of env vars securely! 🔒🔒🔒",
		Long:    "Run commands with groups of env vars. Secrets stored securely in the OS keychain. 🔒",
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			keyfob.Out = cmd.OutOrStdout()
			keyfob.Err = cmd.ErrOrStderr()

			return initializeConfig(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(keyfob.Out, cmd.UsageString())
		},
	}

	cmd.PersistentFlags().BoolVar(&keyfob.NoExitCode, "no-exit-on-fail", false, "Don't return a non-zero exit code on failure.")

	return cmd
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	v.SetConfigName(defaultConfigFilename)
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()
	bindFlags(cmd, v)

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			if err := v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix)); err != nil {
				os.Exit(1)
			}
		}

		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			if err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val)); err != nil {
				os.Exit(1)
			}
		}
	})
}
