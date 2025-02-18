package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *sqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
        return common.ErrorDB(err)
	}

	return nil
}