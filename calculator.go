package main

import "github.com/Sirupsen/logrus"

type Calculator struct {
	Log       *logrus.Logger
	Collector *Collector
}
