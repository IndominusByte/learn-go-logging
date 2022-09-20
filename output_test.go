package main

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logger := logrus.New()
	logger.SetOutput(file)

	logger.Info("Hello info")
	logger.Warn("Hello warn")
	logger.Error("Hello error")
}
