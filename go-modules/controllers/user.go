package controllers

import (
	"net/http"
	"regexp"
)

// these will be a concept of object oriented simulation on go, so we are adding some beheivior on some construct, in other words a method
// This will be a router that handles the requiest to the correct method

// we initialize the struct with empty fields in order to add functionality
type userController struct {
	// RE defined to create the controler
	userIDPattern *regexp.Regexp
}

func (uc userController) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	// well here what we are doing is to atach a string of bytes to an array, the write method alwas return a byte code
	w.Write([]byte("Hello from the user control"))
}

// we dont have constructors on Go so we have to make a liitle trick like creating a constructor function
func newUserController() *userController {
	return &userController {
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
