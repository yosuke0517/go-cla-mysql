package port

import (
	"context"
	"go-cla-mysql/entities/model"
	"go-cla-mysql/usecases/dto"
)

type TodoInputPort interface {
	Create(ctx context.Context) model.Todos
	FindAll(ctx context.Context, max int) dto.TodoOutPutUseCaseDto
	FindByID(ctx context.Context, id int) (model.Todos, error)
	Update(ctx context.Context, todo model.Todo) (model.Todos, error)
}

type TodoOutputPort interface {
	Render(todos *model.Todos) dto.TodoOutPutUseCaseDto
	RenderError(error)
}
