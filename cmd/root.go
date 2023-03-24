package cmd

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	c *cobra.Command
}

func init(){
	logger.SetLevel(logger.InfoLevel) //默认"info"等级
	logger.SetFormatter(&logger.TextFormatter{
		DisableTimestamp: true,			//关闭时间戳
		ForceColors: true,				//颜色
		DisableLevelTruncation: false, //开启level显示
	})
}


func Start(){
	root := NewRootCmd()
	if err := root.c.Execute(); err != nil {
		logger.WithError(err).Fatal("error executing infelkit")
	}
}

//infelkit add/sub --num1 xxx --num2 xxx --loglevel [debug,info] --dryrun
func NewRootCmd() *RootCmd {
	configOptions = NewConfigOptions()
	rootOpts := NewRootOptions()

	rootCmd := &cobra.Command{
		Use: "infelkit",
		Short: "infelkit is a sample example of use go cobra pkg -- command line test",
		ValidArgs: []string{"num-add", "num-sub"},  //有效子args
		ArgAliases: []string{"add", "sub"},  //上面的别名
		Args: cobra.OnlyValidArgs,  //当输入的arg在ValidArgs才有效，否则报错
		DisableFlagsInUseLine: true,
		DisableAutoGenTag: true,
		SilenceUsage: true,			//参数有错时不自动打印帮助信息
		Version: "v1.0",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				logger.WithField("args", cmd.ValidArgs).Info("specify a valid arg")
			}
			//cmd.Help()
		},

	}

	ret := &RootCmd{
		c: rootCmd,
	}

	rootCmd.PersistentPreRunE = validatePreRunFunc(ret, rootOpts)  //在rootCmd.Run之前运行的函数

	flags := rootCmd.Flags()
	flags.StringVarP(&configOptions.LogLevel, "loglevel", "l", configOptions.LogLevel, "log level")
	flags.BoolVar(&configOptions.DryRun, "dryrun", configOptions.DryRun, "do not actually perform the action")

	flags.StringVarP(&rootOpts.Num1, "num1", "a", rootOpts.Num1, "the first number")
	flags.StringVarP(&rootOpts.Num2, "num2", "b", rootOpts.Num2, "the second number")


	//--loglevel 的命令补全
	rootCmd.RegisterFlagCompletionFunc("loglevel", func(c *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"debug", "info"}, cobra.ShellCompDirectiveDefault
	})

	rootCmd.AddCommand(NewAddCmd(rootOpts,flags))
	rootCmd.AddCommand(NewSubCmd(rootOpts, flags))

	return ret
}

func validatePreRunFunc(rootCommand *RootCmd, rootOpts *RootOptions) func(c *cobra.Command, args []string) error {
	return func(c *cobra.Command, args []string) error {

		if c.Root() != c && c.Name() != "help" {
			/*if errs := rootOpts.Validate(); errs != nil {
				for _, err := range errs {
					logger.WithError(err).Error("error validating options")
				}
				return fmt.Errorf("exiting for validation errors")
			}*/
			rootOpts.Log()
		}

		if logger.GetLevel().String() != configOptions.LogLevel {
			lvl, err := logger.ParseLevel(configOptions.LogLevel)
			if err != nil {
				logger.WithError(err).WithField("loglevel", configOptions.LogLevel).Error("loglevel format error")
				return err
			}
			logger.SetLevel(lvl)
		}

		return nil
	}
}