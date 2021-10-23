package repository

import "go-cla-mysql/entities/model"

// todoのCRUDに対するDB用のリポジトリ（ポート）
type TodoRepository interface {
	FindAll(max int) (todos []*model.Todo, err error)
	FindByID(id int) (todos []*model.Todo, err error)
	Create(todo *model.Todo) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
}
