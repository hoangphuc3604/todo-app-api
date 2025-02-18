package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
	"strings"
)

type CreatItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createItemBiz struct {
	storage CreatItemStorage
}
func NewCreateItemBiz(storage CreatItemStorage) *createItemBiz {
	return &createItemBiz{storage: storage}
}
func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsEmpty
	}

	if err := biz.storage.CreateItem(ctx, data); err != nil {
		return common.ErrorCanNotCreateEntity(model.ItemEntityName, err)
	}

	return nil
}