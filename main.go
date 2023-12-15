package main

import (
	"log"
	"github.com/NegarMov/shopping-api/internal/handler"
	"github.com/NegarMov/shopping-api/internal/store/basket"
	"github.com/NegarMov/shopping-api/internal/store/user"
	"github.com/NegarMov/shopping-api/configs"
	"github.com/NegarMov/shopping-api/internal/auth"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

func main() {
	config, err := configs.LoadConfig()
    if err != nil {
        log.Fatal("Error loading configuration file - ", err)
        return
    }

	dbConfig := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		config.DBHost, config.DBPort, config.DBName, config.DBUser, config.DBPass,
	)

	appConfig := fmt.Sprintf("%s:%d", config.ServerIP, config.ServerPort)

	jwtSecret := config.JWTSecret

	app := echo.New()

	db, err := gorm.Open(
		postgres.Open(dbConfig), 
		new(gorm.Config),
	)
	if err != nil {
		log.Fatal(err)
	}

	{
		b := basket.NewSQL(db)
		h := handler.Basket{
			Store: b,
		}

		basketGroup := app.Group("/v1/basket")
		basketGroup.Use(auth.JwtAuthMiddleware(jwtSecret))
		h.Register(basketGroup)
	}

	{
		u := user.NewSQL(db)
		h := handler.User{
			Store: u,
		}

		h.Register(app.Group("/v1/user"))
	}

	if err := app.Start(appConfig); err != nil {
		log.Fatal(err)
	}
}