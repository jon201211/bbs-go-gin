package service

import (
	"bbs-go/model/constants"
	"bbs-go/util/simple"

	"bbs-go/dao"
	"bbs-go/model"
)

var ArticleTagService = newArticleTagService()

func newArticleTagService() *articleTagService {
	return &articleTagService{}
}

type articleTagService struct {
}

func (s *articleTagService) Get(id int64) *model.ArticleTag {
	return dao.ArticleTagDao.Get(simple.DB(), id)
}

func (s *articleTagService) Take(where ...interface{}) *model.ArticleTag {
	return dao.ArticleTagDao.Take(simple.DB(), where...)
}

func (s *articleTagService) Find(cnd *simple.SqlCnd) []model.ArticleTag {
	return dao.ArticleTagDao.Find(simple.DB(), cnd)
}

func (s *articleTagService) FindPageByParams(params *simple.QueryParams) (list []model.ArticleTag, paging *simple.Paging) {
	return dao.ArticleTagDao.FindPageByParams(simple.DB(), params)
}

func (s *articleTagService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.ArticleTag, paging *simple.Paging) {
	return dao.ArticleTagDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *articleTagService) Create(t *model.ArticleTag) error {
	return dao.ArticleTagDao.Create(simple.DB(), t)
}

func (s *articleTagService) Update(t *model.ArticleTag) error {
	return dao.ArticleTagDao.Update(simple.DB(), t)
}

func (s *articleTagService) Updates(id int64, columns map[string]interface{}) error {
	return dao.ArticleTagDao.Updates(simple.DB(), id, columns)
}

func (s *articleTagService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.ArticleTagDao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *articleTagService) DeleteByArticleId(topicId int64) {
	simple.DB().Model(model.ArticleTag{}).Where("topic_id = ?", topicId).UpdateColumn("status", constants.StatusDeleted)
}
