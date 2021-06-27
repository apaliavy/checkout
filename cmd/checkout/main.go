package main

import "github.com/sirupsen/logrus"

func main() {
	logger := logrus.New()
	logger.Info("check app is up and running")
}
