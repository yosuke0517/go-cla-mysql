package presenter

import (
	"fmt"
	"go-cla-practice/entities/model"
	"go-cla-practice/usecases/dto"
	"go-cla-practice/usecases/port"
	"net/http"
)

type Todo struct {
	w http.ResponseWriter
}

func NewTodoOutputPort(w http.ResponseWriter) port.TodoOutputPort {
	return &Todo{
		w: w,
	}
}

// usecasesのTodoOutputPortを実装
func (t *Todo) Render(todos *model.Todos) dto.TodoOutPutUseCaseDto {
	var output dto.TodoOutPutUseCaseDto
	output.Hits = len(*todos)
	output.Todos = *todos
	return output
}

// TODO jsonで返す
func (t *Todo) RenderError(err error) {
	t.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(t.w, err)
}
