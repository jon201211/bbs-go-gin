package admin

import (
	"bbs-go/controller/base"
	"bbs-go/model"
	"bbs-go/service"
	"bbs-go/util/simple"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmailCodeController struct {
	base.BaseController
}

func (c *EmailCodeController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.EmailCodeService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *EmailCodeController) AnyList(ctx *gin.Context) {
	list, paging := service.EmailCodeService.FindPageByParams(simple.NewQueryParams(ctx).PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}

func (c *EmailCodeController) PostCreate(ctx *gin.Context) {
	t := &model.EmailCode{}
	err := simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.EmailCodeService.Create(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *EmailCodeController) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t := service.EmailCodeService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	err = simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.EmailCodeService.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}
