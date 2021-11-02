package injector

import (
	"go-cla-practice/adapters/controllers"
	"go-cla-practice/adapters/gateway"
	"go-cla-practice/adapters/presenter"
	"go-cla-practice/infratructure/db"
	"go-cla-practice/usecases/interactor"
	"go-cla-practice/usecases/port"
	"go-cla-practice/usecases/repository"
)

func InjectDB() db.SqlHandler {
	sqlhandler := db.NewSqlHandler()
	return *sqlhandler
}

func InjectTodoGateway() repository.TodoRepository {
	sqlHandler := InjectDB()
	return gateway.NewTodoGateway(sqlHandler)
}

func InjectOutPutPort() port.TodoOutputPort {
	return presenter.NewTodoOutputPort()
}

func InjectInputPort() port.TodoInputPort {
	OutPutPort := InjectOutPutPort()
	TodoRepo := InjectTodoGateway()
	return interactor.NewTodoInputPort(OutPutPort, TodoRepo)
}

func InjectTodoHandler() controllers.TodoHandler {
	return controllers.NewTodoHandler(InjectInputPort())
}
