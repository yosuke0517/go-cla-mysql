package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-cla-mysql/entities/model"
	"go-cla-mysql/infratructure/db"
	"go-cla-mysql/usecases/repository"
	"log"
)

/*
gateway パッケージは，DB操作に対するアダプターです．
repositoryにて宣言されたメソッドを実装します.
*/

type TodoGateway struct {
	conn db.SqlHandler
}

// NewTodoGateway はTodoGatewayを返します．
// memo: 返り値の設定をrepository.TodoRepositoryにします.
// NewTodoGatewayを呼ぶ際、引数connにmockを渡せばテストが可能になります.
func NewTodoGateway(handler db.SqlHandler) repository.TodoRepository {
	return &TodoGateway{
		conn: handler,
	}
}

// GetDBConn はconnectionを取得します．
func (t *TodoGateway) GetDBConn() db.SqlHandler {
	return t.conn
}

func (t TodoGateway) FindAll(max int) (*model.Todos, error) {
	rows, err := t.conn.Conn.Query("SELECT * FROM `todo`")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Todo is not registered")
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapters/gateway/FindAll")
	}
	var todos model.Todos
	for rows.Next() {
		var todo model.Todo
		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)
		todos = append(todos, todo)
	}
	return &todos, nil
}

func (t TodoGateway) FindByID(ctx context.Context, id int) (*model.Todo, error) {
	row := t.conn.Conn.QueryRowContext(ctx, "SELECT * FROM `todo` WHERE id=?", id)
	todo := model.Todo{}
	err := row.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Todo Not Found. ID = %s", id)
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapters/gateway/GetTodoByID")
	}
	return &todo, nil
}

func (t TodoGateway) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	cmd := fmt.Sprintf("INSERT INTO %s (id, task, limitdate, status) VALUES (?, ?, ?, ?)", `todo`)
	ins, err := t.conn.Conn.Prepare(cmd)
	if err != nil {
		log.Println(err)
	}
	if ins != nil {
		_, err = ins.Exec(todo.ID, todo.Task, todo.LimitDate, todo.Status)
		if err != nil {
			log.Println(err)
		}
	}
	return todo, nil
}

func (t TodoGateway) Update(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	panic("implement me")
}
