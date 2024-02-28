package web

import (
	"net/http"

	"github.com/go-chi/render"
)

type response struct {
	Message string `json:"message"`
}

func respond(w http.ResponseWriter, r *http.Request, status int, msg string) {
	render.Status(r, status)
	render.JSON(w, r, &response{
		Message: msg,
	})
}
