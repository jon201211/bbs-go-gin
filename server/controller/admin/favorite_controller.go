package admin

import (
	"strconv"

	"bbs-go/controller/base"
	"bbs-go/util/simple"

	"bbs-go/model"
	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type FavoriteController struct {
	base.BaseController
}

func (c *FavoriteController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.FavoriteService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *FavoriteController) AnyList(ctx *gin.Context) {
	list, paging := service.FavoriteService.FindPageByParams(simple.NewQueryParams(ctx).PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}

func (c *FavoriteController) PostCreate(ctx *gin.Context) {
	t := &model.Favorite{}
	simple.ReadForm(ctx, t)

	err := service.FavoriteService.Create(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *FavoriteController) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t := service.FavoriteService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	simple.ReadForm(ctx, t)

	err = service.FavoriteService.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}
