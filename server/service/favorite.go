package service

import (
	"bbs-go/model/constants"
	"bbs-go/util/simple/date"
	"errors"

	"bbs-go/util/simple"

	"bbs-go/dao"
	"bbs-go/model"
)

var FavoriteService = newFavoriteService()

func newFavoriteService() *favoriteService {
	return &favoriteService{}
}

type favoriteService struct {
}

func (s *favoriteService) Get(id int64) *model.Favorite {
	return dao.FavoriteDao.Get(simple.DB(), id)
}

func (s *favoriteService) Take(where ...interface{}) *model.Favorite {
	return dao.FavoriteDao.Take(simple.DB(), where...)
}

func (s *favoriteService) Find(cnd *simple.SqlCnd) []model.Favorite {
	return dao.FavoriteDao.Find(simple.DB(), cnd)
}

func (s *favoriteService) FindOne(cnd *simple.SqlCnd) *model.Favorite {
	return dao.FavoriteDao.FindOne(simple.DB(), cnd)
}

func (s *favoriteService) FindPageByParams(params *simple.QueryParams) (list []model.Favorite, paging *simple.Paging) {
	return dao.FavoriteDao.FindPageByParams(simple.DB(), params)
}

func (s *favoriteService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Favorite, paging *simple.Paging) {
	return dao.FavoriteDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *favoriteService) Create(t *model.Favorite) error {
	return dao.FavoriteDao.Create(simple.DB(), t)
}

func (s *favoriteService) Update(t *model.Favorite) error {
	return dao.FavoriteDao.Update(simple.DB(), t)
}

func (s *favoriteService) Updates(id int64, columns map[string]interface{}) error {
	return dao.FavoriteDao.Updates(simple.DB(), id, columns)
}

func (s *favoriteService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.FavoriteDao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *favoriteService) Delete(id int64) {
	dao.FavoriteDao.Delete(simple.DB(), id)
}

func (s *favoriteService) GetBy(userId int64, entityType string, entityId int64) *model.Favorite {
	return dao.FavoriteDao.Take(simple.DB(), "user_id = ? and entity_type = ? and entity_id = ?",
		userId, entityType, entityId)
}

// 收藏文章
func (s *favoriteService) AddArticleFavorite(userId, articleId int64) error {
	article := dao.ArticleDao.Get(simple.DB(), articleId)
	if article == nil || article.Status != constants.StatusOk {
		return errors.New("收藏的文章不存在")
	}
	return s.addFavorite(userId, constants.EntityArticle, articleId)
}

// 收藏主题
func (s *favoriteService) AddTopicFavorite(userId, topicId int64) error {
	topic := dao.TopicDao.Get(simple.DB(), topicId)
	if topic == nil || topic.Status != constants.StatusOk {
		return errors.New("收藏的话题不存在")
	}
	if err := s.addFavorite(userId, constants.EntityTopic, topicId); err != nil {
		return err
	}

	// 发送消息
	MessageService.SendTopicFavoriteMsg(topicId, userId)
	return nil
}

func (s *favoriteService) addFavorite(userId int64, entityType string, entityId int64) error {
	temp := s.GetBy(userId, entityType, entityId)
	if temp != nil { // 已经收藏
		return nil
	}
	return dao.FavoriteDao.Create(simple.DB(), &model.Favorite{
		UserId:     userId,
		EntityType: entityType,
		EntityId:   entityId,
		CreateTime: date.NowTimestamp(),
	})
}
