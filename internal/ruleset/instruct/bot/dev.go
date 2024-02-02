package bot

import (
	"time"

	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/types"
)

var dev = []Executor{
	{
		Flag: "dev_example",
		Run: func(app any, window any, data types.KV) error {
			flog.Info("dev example %s %s", data, time.Now())
			return nil
		},
	},
}
