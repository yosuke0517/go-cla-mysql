package port

import "go-cla-mysql/entities/model"

// 現状はtodoを配列にしただけだが、今後子供のtodoとかを入れたい要望があったときとかの場合にここに追加していく
type TodoOutPutUseCaseDto struct {
	Hits  int `json:"hits"`
	Todos []model.Todo
}

func NewTodoOutPutUseCaseDto(hits int, todos []model.Todo) *TodoOutPutUseCaseDto {
	return &TodoOutPutUseCaseDto{
		Hits:  hits,
		Todos: todos,
	}
}

type TodoInputPort interface {
	Create() model.Todo
	FindAll(max int) model.Todo
	FindByID(id int) (model.Todo, error)
	Update(todo model.Todo) (model.Todo, error)
}

type TodoOutputPort interface {
	Render(todos []*model.Todo)
	RenderError(error)
}
