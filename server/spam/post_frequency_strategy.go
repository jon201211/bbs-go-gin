package spam

import (
	"bbs-go/dao"
	"bbs-go/model"
	"bbs-go/util/simple"
	"bbs-go/util/simple/date"
	"errors"
	"time"
)

// PostFrequencyStrategy 发表频率限制
type PostFrequencyStrategy struct{}

func (PostFrequencyStrategy) Name() string {
	return "PostFrequencyStrategy"
}

func (PostFrequencyStrategy) CheckTopic(user *model.User, topic model.CreateTopicForm) error {
	// 注册时间超过24小时
	if user.CreateTime < date.Timestamp(time.Now().Add(-time.Hour*24)) {
		return nil
	}
	var (
		maxCountInTenMinutes int64 = 1 // 十分钟内最高发帖数量
		maxCountInOneHour    int64 = 2 // 一小时内最高发帖量
		maxCountInOneDay     int64 = 3 // 一天内最高发帖量
	)

	if dao.TopicDao.Count(simple.DB(), simple.NewSqlCnd().Eq("user_id", user.Id).
		Gt("create_time", date.Timestamp(time.Now().Add(-time.Hour*24)))) >= maxCountInOneDay {
		return errors.New("发表太快了，请休息一会儿")
	}

	if dao.TopicDao.Count(simple.DB(), simple.NewSqlCnd().Eq("user_id", user.Id).
		Gt("create_time", date.Timestamp(time.Now().Add(-time.Hour)))) >= maxCountInOneHour {
		return errors.New("发表太快了，请休息一会儿")
	}

	if dao.TopicDao.Count(simple.DB(), simple.NewSqlCnd().Eq("user_id", user.Id).
		Gt("create_time", date.Timestamp(time.Now().Add(-time.Minute*10)))) >= maxCountInTenMinutes {
		return errors.New("发表太快了，请休息一会儿")
	}

	return nil
}

func (s PostFrequencyStrategy) CheckArticle(user *model.User, form model.CreateArticleForm) error {
	// 注册时间超过24小时
	if user.CreateTime < date.Timestamp(time.Now().Add(-time.Hour*24)) {
		return nil
	}
	var (
		maxCountInTenMinutes int64 = 1 // 十分钟内最高发帖数量
		maxCountInOneHour    int64 = 2 // 一小时内最高发帖量
		maxCountInOneDay     int64 = 3 // 一天内最高发帖量
	)

	if dao.ArticleDao.Count(simple.DB(), simple.NewSqlCnd().Eq("user_id", user.Id).
		Gt("create_time", date.Timestamp(time.Now().Add(-time.Hour*24)))) >= maxCountInOneDay {
		return errors.New("发表太快了，请休息一会儿")
	}

	if dao.ArticleDao.Count(simple.DB(), simple.NewSqlCnd().Eq("user_id", user.Id).
		Gt("create_time", date.Timestamp(time.Now().Add(-time.Hour)))) >= maxCountInOneHour {
		return errors.New("发表太快了，请休息一会儿")
	}

	if dao.ArticleDao.Count(simple.DB(), simple.NewSqlCnd().Eq("user_id", user.Id).
		Gt("create_time", date.Timestamp(time.Now().Add(-time.Minute*10)))) >= maxCountInTenMinutes {
		return errors.New("发表太快了，请休息一会儿")
	}

	return nil
}

func (s PostFrequencyStrategy) CheckComment(user *model.User, form model.CreateCommentForm) error {
	// 注册时间超过24小时
	if user.CreateTime < date.Timestamp(time.Now().Add(-time.Hour*24)) {
		return nil
	}

	var (
		maxCountInTenMinutes int64 = 1 // 十分钟内最高发帖数量
		maxCountInOneHour    int64 = 1 // 一小时内最高发帖量
		maxCountInOneDay     int64 = 1 // 一天内最高发帖量
	)

	if dao.CommentDao.Count(simple.DB(), simple.NewSqlCnd().Eq("user_id", user.Id).
		Gt("create_time", date.Timestamp(time.Now().Add(-time.Hour*24)))) >= maxCountInOneDay {
		return errors.New("发表太快了，请休息一会儿")
	}

	if dao.CommentDao.Count(simple.DB(), simple.NewSqlCnd().Eq("user_id", user.Id).
		Gt("create_time", date.Timestamp(time.Now().Add(-time.Hour)))) >= maxCountInOneHour {
		return errors.New("发表太快了，请休息一会儿")
	}

	if dao.CommentDao.Count(simple.DB(), simple.NewSqlCnd().Eq("user_id", user.Id).
		Gt("create_time", date.Timestamp(time.Now().Add(-time.Minute*10)))) >= maxCountInTenMinutes {
		return errors.New("发表太快了，请休息一会儿")
	}
	return nil
}
