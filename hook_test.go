package main

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type GlobalHook struct {
}

func (h *GlobalHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (h *GlobalHook) Fire(e *logrus.Entry) error {
	fmt.Println("Fire", e.Level, e.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&GlobalHook{})

	logger.Info("Hello info")
	logger.Error("Hello error")
}
