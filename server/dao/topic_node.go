package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var TopicNodeDao = newTopicNodeDao()

func newTopicNodeDao() *topicNodeDao {
	return &topicNodeDao{}
}

type topicNodeDao struct {
}

func (d *topicNodeDao) Get(db *gorm.DB, id int64) *model.TopicNode {
	ret := &model.TopicNode{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *topicNodeDao) Take(db *gorm.DB, where ...interface{}) *model.TopicNode {
	ret := &model.TopicNode{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *topicNodeDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.TopicNode) {
	cnd.Find(db, &list)
	return
}

func (d *topicNodeDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.TopicNode {
	ret := &model.TopicNode{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *topicNodeDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.TopicNode, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *topicNodeDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.TopicNode, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.TopicNode{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *topicNodeDao) Create(db *gorm.DB, t *model.TopicNode) (err error) {
	err = db.Create(t).Error
	return
}

func (d *topicNodeDao) Update(db *gorm.DB, t *model.TopicNode) (err error) {
	err = db.Save(t).Error
	return
}

func (d *topicNodeDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.TopicNode{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *topicNodeDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.TopicNode{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *topicNodeDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.TopicNode{}, "id = ?", id)
}
