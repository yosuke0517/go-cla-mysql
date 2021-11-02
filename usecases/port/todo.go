package port

import (
	"go-cla-practice/entities/model"
	"go-cla-practice/usecases/dto"
)

type TodoInputPort interface {
	Create(todo *model.Todo) (bool, error)
	FindAll(max int) (*dto.TodoOutPutUseCaseDto, error)
	FindByID(id int) (*dto.TodoOutPutUseCaseDto, error)
	Update(todo *model.Todo) (bool, error)
}

// TODO 共通のレスポンスの処理入れているのでいらないかもしれない…
type TodoOutputPort interface {
	Render(todos *model.Todos) dto.TodoOutPutUseCaseDto
	RenderError(error)
}
