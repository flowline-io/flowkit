package instruct

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/flowline-io/flowkit/internal/pkg/client"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/setting"
	"github.com/robfig/cron/v3"
	"time"
)

func Cron() {
	c := cron.New(cron.WithSeconds())
	// instruct job
	if setting.DefaultConfig().AccessToken != "" {
		cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(time.Hour))
		if err != nil {
			flog.Panic(err.Error())
		}
		job := &instructJob{client: client.NewFlowbot(setting.DefaultConfig().AccessToken), cache: cache}
		_, err = c.AddJob("*/10 * * * * *", job)
		if err != nil {
			flog.Panic(err.Error())
		}
	}
	c.Start()
}
