package controllers

import (
	"net/http"
	"strconv"

	"github.com/amarmaulana95/api_seg-v2/api/models"
	"github.com/amarmaulana95/api_seg-v2/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetLokasis(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	setype := models.Lokasi{}

	dsetype, err := setype.FindAllLokasis(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dsetype)

}

func (server *Server) GetLokasi(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["prop_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	setype := models.Lokasi{}
	setypeGotten, err := setype.FindLokasiByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, setypeGotten)
}
