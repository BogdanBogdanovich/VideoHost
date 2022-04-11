package database

import _ "fmt"

type User struct {
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	StreamKey string `json:"stream_key"`
}

var Users = []User{
	{UserName: "123", Email: "123@mail.ru", Password: "123", StreamKey: "qwerty"},
	{UserName: "321", Email: "321@mail.ru", Password: "123", StreamKey: "qwerty"},
	{UserName: "111", Email: "111@mail.ru", Password: "123", StreamKey: "qwerty"},
}

func FindUserByEmail(Email string) *User {
	for _, user := range Users {
		if user.Email == Email {
			return &user
		}
	}
	return nil
}

func FindUserByUserName(UserName string) *User {
	for _, user := range Users {
		if user.UserName == UserName {
			return &user
		}
	}
	return nil
}

func SaveUser(SaveUser User) {
	for _, user := range Users {
		if user.UserName == SaveUser.UserName {
			user.Password = SaveUser.Password
			user.StreamKey = SaveUser.StreamKey
			user.Email = SaveUser.Email
			return
		}
	}
	Users = append(Users, SaveUser)
}
