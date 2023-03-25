package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var ProjectDao = newProjectDao()

func newProjectDao() *projectDao {
	return &projectDao{}
}

type projectDao struct {
}

func (d *projectDao) Get(db *gorm.DB, id int64) *model.Project {
	ret := &model.Project{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *projectDao) Take(db *gorm.DB, where ...interface{}) *model.Project {
	ret := &model.Project{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *projectDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Project) {
	cnd.Find(db, &list)
	return
}

func (d *projectDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Project {
	ret := &model.Project{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *projectDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Project, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *projectDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Project, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Project{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *projectDao) Create(db *gorm.DB, t *model.Project) (err error) {
	err = db.Create(t).Error
	return
}

func (d *projectDao) Update(db *gorm.DB, t *model.Project) (err error) {
	err = db.Save(t).Error
	return
}

func (d *projectDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Project{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *projectDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Project{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *projectDao) Delete(db *gorm.DB, id int64) {
	db.Model(&model.Project{}).Delete("id", id)
}
