package auth

import (
	"net/http"
	"strings"
	"github.com/labstack/echo/v4"
)

func JwtAuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := IsAuthorized(authToken, secret)
				if authorized {
					userID, err := ExtractIDFromToken(authToken, secret)
					if err != nil {
						return c.JSON(http.StatusUnauthorized, err.Error())
					}
					c.Set("x-user-id", userID)
					return next(c)
				}
				return c.JSON(http.StatusUnauthorized, err.Error())
			}
			return c.JSON(http.StatusUnauthorized, "not authorized")
		}
	}
}
