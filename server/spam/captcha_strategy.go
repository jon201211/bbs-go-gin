package spam

import (
	"bbs-go/model"
	"bbs-go/service"
	"bbs-go/util/common"

	"github.com/dchest/captcha"
)

type CaptchaStrategy struct{}

func (CaptchaStrategy) Name() string {
	return "CaptchaStrategy"
}

func (CaptchaStrategy) CheckTopic(user *model.User, form model.CreateTopicForm) error {
	if service.SysConfigService.GetConfig().TopicCaptcha && !captcha.VerifyString(form.CaptchaId, form.CaptchaCode) {
		return common.CaptchaError
	}
	return nil
}

func (CaptchaStrategy) CheckArticle(user *model.User, form model.CreateArticleForm) error {
	return nil
}

func (CaptchaStrategy) CheckComment(user *model.User, form model.CreateCommentForm) error {
	return nil
}
