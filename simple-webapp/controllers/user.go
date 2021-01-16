package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDpattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from userController"))
}

func newUserController() *userController {
	return &userController{
		userIDpattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
