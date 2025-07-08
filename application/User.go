// package declaration
package application

// imports
import (
	"fit/infrastructure"
	"fit/model"
)

// type definitions
type User struct {
	User string
	Pass string
	ID   int
	Day  int
}

// data
var (
	Users []User
)

// function definitions
func NewUser(user_record *model.UserRecord) User {

	// success
	return User{
		User: user_record.Username,
		Pass: user_record.Password,
		ID:   user_record.ID,
		Day:  user_record.Day,
	}
}

func CreateUsers() {

	// load the user records from redis
	infrastructure.LoadUserRecords()

	// construct the users
	for i := 0; i < len(model.UserRecords); i++ {

		// construct the english menu
		Users = append(Users, NewUser(&model.UserRecords[i]))
	}
}

func UserHandler() any {

	// initialized data
	var User any = Users[0]

	// success
	return User
}
