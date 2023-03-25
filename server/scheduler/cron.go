package scheduler

import (
	"bbs-go/util/sitemap"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"

	"bbs-go/service"
)

func Start() {
	c := cron.New()

	// Generate RSS
	addCronFunc(c, "@every 30m", func() {
		service.ArticleService.GenerateRss()
		service.TopicService.GenerateRss()
		service.ProjectService.GenerateRss()
	})

	// Generate sitemap
	addCronFunc(c, "0 0 4 ? * *", func() {
		sitemap.Generate()
	})

	c.Start()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}
