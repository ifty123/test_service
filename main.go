package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	cnf "test_service/config"
	"test_service/pkg/database"

	activityRepository "test_service/internal/repository/activity/activity_repo"
	activityService "test_service/internal/services/activity_service"
	activityHandler "test_service/internal/transport/http/activity"

	todoRepository "test_service/internal/repository/todo/todo_repo"
	todoService "test_service/internal/services/todo_service"
	todoHandler "test_service/internal/transport/http/todo"

	"test_service/internal/transport/http/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {

	errChan := make(chan error)

	e := echo.New()
	m := middleware.NewMidleware()

	cnf.InitEnvConfigs()

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
	activityrepo := activityRepository.NewRepo(db.Conn)
	srvAct := activityService.NewService(activityrepo)
	activityHandler.NewHttpHandler(e, srvAct)

	//repo todo
	todorepo := todoRepository.NewRepo(db.Conn)
	srvTodo := todoService.NewService(todorepo)
	todoHandler.NewHttpHandler(e, srvTodo, srvAct)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errChan <- e.Start(":" + cnf.EnvConfigs.LocalServerPort)
	}()

	e.Logger.Print("Starting ", cnf.EnvConfigs.AppName)
	err = <-errChan
	log.Error(err.Error())

}
