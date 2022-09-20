package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("user", "oman").WithField("pass", "123").Info("login")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"user": "oman",
		"pass": "123",
	}).Info("login")
}
