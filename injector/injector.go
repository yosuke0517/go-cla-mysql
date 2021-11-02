package injector

import (
	"go-cla-practice/adapters/controllers"
	"go-cla-practice/adapters/gateway"
	"go-cla-practice/adapters/presenter"
	"go-cla-practice/infratructure/db"
	"go-cla-practice/usecases/interactor"
)

func InjectDB() db.SqlHandler {
	sqlhandler := db.NewSqlHandler()
	return *sqlhandler
}

func InjectTodo() controllers.Todo {
	todo := controllers.Todo{
		OutputFactory: presenter.NewTodoOutputPort,
		InputFactory:  interactor.NewTodoInputPort,
		RepoFactory:   gateway.NewTodoGateway,
		Conn:          InjectDB(),
	}
	return todo
}
