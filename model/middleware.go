package model

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	ID        int
	Username  string
	NickName  string
	HeaderImg string
	jwt.StandardClaims
}
