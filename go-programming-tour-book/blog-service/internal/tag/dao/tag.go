package dao

import (
	"github.com/jinzhu/gorm"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/base"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/tag/model"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/app"
)

type TagDao struct {
	engine *gorm.DB
}

func NewTagDao(db *gorm.DB) *TagDao {
	return &TagDao{engine: db}
}

func (d *TagDao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *TagDao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	offset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, offset, pageSize)
}

func (d *TagDao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{Name: name, State: state, Model: &base.Model{CreatedBy: createBy}}
	return tag.Create(d.engine)
}

func (d *TagDao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{Name: name, State: state, Model: &base.Model{ID: id, ModifiedBy: modifiedBy}}
	return tag.Update(d.engine)
}

func (d *TagDao) DelTag(id uint32) error {
	tag := model.Tag{Model: &base.Model{ID: id}}
	return tag.Del(d.engine)
}
