package port

import (
	"context"
	"go-cla-mysql/entities/model"
	"go-cla-mysql/usecases/dto"
)

type TodoInputPort interface {
	Create(todo *model.Todo) (bool, error)
	FindAll(ctx context.Context, max int) (*dto.TodoOutPutUseCaseDto, error)
	FindByID(ctx context.Context, id int) (*dto.TodoOutPutUseCaseDto, error)
	Update(todo *model.Todo) (bool, error)
}

// TODO 共通のレスポンスの処理入れているのでいらないかもしれない…
type TodoOutputPort interface {
	Render(todos *model.Todos) dto.TodoOutPutUseCaseDto
	RenderError(error)
}
