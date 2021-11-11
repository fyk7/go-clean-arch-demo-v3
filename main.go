package main

import (
	"database/sql"
	"log"
	"time"

	_db "github.com/fyk7/go-clean-arch-demo-v3/app/database"
	_userService "github.com/fyk7/go-clean-arch-demo-v3/app/domain/service"
	_userController "github.com/fyk7/go-clean-arch-demo-v3/app/others/controller"
	_middleware "github.com/fyk7/go-clean-arch-demo-v3/app/others/controller/middleware"
	_userRepository "github.com/fyk7/go-clean-arch-demo-v3/app/others/repository"
	_userInteractor "github.com/fyk7/go-clean-arch-demo-v3/app/usecase/interactor"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	// Dependencies Injection
	var dbConn *sql.DB = _db.NewDB()
	err := dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	timeoutContext := time.Duration(2) * time.Second
	userRepository := _userRepository.NewSqlUserRepository(dbConn)
	userService := _userService.NewUserService(userRepository)
	userUsecase := _userInteractor.NewUserInteractor(userRepository, userService, timeoutContext)

	e := echo.New()
	mdlWare := _middleware.InitMiddleware()
	e.Use(mdlWare.CORS)
	_userController.NewUserController(e, userUsecase)
	log.Fatal(e.Start("localhost:9090"))
}
