package controllers

import (
	"fmt"
	"net/http"

	"github.com/anantjakhmola/goweb-backend/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

//This is used to render the form where a user can create a new user account

//GET signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a fake message.pretend we saved the user")
}
