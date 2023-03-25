package service

import (
	"bbs-go/util/simple"

	"bbs-go/dao"
	"bbs-go/model"
)

var LinkService = newLinkService()

func newLinkService() *linkService {
	return &linkService{}
}

type linkService struct {
}

func (s *linkService) Get(id int64) *model.Link {
	return dao.LinkDao.Get(simple.DB(), id)
}

func (s *linkService) Take(where ...interface{}) *model.Link {
	return dao.LinkDao.Take(simple.DB(), where...)
}

func (s *linkService) Find(cnd *simple.SqlCnd) []model.Link {
	return dao.LinkDao.Find(simple.DB(), cnd)
}

func (s *linkService) FindOne(cnd *simple.SqlCnd) *model.Link {
	return dao.LinkDao.FindOne(simple.DB(), cnd)
}

func (s *linkService) FindPageByParams(params *simple.QueryParams) (list []model.Link, paging *simple.Paging) {
	return dao.LinkDao.FindPageByParams(simple.DB(), params)
}

func (s *linkService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Link, paging *simple.Paging) {
	return dao.LinkDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *linkService) Create(t *model.Link) error {
	return dao.LinkDao.Create(simple.DB(), t)
}

func (s *linkService) Update(t *model.Link) error {
	return dao.LinkDao.Update(simple.DB(), t)
}

func (s *linkService) Updates(id int64, columns map[string]interface{}) error {
	return dao.LinkDao.Updates(simple.DB(), id, columns)
}

func (s *linkService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.LinkDao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *linkService) Delete(id int64) {
	dao.LinkDao.Delete(simple.DB(), id)
}
