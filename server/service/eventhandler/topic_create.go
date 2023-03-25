package eventhandler

import (
	"bbs-go/model"
	"bbs-go/model/constants"
	"bbs-go/service"
	"bbs-go/util/event"
	"bbs-go/util/seo"
	"bbs-go/util/urls"
	"reflect"

	"github.com/sirupsen/logrus"
)

func init() {
	event.RegHandler(reflect.TypeOf(event.TopicCreateEvent{}), handleTopicCreateEvent)
}

func handleTopicCreateEvent(i interface{}) {
	e := i.(event.TopicCreateEvent)

	// 百度链接推送
	seo.Push(urls.TopicUrl(e.TopicId))

	service.UserFollowService.ScanFans(e.UserId, func(fansId int64) {
		logrus.WithField("topicId", e.TopicId).
			WithField("userId", e.UserId).
			WithField("fansId", fansId).
			Info("用户关注，处理帖子")
		if err := service.UserFeedService.Create(&model.UserFeed{
			UserId:     fansId,
			DataId:     e.TopicId,
			DataType:   constants.EntityTopic,
			AuthorId:   e.UserId,
			CreateTime: e.CreateTime,
		}); err != nil {
			logrus.Error(err)
		}
	})
}
