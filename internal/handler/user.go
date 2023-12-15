package handler

import (
	"errors"
	"time"
	"log"
	"net/http"
	"github.com/NegarMov/shopping-api/internal/model"
	"github.com/NegarMov/shopping-api/internal/request"
	"github.com/NegarMov/shopping-api/internal/store/user"
	"github.com/NegarMov/shopping-api/internal/auth"
	"github.com/NegarMov/shopping-api/configs"
	"github.com/labstack/echo/v4"
)

type User struct {
	Store user.User
}

func (u User) SignUp(c echo.Context) error {
	var req request.Credentials

	if err := c.Bind(&req); err != nil {
		log.Println(err)

		return echo.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Println(err)

		return echo.ErrBadRequest
	}

	new_u := model.User{
		Username: 	req.Username,
		Password: 	req.Password,
		CreatedAt: 	time.Now(),
	}

	created_u, err := u.Store.SignUp(new_u)

	if err != nil {
		if errors.Is(err, user.ErrDuplicatedUsername) {
			return echo.NewHTTPError(http.StatusBadRequest, user.ErrDuplicatedUsername)
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, created_u)
}

func (u User) SignIn(c echo.Context) error {
	var req request.Credentials

	if err := c.Bind(&req); err != nil {
		log.Println(err)

		return echo.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Println(err)

		return echo.ErrBadRequest
	}

	found_u, err := u.Store.SignIn(req.Username, req.Password)

	if err != nil {
		if errors.Is(err, user.ErrWrongCredentials) {
			return echo.NewHTTPError(http.StatusUnauthorized, user.ErrWrongCredentials)
		}

		return echo.ErrInternalServerError
	}

	config, _ := configs.LoadConfig()

	token, err := auth.CreateAccessToken(found_u, config.JWTSecret, config.TokenExpiry)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string]string{
		"jwt_token": token,
	})
}

func (u User) Register(g *echo.Group) {
	g.POST("/signup", u.SignUp)
	g.POST("/signin", u.SignIn)
}