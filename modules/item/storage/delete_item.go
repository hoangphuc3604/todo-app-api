package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (s *sqlStorage) DeleteItem(
	ctx context.Context, 
	cond map[string]interface{},
) error {
	deletedStatus := model.ItemStatusDeleted

	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": deletedStatus.String(),
	}).Error; err != nil {
		return err
	}

	return nil
}