package port

import (
	"go-cla-practice/entities/model"
	"go-cla-practice/usecases/dto"
)

type TodoInputPort interface {
	Create(todo *model.Todo) (*dto.TodoOutPutUseCaseDto2, error)
	FindAll(max int) (*dto.TodoOutPutUseCaseDto, error)
	FindByID(id int) (*dto.TodoOutPutUseCaseDto, error)
	Update(todo *model.Todo) (*dto.TodoOutPutUseCaseDto2, error)
}

// TODO intreractorの中でoutputport呼ぶだけにしたいけどresponseWriter渡さないといけなくて、テストしづらいので断念
// TodoOutputPort この中で共通のレスポンス処理を呼ぶ
//type TodoOutputPort interface {
//	Render(todos *model.Todos)
//	RenderError(err error)
//}
