package activityRepository

import (
	"fmt"
	"log"
	activity_model "test_service/internal/models/activity"
	activityRepository "test_service/internal/repository/activity"

	msgErrors "test_service/pkg/errors"

	"github.com/jmoiron/sqlx"
)

const (
	getAllActivity  = `SELECT * from activity`
	GetActivityById = `SELECT * from activity WHERE id = ?`
	SaveActivity    = `INSERT INTO activity (email, title, created_at) VALUES ('$1', '$2', now())`
	LastInsertId    = `SELECT last_insert_id()`
	DeletedById     = `DELETE FROM activity WHERE id = ?`
	UpdateActivity  = `UPDATE activity SET title = '$1', updated_at = now() WHERE id = ?`
)

var statement PreparedStatement

//persiapan mengambil data dari database
type PreparedStatement struct {
	GetAllActivity  *sqlx.Stmt
	GetActivityById *sqlx.Stmt
	SaveActivity    *sqlx.Stmt
	DeletedById     *sqlx.Stmt
	UpdateActivity  *sqlx.Stmt
}

type ActivityRepo struct {
	Conn *sqlx.DB
}

//prepare isi InitPrepare dengan isian Conn, func diberi isian
func NewRepo(Conn *sqlx.DB) activityRepository.ActivityRepository {

	repo := &ActivityRepo{Conn}
	InitPreparedStatement(repo)
	return repo
}

func (p *ActivityRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Conn.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *ActivityRepo) {
	statement = PreparedStatement{
		GetAllActivity:  m.Preparex(getAllActivity),
		GetActivityById: m.Preparex(GetActivityById),
		SaveActivity:    m.Preparex(SaveActivity),
		DeletedById:     m.Preparex(DeletedById),
		UpdateActivity:  m.Preparex(UpdateActivity),
	}
}

func (p *ActivityRepo) GetAllActivity() ([]*activity_model.ActivityModel, error) {

	//fmt.Println(id, "here")
	// activity := []*activity_model.ActivityModel{}
	var activity []*activity_model.ActivityModel
	err := statement.GetAllActivity.Select(&activity)

	if err != nil {
		log.Println("error :", err)
		return nil, fmt.Errorf(err.Error())
	}

	return activity, nil
}

func (p *ActivityRepo) GetActivityById(id int64) (*activity_model.ActivityModel, error) {

	dtActivity := []*activity_model.ActivityModel{}
	err := statement.GetActivityById.Select(&dtActivity, id)

	if err != nil {
		log.Println("Failed Query GetMAhasiswa By ID : ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	if len(dtActivity) < 1 {
		return nil, fmt.Errorf(msgErrors.ErrorDataNotFound)
	}
	return dtActivity[0], nil
}

func (p *ActivityRepo) SaveActivity(payload *activity_model.ActivityModel) (int64, error) {

	var id int64
	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx save activity : ", err.Error())
		return 0, fmt.Errorf(msgErrors.ErrorDB)
	}

	stmt := fmt.Sprintf(`INSERT INTO activity (email, title, created_at) VALUES ('%s', '%s', now())`, payload.Email, payload.Title)

	_, err = tx.Exec(stmt)

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query save activity: ", err.Error())
		return 0, fmt.Errorf(msgErrors.ErrorDB)
	}

	err = tx.QueryRow(LastInsertId).Scan(&id)
	err = tx.Commit()
	if err != nil {
		log.Println("Failed end Tx save activity : ", err.Error())
		return 0, fmt.Errorf(msgErrors.ErrorDB)
	}

	return id, nil
}

func (p *ActivityRepo) DeleteActivityById(id int64) error {

	_, err := statement.DeletedById.Exec(id)

	if err != nil {
		log.Println("Failed Query delete By ID : ", err)
		return err
	}

	return nil
}

func (p *ActivityRepo) UpdateActivity(id int64, title string) error {

	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx update activity : ", err.Error())
		return fmt.Errorf(msgErrors.ErrorDB)
	}

	stmt := fmt.Sprintf(`UPDATE activity SET title = '%s', updated_at = now() WHERE id = ?`, title)

	_, err = tx.Exec(stmt, id)

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query update activity: ", err.Error())
		return fmt.Errorf(msgErrors.ErrorDB)
	}
	return tx.Commit()
}
