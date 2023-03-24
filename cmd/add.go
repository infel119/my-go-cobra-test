package cmd

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"my-go-cobra-test/pkg/builder"
)

func NewAddCmd(rootOpts *RootOptions, rootFlags *pflag.FlagSet) *cobra.Command {
	dockerCmd := &cobra.Command{
		Use: "num-add",
		Short: "execute addition: num1+num2",
		Aliases: []string{"add"},
		Run: func(cmd *cobra.Command, args []string) {
			logger.WithField("processor", cmd.Name()).Info("will execute num1+num2")
			if !configOptions.DryRun {
				builder.StartAdd(rootOpts.toBuild())
			}
		},
	}
	dockerCmd.PersistentFlags().AddFlagSet(rootFlags)

	return dockerCmd
}