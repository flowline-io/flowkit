package agent

import (
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/flowline-io/flowkit/internal/pkg/client"
	"github.com/flowline-io/flowkit/internal/pkg/logs"
	"github.com/flowline-io/flowkit/internal/pkg/util"
	"github.com/flowline-io/flowkit/internal/ruleset/agent/bot"
	"github.com/robfig/cron/v3"
)

type agentJob struct {
	app    fyne.App
	window fyne.Window
	cache  *bigcache.BigCache
	client *client.Tinode
}

func (j *agentJob) RunAnki(c *cron.Cron) {
	util.MustAddFunc(c, "0 * * * * *", func() {
		logs.Info("[agent] anki stats")
		bot.AnkiStats(j.client)
	})
	util.MustAddFunc(c, "0 * * * * *", func() {
		logs.Info("[agent] anki review")
		bot.AnkiReview(j.client)
	})
}

func (j *agentJob) RunClipboard(c *cron.Cron) {
	util.MustAddFunc(c, "*/10 * * * * *", func() {
		logs.Info("[agent] clipboard upload")
		bot.ClipboardUpload(j.window, j.cache, j.client)
	})
}

func (j *agentJob) RunDev(c *cron.Cron) {
	util.MustAddFunc(c, "0 * * * * *", func() {
		logs.Info("[agent] dev import")
		bot.DevImport(j.client)
	})
}
