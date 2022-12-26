package todoRepository

import (
	"fmt"
	"log"
	todo_model "test_service/internal/models/todo"
	todoRepository "test_service/internal/repository/todo"

	msgErrors "test_service/pkg/errors"

	"github.com/jmoiron/sqlx"
)

const (
	getAllTodo   = `SELECT * from todo WHERE activity_group_id = ?`
	GetTodoById  = `SELECT * from todo WHERE id = ?`
	SaveTodo     = `INSERT INTO todo (activity_group_id, is_active, title,priority, created_at) VALUES ('$1', '$2', '$3', '$4', now())`
	LastInsertId = `SELECT last_insert_id()`
	DeletedById  = `UPDATE todo SET deleted_at = now() WHERE id = ?`
	UpdateTodo   = `UPDATE todo SET title = '$1',  priority = '$2', is_active = '$3', updated_at = now() WHERE id = ?`
)

var statement PreparedStatement

//persiapan mengambil data dari database
type PreparedStatement struct {
	GetAllTodo  *sqlx.Stmt
	GetTodoById *sqlx.Stmt
	SaveTodo    *sqlx.Stmt
	DeletedById *sqlx.Stmt
	UpdateTodo  *sqlx.Stmt
}

type TodoRepo struct {
	Conn *sqlx.DB
}

//prepare isi InitPrepare dengan isian Conn, func diberi isian
func NewRepo(Conn *sqlx.DB) todoRepository.TodoRepository {

	repo := &TodoRepo{Conn}
	InitPreparedStatement(repo)
	return repo
}

func (p *TodoRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Conn.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *TodoRepo) {
	statement = PreparedStatement{
		GetAllTodo:  m.Preparex(getAllTodo),
		GetTodoById: m.Preparex(GetTodoById),
		SaveTodo:    m.Preparex(SaveTodo),
		DeletedById: m.Preparex(DeletedById),
		UpdateTodo:  m.Preparex(UpdateTodo),
	}
}

func (p *TodoRepo) GetAllTodo(id int64) ([]*todo_model.TodoModel, error) {

	var Todo []*todo_model.TodoModel
	err := statement.GetAllTodo.Select(&Todo, id)

	if err != nil {
		log.Println("error :", err)
		return nil, fmt.Errorf(err.Error())
	}

	return Todo, nil
}

func (p *TodoRepo) GetTodoById(id int64) (*todo_model.TodoModel, error) {

	dtTodo := []*todo_model.TodoModel{}
	err := statement.GetTodoById.Select(&dtTodo, id)

	if err != nil {
		log.Println("Failed Query todo By ID : ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	if len(dtTodo) < 1 {
		return nil, fmt.Errorf(msgErrors.ErrorDataNotFound)
	}
	return dtTodo[0], nil
}

func (p *TodoRepo) SaveTodo(payload *todo_model.TodoModel) (int64, error) {

	var id int64
	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx save Todo : ", err.Error())
		return 0, fmt.Errorf(msgErrors.ErrorDB)
	}

	isActive := 0
	if payload.IsActive {
		isActive = 1
	}

	stmt := fmt.Sprintf(`INSERT INTO todo (activity_group_id, is_active, title, priority, created_at) VALUES (%d, %d, '%s',  '%s', now())`, payload.ActivityGroupId, isActive, payload.Title, payload.Priority)

	_, err = tx.Exec(stmt)

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query save Todo: ", err.Error())
		return 0, fmt.Errorf(msgErrors.ErrorDB)
	}

	err = tx.QueryRow(LastInsertId).Scan(&id)
	err = tx.Commit()
	if err != nil {
		log.Println("Failed end Tx save Todo : ", err.Error())
		return 0, fmt.Errorf(msgErrors.ErrorDB)
	}

	return id, nil
}

func (p *TodoRepo) DeleteTodoById(id int64) error {

	_, err := statement.DeletedById.Exec(id)

	if err != nil {
		log.Println("Failed Query todo By ID : ", err)
		return err
	}

	return nil
}

func (p *TodoRepo) UpdateTodo(payload *todo_model.TodoModel) error {

	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx update Todo : ", err.Error())
		return fmt.Errorf(msgErrors.ErrorDB)
	}

	stmt := fmt.Sprintf(`UPDATE todo SET title = '%s',  priority = '%s', is_active = '%t', updated_at = now() WHERE id = ?`, payload.Title, payload.Priority, payload.IsActive)

	_, err = tx.Exec(stmt, payload.Id)

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query update Todo: ", err.Error())
		return fmt.Errorf(msgErrors.ErrorDB)
	}
	return tx.Commit()
}
