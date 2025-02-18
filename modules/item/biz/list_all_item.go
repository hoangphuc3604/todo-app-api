package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type ListItemStorage interface {
	ListItems(
		ctx context.Context, 
		filter *model.Filter, 
		paging *common.Paging, 
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBiz struct {
	storage ListItemStorage
}

func NewListItemBiz(storage ListItemStorage) *listItemBiz {
	return &listItemBiz{storage: storage}
}

func (biz *listItemBiz) ListAllItem(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	data, err := biz.storage.ListItems(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrorCanNotListEntity(model.ItemEntityName, err)
	}

	return data, nil
}