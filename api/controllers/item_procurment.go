package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amarmaulana95/api_seg-v2/api/models"
	"github.com/amarmaulana95/api_seg-v2/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetItemProcurments(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parsing form", err)
		return
	}

	search := r.FormValue("q")
	id_provinsi := r.FormValue("id_provinsi")

	setype := models.ItemProcurment{}

	dsetype, err := setype.FindAllItemProcurments(server.DB, search, id_provinsi)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dsetype)
}

func (server *Server) GetItemProcurment(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	setype := models.ItemProcurment{}
	setypeGotten, err := setype.FindItemProcurmentByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, setypeGotten)
}
