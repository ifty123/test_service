package activityRepository

import (
	"fmt"
	"log"
	activity_model "test_service/internal/models/activity"
	activityRepository "test_service/internal/repository/activity"

	"github.com/jmoiron/sqlx"
)

const (
	getAllActivity = `SELECT id, title, email, created_at, updated_at from activity`
)

var statement PreparedStatement

//persiapan mengambil data dari database
type PreparedStatement struct {
	getAllActivity *sqlx.Stmt
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
		getAllActivity: m.Preparex(getAllActivity),
	}
}

func (p *ActivityRepo) GetAllActivity() ([]*activity_model.ActivityModel, error) {

	//fmt.Println(id, "here")
	// activity := []*activity_model.ActivityModel{}
	var activity []*activity_model.ActivityModel
	_, err := statement.getAllActivity.Query(&activity)

	if err != nil {
		log.Println("error :", err)
		return nil, fmt.Errorf(err.Error())
	}

	return activity, nil
}
