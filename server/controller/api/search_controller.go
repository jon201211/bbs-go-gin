package api

import (
	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/model"
	"bbs-go/service"
	"bbs-go/util/es"

	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	base.BaseController
}

func (c *SearchController) AnyReindex(ctx *gin.Context) {
	go service.TopicService.ScanDesc(func(topics []model.Topic) {
		for _, t := range topics {
			topic := service.TopicService.Get(t.Id)
			es.UpdateTopicIndex(topic)
		}
	})
	c.JsonSuccess(ctx, nil)
}

func (c *SearchController) PostTopic(ctx *gin.Context) {
	var (
		page      = simple.FormValueIntDefault(ctx, "page", 1)
		keyword   = simple.FormValue(ctx, "keyword")
		nodeId    = simple.FormValueInt64Default(ctx, "nodeId", 0)
		timeRange = simple.FormValueIntDefault(ctx, "timeRange", 0)
	)

	docs, paging, err := es.SearchTopic(keyword, nodeId, timeRange, page, 20)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
	}

	items := render.BuildSearchTopics(docs)
	c.JsonPageData(ctx, items, paging)
}
