package api

import (
	"bbs-go/controller/base"
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"

	"bbs-go/service"
)

type FavoriteController struct {
	base.BaseController
}

// 是否收藏了
func (c *FavoriteController) GetFavorited(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	entityType := simple.FormValue(ctx, "entityType")
	entityId := simple.FormValueInt64Default(ctx, "entityId", 0)
	if user == nil || len(entityType) == 0 || entityId <= 0 {
		c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("favorited", false).Build())
		return
	} else {
		tmp := service.FavoriteService.GetBy(user.Id, entityType, entityId)
		c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("favorited", tmp != nil).Build())
		return
	}
}

// 取消收藏
func (c *FavoriteController) GetDelete(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	entityType := simple.FormValue(ctx, "entityType")
	entityId := simple.FormValueInt64Default(ctx, "entityId", 0)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	tmp := service.FavoriteService.GetBy(user.Id, entityType, entityId)
	if tmp != nil {
		service.FavoriteService.Delete(tmp.Id)
	}
	c.JsonSuccess(ctx, nil)
	return
}
