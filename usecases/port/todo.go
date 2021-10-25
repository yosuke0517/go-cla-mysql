package port

import (
	"go-cla-mysql/entities/model"
	"go-cla-mysql/usecases/dto"
)

type TodoInputPort interface {
	Create() model.Todos
	FindAll(max int) model.Todos
	FindByID(id int) (model.Todos, error)
	Update(todo model.Todo) (model.Todos, error)
}

type TodoOutputPort interface {
	Render(todos *model.Todos) dto.TodoOutPutUseCaseDto
	RenderError(error)
}
