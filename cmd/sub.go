package cmd

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"my-go-cobra-test/pkg/builder"
)

func NewSubCmd(rootOpts *RootOptions, rootFlags *pflag.FlagSet) *cobra.Command {
	dockerCmd := &cobra.Command{
		Use: "num-sub",
		Short: "execute subtraction: num1-num2",
		Aliases: []string{"sub"},
		Run: func(cmd *cobra.Command, args []string) {
			logger.WithField("processor", cmd.Name()).Info("will execute num1-num2")
			if !configOptions.DryRun {
				builder.StartSub(rootOpts.toBuild())
			}
		},
	}
	dockerCmd.PersistentFlags().AddFlagSet(rootFlags)

	return dockerCmd
}
