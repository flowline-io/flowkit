package bot

import (
	"fyne.io/fyne/v2"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/types"
	"time"
)

var dev = []Executor{
	{
		Flag: "dev_example",
		Run: func(app fyne.App, window fyne.Window, data types.KV) error {
			flog.Info("dev example %s %s", data, time.Now())
			return nil
		},
	},
}
