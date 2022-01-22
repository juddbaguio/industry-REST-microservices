package app

import (
	"net/http"

	"github.com/juddbaguio/industry-REST-microservices/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
