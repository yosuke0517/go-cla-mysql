package repository

import (
	"context"
	"go-cla-mysql/entities/model"
)

// todoのCRUDに対するDB用のリポジトリ（ポート）
type TodoRepository interface {
	// TODO 共通化できる？？
	FindAll(ctx context.Context, max int) (todos *model.Todos, err error)
	FindByID(ctx context.Context, id int) (todos *model.Todos, err error)
	Create(todo *model.Todo) (bool, error)
	Update(todo *model.Todo) (bool, error)
}
