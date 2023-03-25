package base

import (
	"bbs-go/controller/render"
	"bbs-go/service"
	"bbs-go/util/simple"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"bbs-go/model"
)

// BaseController controller
type BaseController struct {
}

// BindAndValidate bind and validate
func (c *BaseController) BindAndValidate(ctx *gin.Context, obj interface{}) bool {
	if err := simple.ReadForm(ctx, obj); err != nil {
		c.JsonFail(ctx, &simple.CodeError{Code: -1, Message: err.Error()})
		return false
	}
	return true
}

// GetCurrentUser get current user from context
func (c *BaseController) GetCurrentUser(ctx *gin.Context) *model.User {
	if currentUser, ok := ctx.Get("CurrentUser"); ok {
		log.Info("CurrentUser:", currentUser)
		return currentUser.(*model.User)
	}
	log.Info("CurrentUser: is null")
	return nil
}

// Success output json data
func (c *BaseController) JsonSuccess(ctx *gin.Context, data interface{}) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"success": true,
		"data":    data,
	})
}

func (c *BaseController) JsonErrorMsg(ctx *gin.Context, message string) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
		"success": false,
		"data":    nil,
	})
}

func (c *BaseController) JsonErrorCode(ctx *gin.Context, code int, message string) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"success": false,
		"data":    nil,
	})
}

func (c *BaseController) JsonPageData(ctx *gin.Context, results interface{}, page *simple.Paging) {

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"success": true,
		"data": &simple.PageResult{
			Results: results,
			Page:    page,
		},
	})

}

func (c *BaseController) JsonCursorData(ctx *gin.Context, results interface{}, cursor string, hasMore bool) {

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"success": true,
		"data": &simple.CursorResult{
			Results: results,
			Cursor:  cursor,
			HasMore: hasMore,
		},
	})

}

// Fail output error
func (c *BaseController) JsonFail(ctx *gin.Context, error *simple.CodeError) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    error.Code,
		"message": error.Message,
	})
	return
}

/*
BuildLoginSuccess 处理登录成功后的返回数据

Parameter:
	user - login user
	ref - 登录来源地址，需要控制登录成功之后跳转到该地址
*/
func (c *BaseController) BuildLoginSuccess(ctx *gin.Context, user *model.User, ref string) {

	token, err := service.UserTokenService.Generate(user.Id)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().
		Put("token", token).
		Put("user", render.BuildUserProfile(user)).
		Put("ref", ref).Build())
	return
}
