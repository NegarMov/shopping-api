package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name 	string 	`json:"username"`
	ID   	uint 	`json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID 		uint 	`json:"id"`
	jwt.StandardClaims
}