package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var SysConfigDao = newSysConfigDao()

func newSysConfigDao() *sysConfigDao {
	return &sysConfigDao{}
}

type sysConfigDao struct {
}

func (d *sysConfigDao) Get(db *gorm.DB, id int64) *model.SysConfig {
	ret := &model.SysConfig{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *sysConfigDao) Take(db *gorm.DB, where ...interface{}) *model.SysConfig {
	ret := &model.SysConfig{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *sysConfigDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.SysConfig) {
	cnd.Find(db, &list)
	return
}

func (d *sysConfigDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.SysConfig {
	ret := &model.SysConfig{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *sysConfigDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.SysConfig, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *sysConfigDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.SysConfig, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.SysConfig{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *sysConfigDao) Create(db *gorm.DB, t *model.SysConfig) (err error) {
	err = db.Create(t).Error
	return
}

func (d *sysConfigDao) Update(db *gorm.DB, t *model.SysConfig) (err error) {
	err = db.Save(t).Error
	return
}

func (d *sysConfigDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.SysConfig{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *sysConfigDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.SysConfig{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *sysConfigDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.SysConfig{}, "id = ?", id)
}

func (d *sysConfigDao) GetByKey(db *gorm.DB, key string) *model.SysConfig {
	if len(key) == 0 {
		return nil
	}
	return d.Take(db, "`key` = ?", key)
}
