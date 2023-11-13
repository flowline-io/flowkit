package agent

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
	// agent job
	if setting.DefaultConfig().AccessToken != "" {
		cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(24*time.Hour))
		if err != nil {
			flog.Panic(err.Error())
		}
		job := &agentJob{cache: cache, client: client.NewTinode(setting.DefaultConfig().AccessToken)}
		job.RunClipboard(c)
		job.RunAnki(c)
		job.RunDev(c)
	}
	c.Start()
}
