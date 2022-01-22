package domain

import (
	"fmt"
	"net/http"

	"github.com/juddbaguio/industry-REST-microservices/utils"
)

var (
	users = map[int64]*User{
		1: {
			Id:        1,
			FirstName: "Judd Misael",
			LastName:  "Baguio",
			Email:     "juddmisaelbaguio@gmail.com",
		},
		2: {
			Id:        2,
			FirstName: "Keith Yvonne",
			LastName:  "Saycon",
			Email:     "kysaycon@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user id %v not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not found",
		}
	}

	return user, nil
}
