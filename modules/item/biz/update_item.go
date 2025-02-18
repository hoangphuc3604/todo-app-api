package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(
		ctx context.Context, 
		cond map[string]interface{},
		data *model.TodoItemUpdate,
	) error
}

type updateItemBiz struct {
	storage UpdateItemStorage
}

func NewUpdateItemBiz(storage UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{storage: storage}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"ID": id})
	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrorCanNotGetEntity(model.ItemEntityName, err)
		}

		return common.ErrorCanNotUpdateEntity(model.ItemEntityName, err)
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return common.ErrorEntityDeleted(model.ItemEntityName, model.ErrItemDeleted)
	}
	
	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"ID": id}, dataUpdate); err != nil {
		return common.ErrorCanNotUpdateEntity(model.ItemEntityName, err)
	}

	return nil
}