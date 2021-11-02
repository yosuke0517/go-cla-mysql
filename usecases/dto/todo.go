package dto

import "go-cla-practice/entities/model"

// 現状はtodoを配列にしただけだが、今後子供のtodoとかを入れたい要望があったときとかの場合にここに追加していく
type TodoOutPutUseCaseDto struct {
	Hits  int `json:"hits"`
	Todos model.Todos
}

func NewTodoOutPutUseCaseDto(hits int, todos []model.Todo) *TodoOutPutUseCaseDto {
	return &TodoOutPutUseCaseDto{
		Hits:  hits,
		Todos: todos,
	}
}
