package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`
	Password string
}

func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}

func (u *User) UpdatePassword(newPassword string) {
	u.Password = newPassword
}

func (u *User) GetHashedPassword() string {
	return u.Password
}
