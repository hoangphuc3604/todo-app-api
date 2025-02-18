package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (s *sqlStorage) UpdateItem(
	ctx context.Context, 
	cond map[string]interface{},
	data *model.TodoItemUpdate,
) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}

	return nil
}