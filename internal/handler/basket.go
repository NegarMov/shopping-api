package handler

import (
	"errors"
	"time"
	"log"
	"strconv"
	"net/http"
	"github.com/NegarMov/shopping-api/internal/model"
	"github.com/NegarMov/shopping-api/internal/request"
	"github.com/NegarMov/shopping-api/internal/store/basket"
	"github.com/labstack/echo/v4"
)

type Basket struct {
	Store basket.Basket
}

func (b Basket) GetAll(c echo.Context) error {
	baskets, err := b.Store.GetAll()
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, baskets)
}

func (b Basket) Create(c echo.Context) error {
	var req request.BasketCreate

	if err := c.Bind(&req); err != nil {
		log.Println(err)

		return echo.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Println(err)

		return echo.ErrBadRequest
	}

	new_b := model.Basket{
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Data: 		req.Data,
		State:		req.State,
	}

	created_b, err := b.Store.Create(new_b)

	if err != nil {
		log.Println(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, created_b)
}

func (b Basket) Update(c echo.Context) error {
	var req request.BasketUpdate
	id_string := c.Param("id")

	u64, parse_err := strconv.ParseUint(id_string, 10, 32)
    if parse_err != nil {
        log.Println(parse_err)

		return echo.ErrBadRequest
    }
    id := uint(u64)

	if err := c.Bind(&req); err != nil {
		log.Println(err)

		return echo.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Println(err)

		return echo.ErrBadRequest
	}

	new_b := model.Basket{
		UpdatedAt: 	time.Now(),
		Data: 		req.Data,
		State:		req.State,
	}

	updated_b, err := b.Store.Update(id, new_b)

	if err != nil {
		if errors.Is(err, basket.ErrBasketCompleted) {
			return echo.NewHTTPError(http.StatusBadRequest, basket.ErrBasketCompleted)
		}

		if errors.Is(err, basket.ErrBasketNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, basket.ErrBasketNotFound)
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, updated_b)
}

func (b Basket) Get(c echo.Context) error {
	id_string := c.Param("id")

	u64, parse_err := strconv.ParseUint(id_string, 10, 32)
    if parse_err != nil {
        log.Println(parse_err)

		return echo.ErrBadRequest
    }
    id := uint(u64)

	found_b, err := b.Store.Get(id)
	if err != nil {
		if errors.Is(err, basket.ErrBasketNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, basket.ErrBasketNotFound)
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, found_b)
}

func (b Basket) Delete(c echo.Context) error {
	id_string := c.Param("id")

	u64, parse_err := strconv.ParseUint(id_string, 10, 32)
    if parse_err != nil {
        log.Println(parse_err)

		return echo.ErrBadRequest
    }
    id := uint(u64)

	deleted_b, err := b.Store.Delete(id)
	if err != nil {
		if errors.Is(err, basket.ErrBasketNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, basket.ErrBasketNotFound)
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, deleted_b)
}

func (b Basket) Register(g *echo.Group) {
	g.GET("/basket", b.GetAll)
	g.POST("/basket", b.Create)
	g.PATCH("/basket/:id", b.Update)
	g.GET("/basket/:id", b.Get)
	g.DELETE("/basket/:id", b.Delete)
}