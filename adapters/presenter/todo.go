package presenter

import (
	"go-cla-practice/entities/model"
	"go-cla-practice/usecases/dto"
	"go-cla-practice/usecases/port"
	"net/http"
)

type Todo struct {
}

func NewTodoOutputPort() port.TodoOutputPort {
	return &Todo{}
}

// usecasesのTodoOutputPortを実装
func (t *Todo) Render(writer http.ResponseWriter, todos *model.Todos) {
	var hits = len(*todos)
	var todoOutPutUseCaseDto = dto.NewTodoOutPutUseCaseDto(hits, *todos)
	Success(writer, todoOutPutUseCaseDto)
}

// TODO jsonで返す
func (t *Todo) RenderError(writer http.ResponseWriter, err error) {
	InternalServerError(writer, "Server Error Please check log")
}
