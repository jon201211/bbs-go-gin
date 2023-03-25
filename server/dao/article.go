package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var ArticleDao = newArticleDao()

func newArticleDao() *articleDao {
	return &articleDao{}
}

type articleDao struct {
}

func (d *articleDao) Get(db *gorm.DB, id int64) *model.Article {
	ret := &model.Article{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *articleDao) Take(db *gorm.DB, where ...interface{}) *model.Article {
	ret := &model.Article{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *articleDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Article) {
	cnd.Find(db, &list)
	return
}

func (d *articleDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Article {
	ret := &model.Article{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *articleDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Article, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *articleDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Article, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Article{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *articleDao) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.Article{})
}

func (d *articleDao) Create(db *gorm.DB, t *model.Article) (err error) {
	err = db.Create(t).Error
	return
}

func (d *articleDao) Update(db *gorm.DB, t *model.Article) (err error) {
	err = db.Save(t).Error
	return
}

func (d *articleDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Article{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *articleDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Article{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *articleDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Article{}, "id = ?", id)
}
