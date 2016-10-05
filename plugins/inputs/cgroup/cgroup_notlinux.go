// +build !linux

package cgroup

import (
	"github.com/fyannk/telegraf"
)

func (g *CGroup) Gather(acc telegraf.Accumulator) error {
	return nil
}
