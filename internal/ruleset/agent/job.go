package agent

import (
	"github.com/allegro/bigcache/v3"
	"github.com/flowline-io/flowkit/internal/pkg/client"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/util"
	"github.com/flowline-io/flowkit/internal/ruleset/agent/bot"
	"github.com/robfig/cron/v3"
)

type agentJob struct {
	cache  *bigcache.BigCache
	client *client.Flowbot
}

func (j *agentJob) RunAnki(c *cron.Cron) {
	util.MustAddFunc(c, "0 * * * * *", func() {
		flog.Info("[agent] anki stats")
		bot.AnkiStats(j.client)
	})
	util.MustAddFunc(c, "0 * * * * *", func() {
		flog.Info("[agent] anki review")
		bot.AnkiReview(j.client)
	})
}

func (j *agentJob) RunDev(c *cron.Cron) {
	util.MustAddFunc(c, "0 * * * * *", func() {
		flog.Info("[agent] dev import")
		bot.DevImport(j.client)
	})
}
