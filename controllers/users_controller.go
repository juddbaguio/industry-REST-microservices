package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/juddbaguio/industry-REST-microservices/services"
	"github.com/juddbaguio/industry-REST-microservices/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetIndent("", "    ")

	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		jsonEncoder.Encode(&utils.ApplicationError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		})
		return
	}

	log.Printf("about to process user_id: %v", userId)
	user, getUserError := services.GetUser(userId)

	if getUserError != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonEncoder.Encode(getUserError)
		return
	}
	w.WriteHeader(http.StatusOK)
	jsonEncoder.Encode(user)
}
