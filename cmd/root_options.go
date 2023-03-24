package cmd

import (
	"github.com/creasty/defaults"
	logger "github.com/sirupsen/logrus"
	"my-go-cobra-test/pkg/builder"
	"strconv"
)

type RootOptions struct {
	Num1	string		//`default:"3"`
	Num2	string		//`default:"4"`
}

func NewRootOptions() *RootOptions {
	rootOpts := &RootOptions{}
	if err := defaults.Set(rootOpts); err != nil {
		logger.WithError(err).WithField("options", "RootOptions").Fatal("error setting infelkit options defaults")
	}
	return rootOpts
}

func (ro *RootOptions) Log(){
	fields := logger.Fields{}
	if ro.Num1 != "" {
		fields["num1"] = ro.Num1
	}
	if ro.Num2 != "" {
		fields["num2"] = ro.Num2
	}

	logger.WithFields(fields).Debug("running with options")
}

func (ro *RootOptions) toBuild() *builder.Build {
	num1, err := strconv.Atoi(ro.Num1)
	if err != nil {
		logger.WithError(err).WithField("num1", ro.Num1).Fatal("num1 format error")
		//os.Exit(1)
	}
	num2, err := strconv.Atoi(ro.Num2)
	if err != nil {
		logger.WithError(err).WithField("num2", ro.Num2).Fatal("num2 format error")
		//os.Exit(1)
	}
	build := &builder.Build{
		Num1: num1,
		Num2: num2,
	}

	return build
}