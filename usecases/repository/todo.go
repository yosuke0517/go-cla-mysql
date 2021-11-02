package repository

import (
	"go-cla-practice/entities/model"
)

// todoのCRUDに対するDB用のリポジトリ（ポート）
type TodoRepository interface {
	// TODO 共通化できる？？
	FindAll(max int) (todos *model.Todos, err error)
	FindByID(id int) (todos *model.Todos, err error)
	Create(todo *model.Todo) (bool, error)
	Update(todo *model.Todo) (bool, error)
}
