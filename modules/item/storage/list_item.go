package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *sqlStorage) ListItems(
	ctx context.Context, 
	filter *model.Filter, 
	paging *common.Paging, 
	moreKeys ...string,
) ([]model.TodoItem, error) {
	var data []model.TodoItem
	db := s.db.Where("status <> ?", "DELETED")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").Limit(paging.Limit).Offset(paging.Offset()).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}