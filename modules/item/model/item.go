package model

import (
	"errors"
	"social-todo-list/common"
)

const (
	ItemEntityName = "item"
)

var (
	ErrTitleIsEmpty = errors.New("title cannot be empty")
	ErrItemDeleted = errors.New("item deleted")
)

type TodoItem struct {
	common.SQLModel
	Title string `json:"title"`
	Description string `json:"description"`
	Status *ItemStatus `json:"status"`
}
func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Title string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status *ItemStatus `json:"status" gorm:"column:status"`
}
func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status *ItemStatus `json:"status" gorm:"column:status"`
}
func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }