package port

import (
	"go-cla-practice/entities/model"
	"net/http"
)

type TodoInputPort interface {
	Create(writer http.ResponseWriter, todo *model.Todo)
	FindAll(writer http.ResponseWriter, max int)
	FindByID(writer http.ResponseWriter, id int)
	Update(writer http.ResponseWriter, todo *model.Todo)
}

// TodoOutputPort この中で共通のレスポンス処理を呼ぶ
type TodoOutputPort interface {
	Render(writer http.ResponseWriter, todos *model.Todos)
	RenderError(writer http.ResponseWriter, err error)
}
