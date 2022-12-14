package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/amarmaulana95/api_seg-v2/api/models"
	"github.com/amarmaulana95/api_seg-v2/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) SegAnalisaMethodAll(w http.ResponseWriter, r *http.Request) {
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

	dsemethod, err := semethod.FindAllSegAnalisaMethodsFull(server.DB, uint32(1), search, limit, offset)
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

	dsemethodTotal := semethod.FindAllSegAnalisaMethodsTotal(server.DB, uint32(1), search)

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

func (server *Server) SegAnalisaMethodDelete(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)

	samethodattach := models.SegAnalisaMethodAttachment{}
	samethoddet := models.SegAnalisaMethodDetail{}
	samethodexcp := models.SegAnalisaMethodException{}
	semethod := models.SegAnalisaMethod{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = samethodattach.DeleteSegAnalisaMethodAttachmentByAnalisa(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	_, err = samethoddet.DeleteSegAnalisaMethodDetailByIdAnalisa(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	_, err = samethodexcp.DeleteSegAnalisaMethodExceptionByIdAnalisa(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	_, err = semethod.DeleteSegAnalisaMethod(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) SegAnalisaMethodSingle(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	dId := "0"
	if len(vars["id"]) > 0 {
		dId = vars["id"]
	}

	uid, err := strconv.Atoi(dId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	semethod := models.SegAnalisaMethod{}
	semethodGotten, err := semethod.FindSegAnalisaMethodByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	semethodAttachment := models.SegAnalisaMethodAttachment{}
	semethodAttachmentGotten, err := semethodAttachment.FindAllSegAnalisaMethodAttachmentsByAnalisa(server.DB, uint32(uid))
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
		dAttachment.Path_file_name = "http://10.14.0.18:9002/open_file?filename=" + rAttachment.File_name

		arrAttachment = append(arrAttachment, dAttachment)
	}

	semethodDetail := models.SegAnalisaMethodDetail{}
	semethodDetailGotten, err := semethodDetail.FindAllSegAnalisaMethodDetailByIdAnalisa(server.DB, uint32(uid))
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

		arrDetail = append(arrDetail, dDetail)
	}

	semethodException := models.SegAnalisaMethodException{}
	semethodExceptionGotten, err := semethodException.FindAllSegAnalisaMethodExceptionsByAnalisa(server.DB, uint32(uid))
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

	responAnalisa := ResponAnalisa{semethodGotten.Id, semethodGotten.Id_analisa_type, semethodGotten.Name, semethodGotten.Description, semethodGotten.Location, semethodGotten.Location_name, semethodGotten.Status_proyek_boost, arrAttachment, arrDetail, arrDataException} //

	respon := &ResponStatusData{200, "Berhasil", responAnalisa}
	responses.JSON(w, http.StatusOK, respon)
}

func (server *Server) SegAnalisaMethodDetailDelete(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)

	samethodattach := models.SegAnalisaMethodDetail{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = samethodattach.DeleteSegAnalisaMethodDetail(server.DB, uint32(uid), 1)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) SegAnalisaMethodInsert(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	err := r.ParseMultipartForm(200000)
	if err != nil {
		fmt.Println("error parsing multiplepart form", err)
		return
	}

	r.ParseForm()

	p_status_proyek_boost := r.FormValue("status_proyek_boost")

	if p_status_proyek_boost == "" {
		p_status_proyek_boost = "0"
	}

	project_boost, err := strconv.Atoi(p_status_proyek_boost)
	if err != nil {
		responses.ERROR(w, http.StatusBadGateway, err)
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	status_proyek_boost := uint32(project_boost)

	dRatingElo := models.SegRatingElo{}

	rSegAnalisaMethod := models.SegAnalisaMethod{}
	rSegAnalisaMethod.Id_analisa_type = 1
	rSegAnalisaMethod.Name = name
	rSegAnalisaMethod.Description = description
	rSegAnalisaMethod.Location = r.FormValue("location")
	rSegAnalisaMethod.Status_proyek_boost = status_proyek_boost

	_, err = rSegAnalisaMethod.SaveSegAnalisaMethod(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	id_data := rSegAnalisaMethod.Id

	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	files := formdata.File["analisa_file"] // grab the filenames

	for i, _ := range files { // loop through the files one by one
		rSegAnalisaMethodAttachment := models.SegAnalisaMethodAttachment{}

		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		currenttime := time.Now()
		hashFilename := md5.Sum([]byte(files[i].Filename + currenttime.String()))
		encToString := hex.EncodeToString(hashFilename[:])
		filename_new := fmt.Sprintf("%s%s", encToString, filepath.Ext(files[i].Filename))
		pathfile := getDir() + filename_new

		rSegAnalisaMethodAttachment.Id_analisa = id_data
		rSegAnalisaMethodAttachment.File_name = filename_new
		rSegAnalisaMethodAttachment.Path_file_name = pathfile

		out, err := os.Create(pathfile)

		defer out.Close()
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		_, err = rSegAnalisaMethodAttachment.SaveSegAnalisaMethodAttachment(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	}
	// file_name => file1;file2;file3

	id_barang_data := strings.Split(r.FormValue("id_barang"), ";") // => 1002;983;756
	eficiency_data := strings.Split(r.FormValue("eficiency"), ";") // => 0.76;0,25;0,0043
	analisa_exception_label_data := strings.Split(r.FormValue("analisa_exception_label"), ";")
	eficiency_type_data := strings.Split(r.FormValue("eficiency_type"), ";")
	price_data := strings.Split(r.FormValue("price"), ";")

	for i, _ := range id_barang_data { // loop through the files one by one

		rSegAnalisaMethodDetail := models.SegAnalisaMethodDetail{}

		eficiency, err := strconv.ParseFloat(eficiency_data[i], 32)
		if err != nil {
			responses.ERROR(w, http.StatusBadGateway, err)
		}

		eficiency_type, err := strconv.Atoi(eficiency_type_data[i])
		if err != nil {
			responses.ERROR(w, http.StatusBadGateway, err)
		}

		price, err := strconv.ParseFloat(price_data[i], 32)
		if err != nil {
			responses.ERROR(w, http.StatusBadGateway, err)
		}

		rSegAnalisaMethodDetail.Id_analisa = id_data
		rSegAnalisaMethodDetail.Id_barang = id_barang_data[i]
		rSegAnalisaMethodDetail.Eficiency = float32(eficiency)
		rSegAnalisaMethodDetail.Barang = analisa_exception_label_data[i]
		rSegAnalisaMethodDetail.Eficiency_type = uint32(eficiency_type)
		rSegAnalisaMethodDetail.Price = float32(price)

		_, err = rSegAnalisaMethodDetail.SaveSegAnalisaMethodDetail(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		dIdDetail := rSegAnalisaMethodDetail.Id
		dId_barang := rSegAnalisaMethodDetail.Id_barang
		if dId_barang == "0" {
			dId_barang = fmt.Sprintf("00-%d", dIdDetail)
		}

		dRatingElo = models.SegRatingElo{}

		dRatingElo.Id_analisa_type = 1
		dRatingElo.Id_analisa = id_data
		dRatingElo.Id_barang = dId_barang
		dRatingElo.Koefisien = rSegAnalisaMethodDetail.Eficiency
		dRatingElo.Rating = float32(rSegAnalisaMethodDetail.Eficiency) * float32(1000)

		_, err = dRatingElo.SaveSegRatingElo(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	}

	p_analisa_type := r.FormValue("analisa_type")

	if p_analisa_type == "" {
		p_analisa_type = "0"
	}

	p_id_analisa_exception := r.FormValue("id_analisa_exception")

	if p_id_analisa_exception == "" {
		p_id_analisa_exception = "0"
	}

	analisa_type_data := strings.Split(p_analisa_type, ";")                 //  => 1;1;2;3;2
	id_analisa_exception_data := strings.Split(p_id_analisa_exception, ";") //  => 10;35;20;21;5
	label_type_data := strings.Split(r.FormValue("label_type"), ";")
	label_exception_data := strings.Split(r.FormValue("label_exception"), ";")

	for i, _ := range id_analisa_exception_data { // loop through the files one by one

		rSegAnalisaMethodException := models.SegAnalisaMethodException{}

		analisa_type, err := strconv.Atoi(analisa_type_data[i])
		if err != nil {
			responses.ERROR(w, http.StatusBadGateway, err)
		}

		id_analisa_exception, err := strconv.Atoi(id_analisa_exception_data[i])
		if err != nil {
			responses.ERROR(w, http.StatusBadGateway, err)
		}

		rSegAnalisaMethodException.Id_analisa = id_data
		rSegAnalisaMethodException.Analisa_type = uint32(analisa_type)
		rSegAnalisaMethodException.Id_analisa_exception = uint32(id_analisa_exception)
		rSegAnalisaMethodException.Label_type = label_type_data[i]
		rSegAnalisaMethodException.Label_exception = label_exception_data[i]

		_, err = rSegAnalisaMethodException.SaveSegAnalisaMethodException(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	}

	respon := &ResponStatusData{200, "Berhasil", rSegAnalisaMethod}
	responses.JSON(w, http.StatusOK, respon)
}

func (server *Server) SegAnalisaMethodExceptionDelete(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)

	samethodattach := models.SegAnalisaMethodException{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = samethodattach.DeleteSegAnalisaMethodException(server.DB, uint32(uid), 1)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) GetSegDashboardEficiency(w http.ResponseWriter, r *http.Request) {

	semethod := models.DashboardAnalisaEficiency{}

	dsemethod, err := semethod.FindDashboardAnalisaEficiencyByID(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	dEfi := DashEfi{}
	dEfi.Data_a[0] = dsemethod.Percent_eficiency
	dEfi.Data_a[1] = 100 - dsemethod.Percent_eficiency
	dEfi.Persen = fmt.Sprintf("%.2f%%", dsemethod.Percent_eficiency)

	respon := &ResponStatusData{200, "Success", dEfi}
	responses.JSON(w, http.StatusOK, respon)

	// responses.JSON(w, http.StatusOK, dEfi)

}
