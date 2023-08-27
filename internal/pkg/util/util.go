package util

import (
	"fyne.io/fyne/v2"
	"github.com/flowline-io/linkit/internal/pkg/logs"
	"github.com/robfig/cron/v3"
)

// WindowSizeToDialog scales the window size to a suitable dialog size.
func WindowSizeToDialog(s fyne.Size) fyne.Size {
	return fyne.NewSize(s.Width*0.8, s.Height*0.8)
}

// MustAddFunc will panic
func MustAddFunc(c *cron.Cron, spec string, cmd func()) {
	_, err := c.AddFunc(spec, cmd)
	if err != nil {
		logs.Panic(err.Error())
	}
}
