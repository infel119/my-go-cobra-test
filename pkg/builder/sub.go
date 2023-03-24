package builder

import logger "github.com/sirupsen/logrus"

func StartSub(buildOpts *Build){
	logger.WithField("num1", buildOpts.Num1).Debug("number1")
	logger.WithField("num2", buildOpts.Num2).Debug("number2")

	res := buildOpts.Num1 - buildOpts.Num2

	logger.WithField("result", res).Infof("%d-%d=%d", buildOpts.Num1, buildOpts.Num2, res)
}
