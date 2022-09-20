package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
}

func TestLevel_2(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
}
