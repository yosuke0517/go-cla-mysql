package controllers

import (
	"encoding/json"
	"fmt"
	"go-cla-practice/adapters/presenter"
	"go-cla-practice/entities/model"
	"go-cla-practice/usecases/port"
	"net/http"
	"strconv"
)

type Todo struct {
	HandlerFactory func(w http.ResponseWriter) TodoHandler
}

type TodoHandler struct {
	todoInputPort port.TodoInputPort
}

func NewTodoHandler(todoInputPort port.TodoInputPort) TodoHandler {
	todoHandler := TodoHandler{todoInputPort: todoInputPort}
	return todoHandler
}

// GetAll は，httpを受け取り，portを組み立てて，inputPort.FindAllを呼び出します．
func (t *TodoHandler) GetAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		res, err := t.todoInputPort.FindAll(writer, 10)
	}
}

func (t *TodoHandler) GetOne() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		strId := request.URL.Query().Get("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			presenter.BadRequest(writer, fmt.Sprintf("parameter invalid! cause: %s", err))
		}
		res, err := t.todoInputPort.FindByID(writer, id)
	}
}

func (t *TodoHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := &model.Todo{}
		if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
			presenter.BadRequest(writer, "Bad request: "+err.Error())
			return
		}
		// TODO validation追加
		res, err := t.todoInputPort.Create(data)
		if err != nil {
			presenter.InternalServerError(writer, "InternalServerError: "+err.Error())
		}
		presenter.Success(writer, res)
	}
}

func (t *TodoHandler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := &model.Todo{}
		if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
			presenter.BadRequest(writer, "Bad request: "+err.Error())
			return
		}
		// TODO validation追加
		res, err := t.todoInputPort.Update(writer, data)
	}
}
