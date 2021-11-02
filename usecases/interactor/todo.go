package interactor

/*
interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持ちます．

interactorはアウトプットポートに依存し(importするということ)，
インプットポートを実装します(interfaceを満たすようにmethodを追加するということ)．
*/

import (
	"go-cla-practice/entities/model"
	"go-cla-practice/usecases/port"
	"go-cla-practice/usecases/repository"
	"net/http"
)

type Todo struct {
	OutputPort port.TodoOutputPort
	TodoRepo   repository.TodoRepository
}

// NewTodoInputPort はUserInputPortを取得します．（controllerで使用）
func NewTodoInputPort(outputPort port.TodoOutputPort, todoRepository repository.TodoRepository) port.TodoInputPort {
	return &Todo{
		OutputPort: outputPort,
		TodoRepo:   todoRepository,
	}
}

func (t *Todo) Create(writer http.ResponseWriter, todo *model.Todo) {
	_, err := t.TodoRepo.Create(todo)
	if err != nil {
		t.OutputPort.RenderError(writer, err)
	}
	var todos model.Todos
	todos = append(todos, *todo)
	t.OutputPort.Render(writer, &todos)
}

// FindAll usecase.UserInputPortを実装している
// FindAll は，TodoRepo.GetUserByIDを呼び出し，dtoに詰めて呼び出し元（controller）へ返します。
func (t *Todo) FindAll(writer http.ResponseWriter, max int) {
	// maxの設定
	const maxLimit int = 10
	todos, err := t.TodoRepo.FindAll(maxLimit)
	if err != nil {
		t.OutputPort.RenderError(writer, err)
	}
	t.OutputPort.Render(writer, todos)
}

func (t *Todo) FindByID(writer http.ResponseWriter, id int) {
	todos, err := t.TodoRepo.FindByID(id)
	if err != nil {
		t.OutputPort.RenderError(writer, err)
	}
	t.OutputPort.Render(writer, todos)
}

func (t *Todo) Update(writer http.ResponseWriter, todo *model.Todo) {
	_, err := t.TodoRepo.Update(todo)
	if err != nil {
		t.OutputPort.RenderError(writer, err)
	}
	var todos model.Todos
	todos = append(todos, *todo)
	t.OutputPort.Render(writer, &todos)
}
