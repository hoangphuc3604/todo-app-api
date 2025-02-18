package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(
		ctx context.Context, 
		cond map[string]interface{},
	) error
}

type deleteItemBiz struct {
	storage DeleteItemStorage
}

func NewDeleteItemBiz(storage DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{storage: storage}
}

func (biz *deleteItemBiz) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"ID": id})
	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrorCanNotGetEntity(model.ItemEntityName, err)
		}

		return common.ErrorCanNotDeleteEntity(model.ItemEntityName, err)
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return common.ErrorEntityDeleted(model.ItemEntityName, model.ErrItemDeleted)
	}

	if err := biz.storage.DeleteItem(ctx, map[string]interface{}{"ID": id}); err != nil {
		return common.ErrorCanNotDeleteEntity(model.ItemEntityName, err)
	}

	return nil
}