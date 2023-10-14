package util

import (
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/robfig/cron/v3"
)

// MustAddFunc will panic
func MustAddFunc(c *cron.Cron, spec string, cmd func()) {
	_, err := c.AddFunc(spec, cmd)
	if err != nil {
		flog.Panic(err.Error())
	}
}
