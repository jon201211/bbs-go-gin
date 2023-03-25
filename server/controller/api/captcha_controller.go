package api

import (
	"bbs-go/controller/base"
	"bbs-go/util/urls"

	"bbs-go/util/simple"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CaptchaController struct {
	base.BaseController
}

func (c *CaptchaController) GetRequest(ctx *gin.Context) {
	captchaId := simple.QueryValue(ctx, "captchaId")
	if simple.IsNotBlank(captchaId) { // reload
		if !captcha.Reload(captchaId) {
			// reload 失败，重新加载验证码
			captchaId = captcha.NewLen(4)
		}
	} else {
		captchaId = captcha.NewLen(4)
	}
	captchaUrl := urls.Url("/api/captcha/show?captchaId=" + captchaId + "&r=" + simple.UUID())
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().
		Put("captchaId", captchaId).
		Put("captchaUrl", captchaUrl).
		Build())
}

func (c *CaptchaController) GetShow(ctx *gin.Context) {
	captchaId := simple.QueryValue(ctx, "captchaId")

	if captchaId == "" {
		ctx.AbortWithStatus(404)
		return
	}

	if !captcha.Reload(captchaId) {
		ctx.AbortWithStatus(404)
		return
	}

	ctx.Header("Content-Type", "image/png")
	if err := captcha.WriteImage(ctx.Writer, captchaId, captcha.StdWidth, captcha.StdHeight); err != nil {
		logrus.Error(err)
	}
}

func (c *CaptchaController) GetVerify(ctx *gin.Context) {
	captchaId := ctx.Param("captchaId")
	captchaCode := ctx.Param("captchaCode")
	success := captcha.VerifyString(captchaId, captchaCode)
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("success", success).Build())
}
