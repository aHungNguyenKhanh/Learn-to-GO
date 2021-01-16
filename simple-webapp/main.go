package main

import (
	"net/http"
	"simple-webapp/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
