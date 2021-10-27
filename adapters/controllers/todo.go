package controllers

import (
	"go-cla-mysql/infratructure/db"
	"go-cla-mysql/usecases/port"
	"go-cla-mysql/usecases/repository"
	"net/http"
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
func (t *Todo) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	outputPort := t.OutputFactory(w)
	repository := t.RepoFactory(t.Conn)
	inputPort := t.InputFactory(outputPort, repository)
	// TODO 第二引数はqueryから取得しなければデフォルトをセットする
	inputPort.FindAll(ctx, 10)
}
