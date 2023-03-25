package eventhandler

import (
	"bbs-go/service"
	"bbs-go/util/event"
	"reflect"
)

func init() {
	event.RegHandler(reflect.TypeOf(event.UnFollowEvent{}), handleUnFollowEvent)
}

func handleUnFollowEvent(i interface{}) {
	e := i.(event.UnFollowEvent)

	// 清理该用户下的信息流
	service.UserFeedService.DeleteByUser(e.UserId, e.OtherId)
}
