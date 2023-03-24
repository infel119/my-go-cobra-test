package cmd

import (
	"github.com/creasty/defaults"
	logger "github.com/sirupsen/logrus"
)

//var validArgs = []string{"num-add", "num-sub"}
//var aliasArgs = []string{"add", "sub"}

var configOptions *ConfigOptions

// ConfigOptions represent the persistent configuration flags of infelkit.
type ConfigOptions struct {
	LogLevel string		`default:"info" name:"log level"`
	DryRun	 bool
}

func NewConfigOptions() *ConfigOptions {
	o := &ConfigOptions{}
	if err := defaults.Set(o); err != nil {
		logger.WithError(err).WithField("options", "ConfigOPtions").Fatal("error setting infelkit options defaults")
	}
	return o
}