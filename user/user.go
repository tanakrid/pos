package user

import "errors"
import "log"

type User struct {
	Username string
	Password string
}

type Customer struct {
	User User
	Wallet int
}

type Employee struct {
	User User
}

func (u User) Authen(username ,pass string, user User) (bool) {
	if username == "" || pass == "" {
		log.Fatalln(errors.New("don't enter password"))
	}
	return username == user.Username && pass == user.Password
}