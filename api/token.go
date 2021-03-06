package main

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token      string `json:"token"`
	Exp        int64  `json:"exp"`
	UserId     int    `json:"userId"`
	Authorized bool   `json:"authorized"`
	Admin      bool   `json:"admin"`
}

func newToken(userId int, admin bool) (*Token, error) {
	expires_in := time.Now().Add(time.Minute * 15).Unix()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["admin"] = admin
	atClaims["user_id"] = userId
	atClaims["exp"] = expires_in
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	return &Token{
		Token:      token,
		Exp:        expires_in,
		UserId:     userId,
		Admin:      admin,
		Authorized: true,
	}, nil
}
