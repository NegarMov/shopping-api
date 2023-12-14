package main

import (
	"log"
	"github.com/NegarMov/shopping-api/internal/handler"
	"github.com/NegarMov/shopping-api/internal/store/basket"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := echo.New()

	db, err := gorm.Open(sqlite.Open("shop.db"), new(gorm.Config))
	if err != nil {
		log.Fatal(err)
	}

	{
		b := basket.NewSQL(db)
		h := handler.Basket{
			Store: b,
		}

		h.Register(app.Group("/v1"))
	}

	if err := app.Start("127.0.0.1:80"); err != nil {
		log.Fatal(err)
	}
}