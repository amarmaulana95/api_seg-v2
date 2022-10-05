package controllers

import (
	"net/http"

	"github.com/amarmaulana95/api_seg-v2/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
