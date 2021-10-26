package interactor

/*
interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持ちます．

interactorはアウトプットポートに依存し(importするということ)，
インプットポートを実装します(interfaceを満たすようにmethodを追加するということ)．
*/

import (
	"context"
	"go-cla-mysql/entities/model"
	"go-cla-mysql/usecases/dto"
	"go-cla-mysql/usecases/port"
	"go-cla-mysql/usecases/repository"
)

type Todo struct {
	OutputPort port.TodoOutputPort
	TodoRepo   repository.TodoRepository
}

func (t *Todo) Create(ctx context.Context) model.Todos {
	panic("implement me")
}

// FindAll usecase.UserInputPortを実装している
// FindAll は，TodoRepo.GetUserByIDを呼び出し，その結果をOutputPort.Render or OutputPort.RenderErrorに渡します．
func (t *Todo) FindAll(ctx context.Context, max int) (*dto.TodoOutPutUseCaseDto, error) {
	// maxの設定
	const maxLimit int = 10
	todos, err := t.TodoRepo.FindAll(maxLimit)
	if err != nil {
		t.OutputPort.RenderError(err)
		return nil, nil
	}
	// TODO hitsを求めて、Dtoに変換して返す
	var hits = len(*todos)
	var todoOutPutUseCaseDto = dto.NewTodoOutPutUseCaseDto(hits, *todos)
	return todoOutPutUseCaseDto, nil
}

func (t *Todo) FindByID(ctx context.Context, id int) (model.Todos, error) {
	panic("implement me")
}

func (t *Todo) Update(ctx context.Context, todo model.Todo) (model.Todos, error) {
	panic("implement me")
}

// NewTodoInputPort はUserInputPortを取得します．（controllerで使用）
func NewTodoInputPort(outputPort port.TodoOutputPort, todoRepository repository.TodoRepository) *Todo {
	return &Todo{
		// repositoryをDI
		OutputPort: outputPort,
		TodoRepo:   todoRepository,
	}
}
