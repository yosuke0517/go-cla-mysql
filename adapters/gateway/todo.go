package gateway

import (
	"database/sql"
	"errors"
	"fmt"
	"go-cla-practice/entities/model"
	"go-cla-practice/infratructure/db"
	"go-cla-practice/usecases/repository"
	"log"
)

/*
gateway パッケージは，DB操作に対するアダプターです．
repositoryにて宣言されたメソッドを実装します.
*/

var tableName = "todos"

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

func (t TodoGateway) FindAll(max int) (*model.Todos, error) {
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
		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status, &todo.Deleted)
		todos = append(todos, todo)
	}
	return &todos, nil
}

// 1件なんだけど返却はFindAllと揃えた
func (t TodoGateway) FindByID(id int) (*model.Todos, error) {
	cmd := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	row := t.Conn.QueryRow(cmd, id)
	todo := model.Todo{}
	err := row.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status, &todo.Deleted)
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

func (t TodoGateway) Create(todo *model.Todo) (*model.Todos, error) {
	cmd := fmt.Sprintf("INSERT INTO %s (id, task, limitdate, status) VALUES (?, ?, ?, ?)", `todos`)
	ins, err := t.Conn.Prepare(cmd)
	var todos model.Todos
	if err != nil {
		return &todos, err
	}
	if ins != nil {
		_, err = ins.Exec(todo.ID, todo.Task, todo.LimitDate, todo.Status)
		if err != nil {
			return &todos, err
		}
	}
	todos = append(todos, *todo)
	return &todos, nil
}

func (t TodoGateway) Update(todo *model.Todo) (*model.Todos, error) {
	cmd := fmt.Sprintf("UPDATE %s SET task = ?, limitDate = ?, status = ?, deleted = ? WHERE id = ?", tableName)
	upd, err := t.Conn.Prepare(cmd)
	var todos model.Todos
	if err != nil {
		return &todos, err
	}
	if upd != nil {
		_, err = upd.Exec(todo.Task, todo.LimitDate, todo.Status, todo.Deleted, todo.ID)
		if err != nil {
			return &todos, err
		}
	}
	todos = append(todos, *todo)
	return &todos, nil
}
