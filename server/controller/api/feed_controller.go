package api

import (
	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/service"
	"bbs-go/util/simple"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FeedController struct {
	base.BaseController
}

func (c *FeedController) GetTopics(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	cursor := simple.FormValueInt64Default(ctx, "cursor", 0)
	topics, cursor, hasMore := service.UserFeedService.GetTopics(user.Id, cursor)
	c.JsonCursorData(ctx, render.BuildSimpleTopics(topics, user), strconv.FormatInt(cursor, 10), hasMore)
	return
}
