package dao

import (
	"bbs-go/model"
	"bbs-go/util/simple"

	"gorm.io/gorm"
)

var OperateLogDao = newOperateLogDao()

func newOperateLogDao() *operateLogDao {
	return &operateLogDao{}
}

type operateLogDao struct {
}

func (d *operateLogDao) Get(db *gorm.DB, id int64) *model.OperateLog {
	ret := &model.OperateLog{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *operateLogDao) Take(db *gorm.DB, where ...interface{}) *model.OperateLog {
	ret := &model.OperateLog{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *operateLogDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.OperateLog) {
	cnd.Find(db, &list)
	return
}

func (d *operateLogDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.OperateLog {
	ret := &model.OperateLog{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *operateLogDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.OperateLog, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *operateLogDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.OperateLog, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.OperateLog{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *operateLogDao) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.OperateLog{})
}

func (d *operateLogDao) Create(db *gorm.DB, t *model.OperateLog) (err error) {
	err = db.Create(t).Error
	return
}

func (d *operateLogDao) Update(db *gorm.DB, t *model.OperateLog) (err error) {
	err = db.Save(t).Error
	return
}

func (d *operateLogDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.OperateLog{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *operateLogDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.OperateLog{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *operateLogDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.OperateLog{}, "id = ?", id)
}
