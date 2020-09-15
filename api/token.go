package main

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Token struct {
	Token      string `json:"token"`
	Exp        int64  `json:"exp"`
	UserId     int64  `json:"userId"`
	Authorized bool   `json:"authorized"`
	Admin      bool   `json:"admin"`
}

func newToken(userId int64, admin bool) (*Token, error) {
	os.Setenv("ACCESS_SECRET", "asdf")
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
