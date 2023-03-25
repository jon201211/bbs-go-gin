package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var TopicDao = newTopicDao()

func newTopicDao() *topicDao {
	return &topicDao{}
}

type topicDao struct {
}

func (d *topicDao) Get(db *gorm.DB, id int64) *model.Topic {
	ret := &model.Topic{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *topicDao) Take(db *gorm.DB, where ...interface{}) *model.Topic {
	ret := &model.Topic{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *topicDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Topic) {
	cnd.Find(db, &list)
	return
}

func (d *topicDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Topic {
	ret := &model.Topic{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *topicDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Topic, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *topicDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Topic, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Topic{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *topicDao) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.Topic{})
}

func (d *topicDao) Create(db *gorm.DB, t *model.Topic) (err error) {
	err = db.Create(t).Error
	return
}

func (d *topicDao) Update(db *gorm.DB, t *model.Topic) (err error) {
	err = db.Save(t).Error
	return
}

func (d *topicDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Topic{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *topicDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Topic{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *topicDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Topic{}, "id = ?", id)
}
