package main

import (
	"encoding/json"
	// "fmt"
	"github.com/matryer/is"
	// "io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserLogin(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	r := httptest.NewRequest("GET", "/login", nil)
	w := httptest.NewRecorder()

	srv.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusGone)
}

func TestUserRegister(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	r := httptest.NewRequest("GET", "/register", nil)
	w := httptest.NewRecorder()

	type p struct {
		Username bool `json:"username"`
	}

	var data p
	srv.ServeHTTP(w, r)

	resp := w.Result()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	json.NewDecoder(resp.Body).Decode(data)
	is.Equal(data.Username, "asdf")
}
