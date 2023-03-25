package api

import (
	"bbs-go/controller/base"
	"bbs-go/service"
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	base.BaseController
}

func (c *LikeController) GetIsLiked(ctx *gin.Context) {
	var (
		user           = service.UserTokenService.GetCurrent(ctx)
		entityType     = simple.FormValue(ctx, "entityType")
		entityIds      = simple.FormValueInt64Array(ctx, "entityIds")
		likedEntityIds []int64
	)
	if user != nil {
		likedEntityIds = service.UserLikeService.IsLiked(user.Id, entityType, entityIds)
	}
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("liked", likedEntityIds).Build())
}

func (c *LikeController) GetLiked(ctx *gin.Context) {
	var (
		user       = service.UserTokenService.GetCurrent(ctx)
		entityType = simple.FormValue(ctx, "entityType")
		entityId   = simple.FormValueInt64Default(ctx, "entityId", 0)
	)
	if user == nil || simple.IsBlank(entityType) || entityId <= 0 {
		c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("liked", false).Build())
		return
	} else {
		liked := service.UserLikeService.Exists(user.Id, entityType, entityId)
		c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("liked", liked).Build())
		return
	}
}
