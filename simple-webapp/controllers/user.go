package controllers

import (
	"net/http"
	"regexp"
	"simple-webapp/models"
)

type userController struct {
	userIDpattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from userController"))
}

func (uc userController) getALL(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (uc userController) get(id int, w http.ResponseWriter) {
	//TODO
}

func (uc userController) post(w http.ResponseWriter) {
	//TODO
}

func (uc userController) put(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (uc userController) delete(id int, w http.ResponseWriter) {
	//TODO
}

func (uc userController) parseRequest(r *http.Request) (models.User, error) {
	//TODO
}

func newUserController() *userController {
	return &userController{
		userIDpattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
