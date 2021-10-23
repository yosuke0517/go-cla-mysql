package interactor

/*
interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持ちます．

interactorはアウトプットポートに依存し(importするということ)，
インプットポートを実装します(interfaceを満たすようにmethodを追加するということ)．
*/

import (
	"go-cla-mysql/usecases/port"
	"go-cla-mysql/usecases/repository"
)

type Todo struct {
	OutputPort port.TodoOutputPort
	TodoRepo   repository.TodoRepository
}

// NewUserInputPort はUserInputPortを取得します．（controllerで使用）
func NewTodoInputPort(outputPort port.TodoOutputPort, todoRepository repository.TodoRepository) *Todo {
	return &Todo{
		// outputPort, repositoryをDI
		OutputPort: outputPort,
		TodoRepo:   todoRepository,
	}
}

// usecase.UserInputPortを実装している
// FindAll は，TodoRepo.GetUserByIDを呼び出し，その結果をOutputPort.Render or OutputPort.RenderErrorに渡します．
func (t *Todo) FindAll(max int) {
	// maxの設定
	const maxLimit int = 10
	todos, err := t.TodoRepo.FindAll(maxLimit)
	if err != nil {
		t.OutputPort.RenderError(err)
		return
	}
	// TODO hitsを求めて、Dtoに変換して返す
	//var hits = len(todos)
	//var todoOutPutUseCaseDto = port.NewTodoOutPutUseCaseDto(hits, todos)
	t.OutputPort.Render(todos)
}
