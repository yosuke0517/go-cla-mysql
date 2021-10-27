package controllers

import (
	"github.com/labstack/echo"
	"go-cla-mysql/infratructure/db"
	"go-cla-mysql/usecases/port"
	"go-cla-mysql/usecases/repository"
	"log"
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
func (t *Todo) GetAll() echo.HandlerFunc {
	return func(context echo.Context) error {
		ctx := context.Request().Context()
		outputPort := t.OutputFactory(context.Response().Writer)
		repository := t.RepoFactory(t.Conn)
		inputPort := t.InputFactory(outputPort, repository)
		// TODO 第二引数はqueryから取得しなければデフォルトをセットする
		hoge, err := inputPort.FindAll(ctx, 10)
		if err != nil {
			log.Fatalf("GetAll error cause, %s", err)
		}
		return context.JSON(http.StatusOK, hoge)
	}
}
