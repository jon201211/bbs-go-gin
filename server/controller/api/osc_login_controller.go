package api

import (
	"bbs-go/controller/base"
	"bbs-go/service"
	"bbs-go/util/oauth/osc"
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"
)

type OscLoginController struct {
	base.BaseController
}

// GetAuthorize 获取登录授权地址
func (c *OscLoginController) GetAuthorize(ctx *gin.Context) {
	loginMethod := service.SysConfigService.GetLoginMethod()
	if !loginMethod.Osc {
		c.JsonErrorMsg(ctx, "开源中国账号登录/注册已禁用")
		return
	}

	ref := simple.FormValue(ctx, "ref")
	url := osc.AuthCodeURL(map[string]string{"ref": ref})
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("url", url).Build())
	return
}

// GetCallback 获取回调信息获取
func (c *OscLoginController) GetCallback(ctx *gin.Context) {
	loginMethod := service.SysConfigService.GetLoginMethod()
	if !loginMethod.Osc {
		c.JsonErrorMsg(ctx, "开源中国账号登录/注册已禁用")
		return
	}

	code := simple.FormValue(ctx, "code")
	state := simple.FormValue(ctx, "state")

	thirdAccount, err := service.ThirdAccountService.GetOrCreateByOSC(code, state)
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
