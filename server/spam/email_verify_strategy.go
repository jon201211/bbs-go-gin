package spam

import (
	"bbs-go/model"
	"bbs-go/service"
	"bbs-go/util/common"
)

type EmailVerifyStrategy struct{}

func (EmailVerifyStrategy) Name() string {
	return "EmailVerifyStrategy"
}

func (EmailVerifyStrategy) CheckTopic(user *model.User, form model.CreateTopicForm) error {
	if service.SysConfigService.IsCreateTopicEmailVerified() && !user.EmailVerified {
		return common.EmailNotVerified
	}
	return nil
}

func (EmailVerifyStrategy) CheckArticle(user *model.User, form model.CreateArticleForm) error {
	if service.SysConfigService.IsCreateArticleEmailVerified() && !user.EmailVerified {
		return common.EmailNotVerified
	}
	return nil
}

func (EmailVerifyStrategy) CheckComment(user *model.User, form model.CreateCommentForm) error {
	if service.SysConfigService.IsCreateCommentEmailVerified() && !user.EmailVerified {
		return common.EmailNotVerified
	}
	return nil
}
