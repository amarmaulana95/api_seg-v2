package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amarmaulana95/api_seg-v2/api/models"
	"github.com/amarmaulana95/api_seg-v2/api/responses"
)

func (server *Server) SegAnalisaInovasiAll(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parsing form", err)
		return
	}

	responData := ResponStatusDataView{}

	search := r.FormValue("q")

	page_data := "1"
	if len(r.FormValue("_page")) > 0 {
		page_data = r.FormValue("_page")
	}

	page, err := strconv.ParseUint(page_data, 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	limit_data := "100"
	if len(r.FormValue("_limit")) > 0 {
		limit_data = r.FormValue("_limit")
	}

	limit, err := strconv.ParseUint(limit_data, 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	offset := ((page - 1) * limit)

	responAnalisa := ResponAnalisa{}
	arrResponAnalisa := []ResponAnalisa{}

	semethod := models.SegAnalisaMethod{}

	dsemethod, err := semethod.FindAllSegAnalisaMethodsFull(server.DB, uint32(2), search, limit, offset)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	dAnalisa := models.SegAnalisaMethod{}
	arrAnalisa := []models.SegAnalisaMethod{}
	for i, _ := range dsemethod { // loop through the files one by one
		dAnalisa.Id = dsemethod[i].Id
		dAnalisa.Id_analisa_type = dsemethod[i].Id_analisa_type
		dAnalisa.Name = dsemethod[i].Name
		dAnalisa.Description = dsemethod[i].Description
		dAnalisa.Location = dsemethod[i].Location
		dAnalisa.Location_name = dsemethod[i].Location_name
		dAnalisa.Status_proyek_boost = dsemethod[i].Status_proyek_boost

		arrAnalisa = append(arrAnalisa, dAnalisa)

		// ----------------------------------------------

		semethodAttachment := models.SegAnalisaMethodAttachment{}
		semethodAttachmentGotten, err := semethodAttachment.FindAllSegAnalisaMethodAttachmentsByAnalisa(server.DB, uint32(dsemethod[i].Id))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		dAttachment := DataAttachment{}
		arrAttachment := []DataAttachment{}

		for _, rAttachment := range semethodAttachmentGotten {
			dAttachment.Id = rAttachment.Id
			dAttachment.File_name = "file_data"
			dAttachment.Attachment = rAttachment.File_name
			dAttachment.Path_file_name = "http://localhost/open_file?filename=" + rAttachment.File_name

			arrAttachment = append(arrAttachment, dAttachment)
		}

		semethodDetail := models.SegAnalisaMethodDetail{}
		semethodDetailGotten, err := semethodDetail.FindAllSegAnalisaMethodDetailByIdAnalisa(server.DB, uint32(dsemethod[i].Id))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		dDetail := DataDetail{}
		arrDetail := []DataDetail{}
		for i, _ := range semethodDetailGotten { // loop through the files one by one
			dDetail.Id = semethodDetailGotten[i].Id
			dDetail.Id_barang = semethodDetailGotten[i].Id_barang
			dDetail.Label_barang = semethodDetailGotten[i].Barang
			dDetail.Eficiency = semethodDetailGotten[i].Eficiency
			dDetail.Eficiency_type = float32(semethodDetailGotten[i].Eficiency_type)
			dDetail.Price = float32(semethodDetailGotten[i].Price)

			arrDetail = append(arrDetail, dDetail)
		}

		semethodException := models.SegAnalisaMethodException{}
		semethodExceptionGotten, err := semethodException.FindAllSegAnalisaMethodExceptionsByAnalisa(server.DB, uint32(dsemethod[i].Id))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		dDataException := DataException{}
		arrDataException := []DataException{}
		for i, _ := range semethodExceptionGotten { // loop through the files one by one
			dDataException.Id = semethodExceptionGotten[i].Id
			dDataException.Analisa_type = semethodExceptionGotten[i].Analisa_type
			dDataException.Label_type = semethodExceptionGotten[i].Label_type
			dDataException.Id_analisa_exception = semethodExceptionGotten[i].Id_analisa_exception
			dDataException.Label_exception = semethodExceptionGotten[i].Label_exception

			arrDataException = append(arrDataException, dDataException)
		}

		// ----------------------------------------------

		responAnalisa = ResponAnalisa{dAnalisa.Id, dAnalisa.Id_analisa_type, dAnalisa.Name, dAnalisa.Description, dAnalisa.Location, dAnalisa.Location_name, dAnalisa.Status_proyek_boost, arrAttachment, arrDetail, arrDataException} //
		arrResponAnalisa = append(arrResponAnalisa, responAnalisa)
	}

	dsemethodTotal := semethod.FindAllSegAnalisaMethodsTotal(server.DB, uint32(2), search)

	selisih := dsemethodTotal % limit

	total_pages := 1

	if selisih == 0 {
		total_pages = (int(dsemethodTotal) / int(limit))
	} else {
		total_pages = (int(dsemethodTotal) / int(limit)) + 1
	}

	responData = ResponStatusDataView{uint32(page), uint32(limit), uint32(total_pages), uint32(dsemethodTotal), arrResponAnalisa}

	responses.JSON(w, http.StatusOK, responData)
}
