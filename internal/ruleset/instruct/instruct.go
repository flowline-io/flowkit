package instruct

import (
	"context"
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/flowline-io/linkit/internal/pkg/client"
	"github.com/flowline-io/linkit/internal/pkg/logs"
	"github.com/flowline-io/linkit/internal/pkg/setting"
	"github.com/robfig/cron/v3"
	"time"
)

func Cron(app fyne.App, window fyne.Window) {
	c := cron.New(cron.WithSeconds())
	// instruct job
	if setting.Get().AccessToken != "" {
		cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(time.Hour))
		if err != nil {
			logs.Panic(err.Error())
		}
		job := &instructJob{app: app, window: window, client: client.NewTinode(setting.Get().AccessToken), cache: cache}
		_, err = c.AddJob("*/10 * * * * *", job)
		if err != nil {
			logs.Panic(err.Error())
		}
	}
	c.Start()
}
