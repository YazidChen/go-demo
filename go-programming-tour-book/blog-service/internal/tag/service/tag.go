package service

import (
	"context"
	"yazidchen.com/go-programming-tour-book/blog-service/global"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/tag/dao"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/tag/dto"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/tag/model"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/app"
)

type TagService struct {
	ctx context.Context
	dao *dao.TagDao
}

func NewTagService(ctx context.Context) *TagService {
	return &TagService{ctx: ctx, dao: dao.NewTagDao(global.DBEngine)}
}

func (s *TagService) CountTag(param *dto.TagGetReq) (int, error) {
	return s.dao.CountTag(param.Name, param.State)
}

func (s *TagService) GetTagList(param *dto.TagGetReq, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (s *TagService) CreateTag(param *dto.TagCreateReq) error {
	return s.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (s *TagService) UpdateTag(param *dto.TagUpdateReq) error {
	return s.dao.UpdateTag(param.ID, param.Name, *param.State, param.ModifiedBy)
}

func (s *TagService) DelTag(param *dto.TagDelReq) error {
	return s.dao.DelTag(param.ID)
}
