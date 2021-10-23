package port

import "go-cla-mysql/entities/model"

// 現状はtodoを配列にしただけだが、今後子供のtodoとかを入れたい要望があったときとかの場合にここに追加していく
type TodosUseCaseDto struct {
	Todos []model.Todo
}

type TodoOutPutUseCaseDto struct {
	Hits  int `json:"hits"`
	Todos TodosUseCaseDto
}

type TodoInputPort interface {
	Create() TodosUseCaseDto
	FindAll() TodoOutPutUseCaseDto
	FindByID(id int) (TodosUseCaseDto, error)
	// Save(domain.Todoなどの単数系) (int64, error)
	// DeleteByID(int) error
}

type TodoOutputPort interface {
	Render(*model.Todo)
	RenderError(error)
}
