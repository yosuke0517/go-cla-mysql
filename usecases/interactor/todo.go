package interactor

/*
interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持ちます．

interactorはアウトプットポートに依存し(importするということ)，
インプットポートを実装します(interfaceを満たすようにmethodを追加するということ)．
*/

import (
	"go-cla-practice/entities/model"
	"go-cla-practice/usecases/dto"
	"go-cla-practice/usecases/port"
	"go-cla-practice/usecases/repository"
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

func (t *Todo) Create(todo *model.Todo) (bool, error) {
	isCreated, err := t.TodoRepo.Create(todo)
	return isCreated, err
}

// FindAll usecase.UserInputPortを実装している
// FindAll は，TodoRepo.GetUserByIDを呼び出し，dtoに詰めて呼び出し元（controller）へ返します。
func (t *Todo) FindAll(max int) (*dto.TodoOutPutUseCaseDto, error) {
	// maxの設定
	const maxLimit int = 10
	todos, err := t.TodoRepo.FindAll(maxLimit)
	if err != nil {
		t.OutputPort.RenderError(err)
		return nil, nil
	}
	var hits = len(*todos)
	var todoOutPutUseCaseDto = dto.NewTodoOutPutUseCaseDto(hits, *todos)
	return todoOutPutUseCaseDto, nil
}

func (t *Todo) FindByID(id int) (*dto.TodoOutPutUseCaseDto, error) {
	todo, err := t.TodoRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	var hits = len(*todo)
	var todoOutPutUseCaseDto = dto.NewTodoOutPutUseCaseDto(hits, *todo)
	return todoOutPutUseCaseDto, nil
}

func (t *Todo) Update(todo *model.Todo) (bool, error) {
	isCreated, err := t.TodoRepo.Update(todo)
	return isCreated, err
}
