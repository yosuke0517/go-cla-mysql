package repository

import (
	"context"
	"go-cla-mysql/entities/model"
)

// todoのCRUDに対するDB用のリポジトリ（ポート）
type TodoRepository interface {
	// TODO 共通化できる？？
	FindAll(max int) (todos *model.Todos, err error)
	FindByID(ctx context.Context, id int) (todos *model.Todo, err error)
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	Update(ctx context.Context, todo *model.Todo) (*model.Todo, error)
}
