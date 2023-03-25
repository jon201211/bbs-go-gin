package service

import (
	"bbs-go/model/constants"
	"strings"

	"bbs-go/util/simple"

	"bbs-go/cache"
	"bbs-go/dao"
	"bbs-go/model"
)

var TagService = newTagService()

func newTagService() *tagService {
	return &tagService{}
}

type tagService struct {
}

func (s *tagService) Get(id int64) *model.Tag {
	return dao.TagDao.Get(simple.DB(), id)
}

func (s *tagService) Take(where ...interface{}) *model.Tag {
	return dao.TagDao.Take(simple.DB(), where...)
}

func (s *tagService) Find(cnd *simple.SqlCnd) []model.Tag {
	return dao.TagDao.Find(simple.DB(), cnd)
}

func (s *tagService) FindOne(cnd *simple.SqlCnd) *model.Tag {
	return dao.TagDao.FindOne(simple.DB(), cnd)
}

func (s *tagService) FindPageByParams(params *simple.QueryParams) (list []model.Tag, paging *simple.Paging) {
	return dao.TagDao.FindPageByParams(simple.DB(), params)
}

func (s *tagService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Tag, paging *simple.Paging) {
	return dao.TagDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *tagService) Create(t *model.Tag) error {
	return dao.TagDao.Create(simple.DB(), t)
}

func (s *tagService) Update(t *model.Tag) error {
	if err := dao.TagDao.Update(simple.DB(), t); err != nil {
		return err
	}
	cache.TagCache.Invalidate(t.Id)
	return nil
}

// func (s *tagService) Updates(id int64, columns map[string]interface{}) error {
// 	return dao.TagDao.Updates(simple.DB(), id, columns)
// }
//
// func (s *tagService) UpdateColumn(id int64, name string, value interface{}) error {
// 	return dao.TagDao.UpdateColumn(simple.DB(), id, name, value)
// }
//
// func (s *tagService) Delete(id int64) {
// 	dao.TagDao.Delete(simple.DB(), id)
// }

// 自动完成
func (s *tagService) Autocomplete(input string) []model.Tag {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil
	}
	return dao.TagDao.Find(simple.DB(), simple.NewSqlCnd().Where("status = ? and name like ?",
		constants.StatusOk, "%"+input+"%").Limit(6))
}

func (s *tagService) GetOrCreate(name string) (*model.Tag, error) {
	return dao.TagDao.GetOrCreate(simple.DB(), name)
}

func (s *tagService) GetByName(name string) *model.Tag {
	return dao.TagDao.GetByName(name)
}

func (s *tagService) GetTags() []model.TagResponse {
	list := dao.TagDao.Find(simple.DB(), simple.NewSqlCnd().Where("status = ?", constants.StatusOk))

	var tags []model.TagResponse
	for _, tag := range list {
		tags = append(tags, model.TagResponse{TagId: tag.Id, TagName: tag.Name})
	}
	return tags
}

func (s *tagService) GetTagInIds(tagIds []int64) []model.Tag {
	return dao.TagDao.GetTagInIds(tagIds)
}

// 扫描
func (s *tagService) Scan(callback func(tags []model.Tag)) {
	var cursor int64
	for {
		list := dao.TagDao.Find(simple.DB(), simple.NewSqlCnd().Where("id > ?", cursor).Asc("id").Limit(100))
		if len(list) == 0 {
			break
		}
		cursor = list[len(list)-1].Id
		callback(list)
	}
}
