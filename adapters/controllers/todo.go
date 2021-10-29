package controllers

import (
	"fmt"
	"go-cla-mysql/adapters/presenter"
	"go-cla-mysql/infratructure/db"
	"go-cla-mysql/usecases/port"
	"go-cla-mysql/usecases/repository"
	"net/http"
	"strconv"
)

type Todo struct {
	OutputFactory func(w http.ResponseWriter) port.TodoOutputPort
	// -> presenter.NewTodoOutputPort
	InputFactory func(o port.TodoOutputPort, repo repository.TodoRepository) port.TodoInputPort
	// -> interactor.NewTodoInputPort
	RepoFactory func(c db.SqlHandler) repository.TodoRepository
	// -> gateway.NewTodoGateway
	Conn db.SqlHandler
}

// GetAll は，httpを受け取り，portを組み立てて，inputPort.FindAllを呼び出します．
func (t *Todo) GetAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		outputPort := t.OutputFactory(writer)
		repository := t.RepoFactory(t.Conn)
		inputPort := t.InputFactory(outputPort, repository)
		// TODO 第二引数はqueryから取得しなければデフォルトをセットする
		res, err := inputPort.FindAll(ctx, 10)
		if err != nil {
			presenter.InternalServerError(writer, fmt.Sprintf("cause: %s", err))
		}
		presenter.Success(writer, res)
	}
}

func (t *Todo) GetOne() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		strId := request.URL.Query().Get("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			presenter.BadRequest(writer, fmt.Sprintf("parameter invalid! cause: %s", err))
		}
		outputPort := t.OutputFactory(writer)
		repository := t.RepoFactory(t.Conn)
		inputPort := t.InputFactory(outputPort, repository)
		res, err := inputPort.FindByID(ctx, id)
		if err != nil {
			presenter.InternalServerError(writer, fmt.Sprintf("cause: %s", err))
		}
		presenter.Success(writer, res)
	}
}
