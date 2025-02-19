package controllers

import (
	"github.com/iota-uz/iota-sdk/modules/core/services"
	"github.com/iota-uz/iota-sdk/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iota-uz/iota-sdk/pkg/application"
)

func NewLoginController(app application.Application) application.Controller {
	return &LoginController{
		app:         app,
		authService: app.Service(services.AuthService{}).(*services.AuthService),
	}
}

type LoginController struct {
	app         application.Application
	authService *services.AuthService
}

func (c *LoginController) Key() string {
	return "/login"
}

func (c *LoginController) Register(r *mux.Router) {
	getRouter := r.PathPrefix("/login").Subrouter()
	getRouter.Use(middleware.WithLocalizer(c.app.Bundle()))
	getRouter.Use(middleware.WithTransaction())
	getRouter.HandleFunc("", c.Get).Methods(http.MethodGet)
}

func (c *LoginController) Get(w http.ResponseWriter, r *http.Request) {
	cookie, err := c.authService.CookieAuthenticate(r.Context(), "test@gmail.com", "TestPass123!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
