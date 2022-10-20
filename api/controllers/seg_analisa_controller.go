package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amarmaulana95/api_seg-v2/api/models"
	"github.com/amarmaulana95/api_seg-v2/api/responses"
)

func (server *Server) GetSegAnalisaException(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parsing form", err)
		return
	}
	search := r.FormValue("q")
	p_type := r.FormValue("type")
	if p_type == "" {
		p_type = "0"
	}
	type_analisa, err := strconv.Atoi(p_type)
	if err != nil {
		responses.ERROR(w, http.StatusBadGateway, err)
	}

	semethod := models.SegAnalisaMethod{}

	dsemethod, err := semethod.FindAllSegAnalisaMethodsByException(server.DB, search, uint32(type_analisa))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dsemethod)
}

func (server *Server) GetSegDashboardTotalBulananAnalisa(w http.ResponseWriter, r *http.Request) {

	semethod := models.DashboardAnalisaTotalBulanan{}

	dsemethod, err := semethod.FindAllDashboardAnalisaTotalBulanans(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dsemethod)
}

func (server *Server) GetSegDashboardTotalAnalisa(w http.ResponseWriter, r *http.Request) {

	semethod := models.DashboardAnalisaTotal{}

	dsemethod, err := semethod.FindAllDashboardAnalisaTotals(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dsemethod)
}

