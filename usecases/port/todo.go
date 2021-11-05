package port

import (
	"go-cla-practice/entities/model"
)

type TodoInputPort interface {
	Create(todo *model.Todo)
	FindAll(max int)
	FindByID(id int)
	Update(todo *model.Todo)
}

// TodoOutputPort この中で共通のレスポンス処理を呼ぶ
type TodoOutputPort interface {
	Render(todos *model.Todos)
	RenderError(err error)
}
