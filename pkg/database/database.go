package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
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
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.dbname"),
		Port:     viper.GetString("db.port"),
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
