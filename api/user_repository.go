package main

import (
	models "main/models"
)

type BaseUserRepository interface {
	GetUserById(id int) *models.User
	GetUserByUsername(username string) *models.User
	InsertUser(user models.User) bool
	GetAllUsers() []models.User
}

type LiveUserRepository struct{}

func (u LiveUserRepository) GetUserById(id int) *models.User {
	return &models.User{}
}

func (u LiveUserRepository) GetUserByUsername(username string) *models.User {
	return &models.User{}
}

func (u LiveUserRepository) InsertUser(user models.User) bool {
	return true
}

func (u LiveUserRepository) GetAllUsers() []models.User {
	return []models.User{}
}

type MockUserRepository struct {
	users []models.User
}

func (u *MockUserRepository) InsertUser(user models.User) bool {
	u.users = append(u.users, user)
	return true
}

func (u *MockUserRepository) GetUserByUsername(username string) *models.User {
	for _, user := range u.users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

func (u *MockUserRepository) GetAllUsers() []models.User {
	return u.users
}

func (u *MockUserRepository) GetUserById(id int) *models.User {
	for _, element := range u.users {
		if element.ID == id {
			return &element
		}
	}
	return nil
}
