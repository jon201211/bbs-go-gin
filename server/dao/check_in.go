package dao

import (
	"bbs-go/model"
	"bbs-go/util/simple"

	"gorm.io/gorm"
)

var CheckInDao = newCheckInDao()

func newCheckInDao() *checkInDao {
	return &checkInDao{}
}

type checkInDao struct {
}

func (d *checkInDao) Get(db *gorm.DB, id int64) *model.CheckIn {
	ret := &model.CheckIn{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *checkInDao) Take(db *gorm.DB, where ...interface{}) *model.CheckIn {
	ret := &model.CheckIn{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *checkInDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.CheckIn) {
	cnd.Find(db, &list)
	return
}

func (d *checkInDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.CheckIn {
	ret := &model.CheckIn{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *checkInDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.CheckIn, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *checkInDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.CheckIn, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.CheckIn{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *checkInDao) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.CheckIn{})
}

func (d *checkInDao) Create(db *gorm.DB, t *model.CheckIn) (err error) {
	err = db.Create(t).Error
	return
}

func (d *checkInDao) Update(db *gorm.DB, t *model.CheckIn) (err error) {
	err = db.Save(t).Error
	return
}

func (d *checkInDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.CheckIn{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *checkInDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.CheckIn{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *checkInDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.CheckIn{}, "id = ?", id)
}
