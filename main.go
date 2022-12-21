package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test_service/pkg/database"

	activityRepository "test_service/internal/repository/activity/activity_repo"
	activityService "test_service/internal/services/activity_service"
	activityHandler "test_service/internal/transport/http/activity"

	"test_service/internal/transport/http/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {

	errChan := make(chan error)

	e := echo.New()
	m := middleware.NewMidleware()

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config-dev")

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	connect := database.MakeInitialize()

	db, err := database.Initialize(connect)
	if err != nil {
		log.Fatal("Failed to Connect Mysql Database: " + err.Error())
	}

	defer func() {
		err := db.Conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	e.Use(m.CORS)

	//repo activity
	sqlrepo := activityRepository.NewRepo(db.Conn)
	srv := activityService.NewService(sqlrepo)
	activityHandler.NewHttpHandler(e, srv)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errChan <- e.Start(":" + viper.GetString("server.port"))
	}()

	e.Logger.Print("Starting ", viper.GetString("appName"))
	err = <-errChan
	log.Error(err.Error())

}
