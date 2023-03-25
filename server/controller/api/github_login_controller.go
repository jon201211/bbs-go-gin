package api

import (
	"bbs-go/controller/base"
	"bbs-go/service"
	"bbs-go/util/oauth/github"
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"
)

type GithubLoginController struct {
	base.BaseController
}

// 获取Github登录授权地址
func (c *GithubLoginController) GetAuthorize(ctx *gin.Context) {
	loginMethod := service.SysConfigService.GetLoginMethod()
	if !loginMethod.Github {
		c.JsonErrorMsg(ctx, "Github登录/注册已禁用")
		return
	}

	ref := simple.FormValue(ctx, "ref")
	url := github.AuthCodeURL(map[string]string{"ref": ref})
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("url", url).Build())
	return
}

// 获取Github回调信息获取
func (c *GithubLoginController) GetCallback(ctx *gin.Context) {
	loginMethod := service.SysConfigService.GetLoginMethod()
	if !loginMethod.Github {
		c.JsonErrorMsg(ctx, "Github登录/注册已禁用")
		return
	}

	code := simple.FormValue(ctx, "code")
	state := simple.FormValue(ctx, "state")

	thirdAccount, err := service.ThirdAccountService.GetOrCreateByGithub(code, state)
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
