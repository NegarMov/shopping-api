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

func str2uint(s string) (uint, error) {
	u64, parse_err := strconv.ParseUint(s, 10, 32)
    if parse_err != nil {
        log.Println(parse_err)
		return 0, parse_err
    }
    return uint(u64), nil
}

type Basket struct {
	Store basket.Basket
}

func (b Basket) GetAll(c echo.Context) error {
	user_id, parse_err := str2uint(c.Get("x-user-id").(string))
    if parse_err != nil {
        return echo.ErrInternalServerError
    }

	baskets, err := b.Store.GetAll(user_id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, baskets)
}

func (b Basket) Create(c echo.Context) error {
	user_id, parse_err := str2uint(c.Get("x-user-id").(string))
    if parse_err != nil {
        return echo.ErrInternalServerError
    }

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
		UserID: 	user_id,
	}

	created_b, err := b.Store.Create(new_b)

	if err != nil {
		log.Println(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, created_b)
}

func (b Basket) Update(c echo.Context) error {
	user_id, parse_err := str2uint(c.Get("x-user-id").(string))
    if parse_err != nil {
        return echo.ErrInternalServerError
    }

	var req request.BasketUpdate

	basket_id, parse_err := str2uint(c.Param("id"))
    if parse_err != nil {
        return echo.ErrBadRequest
    }

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
		UserID: 	user_id,
	}

	updated_b, err := b.Store.Update(basket_id, new_b)

	if err != nil {
		if errors.Is(err, basket.ErrBasketCompleted) {
			return echo.NewHTTPError(http.StatusBadRequest, basket.ErrBasketCompleted)
		}

		if errors.Is(err, basket.ErrBasketNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, basket.ErrBasketNotFound)
		}

		if errors.Is(err, basket.ErrAccessDenied) {
			return echo.NewHTTPError(http.StatusNotFound, basket.ErrAccessDenied)
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, updated_b)
}

func (b Basket) Get(c echo.Context) error {
	user_id, parse_err := str2uint(c.Get("x-user-id").(string))
    if parse_err != nil {
        return echo.ErrInternalServerError
    }

	basket_id, parse_err := str2uint(c.Param("id"))
    if parse_err != nil {
        return echo.ErrBadRequest
    }

	found_b, err := b.Store.Get(basket_id, user_id)
	if err != nil {
		if errors.Is(err, basket.ErrBasketNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, basket.ErrBasketNotFound)
		}

		if errors.Is(err, basket.ErrAccessDenied) {
			return echo.NewHTTPError(http.StatusNotFound, basket.ErrAccessDenied)
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, found_b)
}

func (b Basket) Delete(c echo.Context) error {
	user_id, parse_err := str2uint(c.Get("x-user-id").(string))
    if parse_err != nil {
        return echo.ErrInternalServerError
    }

	basket_id, parse_err := str2uint(c.Param("id"))
    if parse_err != nil {
        return echo.ErrBadRequest
    }

	deleted_b, err := b.Store.Delete(basket_id, user_id)
	if err != nil {
		if errors.Is(err, basket.ErrBasketNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, basket.ErrBasketNotFound)
		}

		if errors.Is(err, basket.ErrAccessDenied) {
			return echo.NewHTTPError(http.StatusNotFound, basket.ErrAccessDenied)
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, deleted_b)
}

func (b Basket) Register(g *echo.Group) {
	g.GET("", b.GetAll)
	g.POST("", b.Create)
	g.PATCH("/:id", b.Update)
	g.GET("/:id", b.Get)
	g.DELETE("/:id", b.Delete)
}