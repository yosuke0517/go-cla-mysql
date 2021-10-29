package injector

import (
	"go-cla-mysql/adapters/controllers"
	"go-cla-mysql/adapters/gateway"
	"go-cla-mysql/adapters/presenter"
	"go-cla-mysql/infratructure/db"
	"go-cla-mysql/usecases/interactor"
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
