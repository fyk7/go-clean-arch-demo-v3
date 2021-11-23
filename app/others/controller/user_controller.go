package controller

import (
	"errors"
	"fmt"
	"net/http"

	interactor "github.com/fyk7/go-clean-arch-demo-v3/app/usecase/interactor"
	"github.com/labstack/echo/v4"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

// TODO ginのContextを直接埋め込んでしまう。他のファイルにinterfaceは定義しない
type UserController interface {
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
}

func NewUserController(e *echo.Echo, us interactor.UserInteractor) {
	controller := &userController{
		userInteractor: us,
	}
	e.GET("/users", controller.GetUsers)
	e.POST("/user", controller.CreateUser)
	e.GET("/user/:email", controller.GetByEmail)
}

func (uc *userController) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := uc.userInteractor.ListUser(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) GetByEmail(c echo.Context) error {
	type (
		GetUserResponse struct {
			ID    string
			Email string
		}
	)
	ctx := c.Request().Context()
	// TODO validate email in controller or usecase
	email := c.Param("email")

	u, err := uc.userInteractor.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	user_resp := GetUserResponse{
		ID:    u.ID,
		Email: u.Email,
	}
	return c.JSON(http.StatusOK, user_resp)

}

func (uc *userController) CreateUser(c echo.Context) error {
	type (
		// TODO Define validator
		CreateUserRequest struct {
			Email string `json:"email"`
		}
		CreateUserResponse struct {
			ID    string
			Email string
		}
	)
	ctx := c.Request().Context()
	var params CreateUserRequest

	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	user, err := uc.userInteractor.Create(ctx, params.Email)
	if !errors.Is(err, nil) {
		fmt.Println(err)
		// TODO judge error
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}
	user_resp := CreateUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	return c.JSON(http.StatusCreated, user_resp)
}
