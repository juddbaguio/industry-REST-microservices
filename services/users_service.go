package services

import (
	"github.com/juddbaguio/industry-REST-microservices/domain"
	"github.com/juddbaguio/industry-REST-microservices/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
