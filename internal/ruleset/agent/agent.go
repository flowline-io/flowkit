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
	if setting.AppConfig().AccessToken == "" {
		return
	}
	// agent job
	c := cron.New(cron.WithSeconds())
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(24*time.Hour))
	if err != nil {
		flog.Panic(err.Error())
	}
	job := &agentJob{cache: cache, client: client.NewFlowbot(setting.AppConfig().AccessToken)}
	job.RunClipboard(c)
	job.RunAnki(c)
	job.RunDev(c)
	c.Start()
}
