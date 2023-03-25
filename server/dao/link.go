package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var LinkDao = newLinkDao()

func newLinkDao() *linkDao {
	return &linkDao{}
}

type linkDao struct {
}

func (d *linkDao) Get(db *gorm.DB, id int64) *model.Link {
	ret := &model.Link{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *linkDao) Take(db *gorm.DB, where ...interface{}) *model.Link {
	ret := &model.Link{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *linkDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Link) {
	cnd.Find(db, &list)
	return
}

func (d *linkDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Link {
	ret := &model.Link{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *linkDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Link, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *linkDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Link, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Link{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *linkDao) Create(db *gorm.DB, t *model.Link) (err error) {
	err = db.Create(t).Error
	return
}

func (d *linkDao) Update(db *gorm.DB, t *model.Link) (err error) {
	err = db.Save(t).Error
	return
}

func (d *linkDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Link{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *linkDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Link{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *linkDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Link{}, "id = ?", id)
}
