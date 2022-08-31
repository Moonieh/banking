package models

type User struct {
	Username string
	Email    string
	Password string
}

type Account struct {
	Type    string
	Name    string
	Balance int
	UserID  uint
}
