package service

import (
	"bbs-go/model/constants"
	"errors"
	"strings"

	"bbs-go/util/simple/date"
	"bbs-go/util/simple/json"

	"github.com/sirupsen/logrus"

	"bbs-go/util/simple"

	"bbs-go/dao"
	"bbs-go/model"
)

var CommentService = newCommentService()

func newCommentService() *commentService {
	return &commentService{}
}

type commentService struct {
}

func (s *commentService) Get(id int64) *model.Comment {
	return dao.CommentDao.Get(simple.DB(), id)
}

func (s *commentService) Take(where ...interface{}) *model.Comment {
	return dao.CommentDao.Take(simple.DB(), where...)
}

func (s *commentService) Find(cnd *simple.SqlCnd) []model.Comment {
	return dao.CommentDao.Find(simple.DB(), cnd)
}

func (s *commentService) FindOne(cnd *simple.SqlCnd) *model.Comment {
	return dao.CommentDao.FindOne(simple.DB(), cnd)
}

func (s *commentService) FindPageByParams(params *simple.QueryParams) (list []model.Comment, paging *simple.Paging) {
	return dao.CommentDao.FindPageByParams(simple.DB(), params)
}

func (s *commentService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Comment, paging *simple.Paging) {
	return dao.CommentDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *commentService) Count(cnd *simple.SqlCnd) int64 {
	return dao.CommentDao.Count(simple.DB(), cnd)
}

func (s *commentService) Create(t *model.Comment) error {
	return dao.CommentDao.Create(simple.DB(), t)
}

func (s *commentService) Update(t *model.Comment) error {
	return dao.CommentDao.Update(simple.DB(), t)
}

func (s *commentService) Updates(id int64, columns map[string]interface{}) error {
	return dao.CommentDao.Updates(simple.DB(), id, columns)
}

func (s *commentService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.CommentDao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *commentService) Delete(id int64) error {
	return dao.CommentDao.UpdateColumn(simple.DB(), id, "status", constants.StatusDeleted)
}

// Publish 发表评论
func (s *commentService) Publish(userId int64, form model.CreateCommentForm) (*model.Comment, error) {
	form.Content = strings.TrimSpace(form.Content)
	if simple.IsBlank(form.EntityType) {
		return nil, errors.New("参数非法")
	}
	if form.EntityId <= 0 {
		return nil, errors.New("参数非法")
	}
	if simple.IsBlank(form.Content) {
		return nil, errors.New("请输入评论内容")
	}

	comment := &model.Comment{
		UserId:      userId,
		EntityType:  form.EntityType,
		EntityId:    form.EntityId,
		Content:     form.Content,
		ContentType: simple.DefaultIfBlank(form.ContentType, constants.ContentTypeMarkdown),
		QuoteId:     form.QuoteId,
		Status:      constants.StatusOk,
		CreateTime:  date.NowTimestamp(),
	}

	if len(form.ImageList) > 0 {
		imageListStr, err := json.ToStr(form.ImageList)
		if err == nil {
			comment.ImageList = imageListStr
		} else {
			logrus.Error(err)
		}
	}

	if err := s.Create(comment); err != nil {
		return nil, err
	}

	if form.EntityType == constants.EntityTopic {
		TopicService.OnComment(form.EntityId, comment)
	}

	UserService.IncrCommentCount(userId)         // 用户跟帖计数
	UserService.IncrScoreForPostComment(comment) // 获得积分
	MessageService.SendCommentMsg(comment)       // 发送消息

	return comment, nil
}

// // 统计数量
// func (s *commentService) Count(entityType string, entityId int64) int64 {
// 	var count int64 = 0
// 	simple.DB().Model(&model.Comment{}).Where("entity_type = ? and entity_id = ?", entityType, entityId).Count(&count)
// 	return count
// }

// GetComments 列表
func (s *commentService) GetComments(entityType string, entityId int64, cursor int64) (comments []model.Comment, nextCursor int64, hasMore bool) {
	limit := 50
	cnd := simple.NewSqlCnd().Eq("entity_type", entityType).Eq("entity_id", entityId).Eq("status", constants.StatusOk).Desc("id").Limit(limit)
	if cursor > 0 {
		cnd.Lt("id", cursor)
	}
	comments = dao.CommentDao.Find(simple.DB(), cnd)
	if len(comments) > 0 {
		nextCursor = comments[len(comments)-1].Id
		hasMore = len(comments) >= limit
	} else {
		nextCursor = cursor
	}
	return
}

// ScanByUser 按照用户扫描数据
func (s *commentService) ScanByUser(userId int64, callback func(comments []model.Comment)) {
	var cursor int64 = 0
	for {
		list := dao.CommentDao.Find(simple.DB(), simple.NewSqlCnd().
			Eq("user_id", userId).Gt("id", cursor).Asc("id").Limit(1000))
		if len(list) == 0 {
			break
		}
		cursor = list[len(list)-1].Id
		callback(list)
	}
}
