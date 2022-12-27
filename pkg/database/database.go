package database

import (
	"fmt"

	cnf "test_service/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Conn *sqlx.DB
}

type Config struct {
	Host     string
	Username string
	Password string
	Database string
	Port     string
}

func MakeInitialize() *Config {
	return &Config{
		Host:     cnf.EnvConfigs.DbHost,
		Username: cnf.EnvConfigs.DbUsername,
		Password: cnf.EnvConfigs.DbPassword,
		Database: cnf.EnvConfigs.DbName,
		Port:     cnf.EnvConfigs.DbPort,
	}
}

func Initialize(config *Config) (Database, error) {
	db := Database{}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Port, config.Database)
	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
