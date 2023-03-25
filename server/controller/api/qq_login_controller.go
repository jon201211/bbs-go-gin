package api

import (
	"bbs-go/controller/base"
	"bbs-go/service"
	"bbs-go/util/oauth/qq"
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"
)

type QQLoginController struct {
	base.BaseController
}

// 获取QQ登录授权地址
func (c *QQLoginController) GetAuthorize(ctx *gin.Context) {
	loginMethod := service.SysConfigService.GetLoginMethod()
	if !loginMethod.QQ {
		c.JsonErrorMsg(ctx, "QQ登录/注册已禁用")
		return
	}

	ref := simple.FormValue(ctx, "ref")
	url := qq.AuthorizeUrl(map[string]string{"ref": ref})
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("url", url).Build())
	return
}

// 获取QQ回调信息获取
func (c *QQLoginController) GetCallback(ctx *gin.Context) {
	loginMethod := service.SysConfigService.GetLoginMethod()
	if !loginMethod.QQ {
		c.JsonErrorMsg(ctx, "QQ登录/注册已禁用")
		return
	}

	code := simple.FormValue(ctx, "code")
	state := simple.FormValue(ctx, "state")

	thirdAccount, err := service.ThirdAccountService.GetOrCreateByQQ(code, state)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	user, codeErr := service.UserService.SignInByThirdAccount(thirdAccount)
	if codeErr != nil {
		c.JsonErrorMsg(ctx, codeErr.Error())
		return
	} else {
		c.BuildLoginSuccess(ctx, user, "")
		return
	}
}
