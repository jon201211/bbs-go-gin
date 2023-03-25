package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var CommentDao = newCommentDao()

func newCommentDao() *commentDao {
	return &commentDao{}
}

type commentDao struct {
}

func (d *commentDao) Get(db *gorm.DB, id int64) *model.Comment {
	ret := &model.Comment{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *commentDao) Take(db *gorm.DB, where ...interface{}) *model.Comment {
	ret := &model.Comment{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *commentDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Comment) {
	cnd.Find(db, &list)
	return
}

func (d *commentDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Comment {
	ret := &model.Comment{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *commentDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Comment, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *commentDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Comment, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Comment{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *commentDao) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.Comment{})
}

func (d *commentDao) Create(db *gorm.DB, t *model.Comment) (err error) {
	err = db.Create(t).Error
	return
}

func (d *commentDao) Update(db *gorm.DB, t *model.Comment) (err error) {
	err = db.Save(t).Error
	return
}

func (d *commentDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Comment{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *commentDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Comment{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *commentDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Comment{}, "id = ?", id)
}
