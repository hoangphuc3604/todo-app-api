package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"

	"gorm.io/gorm"
)

func (s *sqlStorage) GetItem(
	ctx context.Context, 
	cond map[string]interface{},
) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &data, nil
}