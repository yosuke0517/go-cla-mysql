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
	TodoRepo repository.TodoRepository
}

// NewTodoInputPort はUserInputPortを取得します．（controllerで使用）
func NewTodoInputPort(todoRepository repository.TodoRepository) port.TodoInputPort {
	return &Todo{
		TodoRepo: todoRepository,
	}
}

func (t *Todo) Create(todo *model.Todo) (*dto.TodoOutPutUseCaseDto2, error) {
	isCreate, err := t.TodoRepo.Create(todo)
	var result = dto.NewTodoOutPutUseCaseDto2(isCreate)
	if err != nil {
		return result, err
	}
	return result, err
}

// FindAll usecase.UserInputPortを実装している
// FindAll は，TodoRepo.GetUserByIDを呼び出し，dtoに詰めて呼び出し元（controller）へ返します。
func (t *Todo) FindAll(max int) (*dto.TodoOutPutUseCaseDto, error) {
	// maxの設定
	const maxLimit int = 10
	todos, err := t.TodoRepo.FindAll(maxLimit)
	var hits = len(*todos)
	var todoOutPutUseCaseDto = dto.NewTodoOutPutUseCaseDto(hits, *todos)
	if err != nil {
		return nil, err
	}
	return todoOutPutUseCaseDto, err
}

func (t *Todo) FindByID(id int) (*dto.TodoOutPutUseCaseDto, error) {
	todos, err := t.TodoRepo.FindByID(id)
	var hits = len(*todos)
	var todoOutPutUseCaseDto = dto.NewTodoOutPutUseCaseDto(hits, *todos)
	if err != nil {
		return nil, err
	}
	return todoOutPutUseCaseDto, err
}

func (t *Todo) Update(todo *model.Todo) (*dto.TodoOutPutUseCaseDto2, error) {
	isUpdated, err := t.TodoRepo.Update(todo)
	var result = dto.NewTodoOutPutUseCaseDto2(isUpdated)
	if err != nil {

		return result, err
	}
	return result, err
}
