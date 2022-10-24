package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/amarmaulana95/api_seg-v2/api/auth"
	"github.com/amarmaulana95/api_seg-v2/api/models"
	"github.com/amarmaulana95/api_seg-v2/api/responses"
	"github.com/amarmaulana95/api_seg-v2/api/utils/formaterror"
)

func (server *Server) GetSegAnalisaTypes(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	setype := models.SegAnalisaType{}

	dsetype, err := setype.FindAllSegAnalisaTypes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// responses.JSON(w, http.StatusOK, dsetype)
	respon := &ResponStatusData{200, "Success", dsetype}
	responses.JSON(w, http.StatusOK, respon)
}

func (server *Server) GetSegAnalisaType(w http.ResponseWriter, r *http.Request) {

	ff := mux.Vars(r)
	uid, err := strconv.ParseUint(ff["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	setype := models.SegAnalisaType{}
	setypeGotten, err := setype.FindSegAnalisaTypeByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	respon := &ResponStatusData{200, "Success", setypeGotten}
	responses.JSON(w, http.StatusOK, respon)
	// responses.JSON(w, http.StatusOK, setypeGotten)
}

func (server *Server) UpdateSegAnalisaType(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	setype := models.SegAnalisaType{}
	err = json.Unmarshal(body, &setype)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	setype.Prepare()

	updatedSegAnalisaType, err := setype.UpdateASegAnalisaType(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedSegAnalisaType)
}

func (server *Server) DeleteSegAnalisaType(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	setype := models.SegAnalisaType{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = setype.DeleteASegAnalisaType(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) CreateSegAnalisaType(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	setype := models.SegAnalisaType{}
	err = json.Unmarshal(body, &setype)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	setype.Prepare()

	setypeCreated, err := setype.SaveSegAnalisaType(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, setypeCreated.Id))
	responses.JSON(w, http.StatusCreated, setypeCreated)
}
