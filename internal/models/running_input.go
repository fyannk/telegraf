package models

import (
	"time"

	"github.com/fyannk/telegraf"
)

type RunningInput struct {
	Name   string
	Input  telegraf.Input
	Config *InputConfig
}

// InputConfig containing a name, interval, and filter
type InputConfig struct {
	Name              string
	NameOverride      string
	MeasurementPrefix string
	MeasurementSuffix string
	Tags              map[string]string
	Filter            Filter
	Interval          time.Duration
}
