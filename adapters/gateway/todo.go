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
	db.SqlHandler
}

// NewTodoGateway はTodoGatewayを返します．
// memo: 返り値の設定をrepository.TodoRepositoryにします.
// NewTodoGatewayを呼ぶ際、引数connにmockを渡せばテストが可能になります.
func NewTodoGateway(handler db.SqlHandler) repository.TodoRepository {
	return &TodoGateway{
		handler,
	}
}

func (t TodoGateway) FindAll(ctx context.Context, max int) (*model.Todos, error) {
	tableName := "todo"
	cmd := fmt.Sprintf(`SELECT * FROM %s`, tableName)
	rows, err := t.Conn.Query(cmd)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Todo is not registered")
		}
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Internal Server Error. adapters/gateway/FindAll: %s", err))
	}
	var todos model.Todos
	for rows.Next() {
		var todo model.Todo
		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)
		todos = append(todos, todo)
	}
	return &todos, nil
}

// 1件なんだけど返却はFindAllと揃えた
func (t TodoGateway) FindByID(ctx context.Context, id int) (*model.Todos, error) {
	tableName := "todo"
	cmd := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	row := t.Conn.QueryRow(cmd, id)
	todo := model.Todo{}
	err := row.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Todo Not Found. ID = %s", id)
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapters/gateway/GetTodoByID")
	}
	var todos model.Todos
	todos = append(todos, todo)
	return &todos, nil
}

func (t TodoGateway) Create(todo *model.Todo) (bool, error) {
	cmd := fmt.Sprintf("INSERT INTO %s (id, task, limitdate, status) VALUES (?, ?, ?, ?)", `todo`)
	ins, err := t.Conn.Prepare(cmd)
	if err != nil {
		return false, err
	}
	if ins != nil {
		_, err = ins.Exec(todo.ID, todo.Task, todo.LimitDate, todo.Status)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (t TodoGateway) Update(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	panic("implement me")
}
