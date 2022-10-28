package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	"github.com/amarmaulana95/api_seg-v2/api/models"
	"github.com/amarmaulana95/api_seg-v2/api/responses"
	"github.com/gorilla/mux"
)

type RatingTenderItem struct {
	Select_item_a float32 `json:"select_item_a"`
	Select_item_b float32 `json:"select_item_b"`
}

func (server *Server) SegTenderEstimateCalculation(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println(err)
	}

	rtesc := RawTenderEstimateCalculation{}
	err = json.Unmarshal(body, &rtesc)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tenderEstimate := models.SegTenderEstimate{}

	tenderEstimate.Project_name = rtesc.Project_name
	tenderEstimate.Location = rtesc.Province
	tenderEstimate.Construction_type = rtesc.Construction_type
	tenderEstimate.Building_designation = rtesc.Building
	tenderEstimate.Class = rtesc.Tipe_class

	_, err = tenderEstimate.SaveSegTenderEstimate(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	id_data := tenderEstimate.Id

	for i, _ := range rtesc.Component_item { // loop through the files one by one
		tenderEstimateDetail := models.SegTenderEstimateDetail{}

		tenderEstimateDetail.Id_tender_estimate = id_data
		tenderEstimateDetail.Id_barang = rtesc.Component_item[i].Item
		tenderEstimateDetail.Price = rtesc.Component_item[i].Price
		tenderEstimateDetail.Method = rtesc.Component_item[i].Method
		tenderEstimateDetail.Innovation = rtesc.Component_item[i].Innovation
		tenderEstimateDetail.Value_enginering = rtesc.Component_item[i].Value_eng
		tenderEstimateDetail.Finance_enginering = rtesc.Component_item[i].Finance_eng

		_, err = tenderEstimateDetail.SaveSegTenderEstimateDetail(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	}

	for i, _ := range rtesc.Project_boost { // loop through the files one by one
		tenderEstimateProyekBoost := models.SegTenderEstimateProyekBoost{}

		tenderEstimateProyekBoost.Id_tender_estimate = id_data
		tenderEstimateProyekBoost.Analisa_type = rtesc.Project_boost[i].Select_object
		tenderEstimateProyekBoost.Id_analisa = rtesc.Project_boost[i].Name_object

		_, err = tenderEstimateProyekBoost.SaveSegTenderEstimateProyekBoost(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	}

	responses.JSON(w, http.StatusOK, id_data)
}

func (server *Server) SegTenderEstimateCalculationRating(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	uid, err_mux := strconv.ParseUint(vars["id"], 10, 32)
	if err_mux != nil {
		responses.ERROR(w, http.StatusBadRequest, err_mux)
		return
	}

	semethodEstimateHead := models.SegTenderEstimate{}
	semethodEstimateHeadGotten, err := semethodEstimateHead.FindSegTenderEstimateByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	dRatingTenderEstimate := RatingTenderEstimate{}
	dRatingTenderEstimate.Id = semethodEstimateHeadGotten.Id
	dRatingTenderEstimate.Project = semethodEstimateHeadGotten.Project_name
	dRatingTenderEstimate.Location = semethodEstimateHeadGotten.Location
	dRatingTenderEstimate.Const_type = semethodEstimateHeadGotten.Construction_type
	dRatingTenderEstimate.BuiLding = semethodEstimateHeadGotten.Building_designation
	dRatingTenderEstimate.Class_project = semethodEstimateHeadGotten.Class

	semethodEstimate := models.SegTenderEstimateDetail{}
	semethodEstimateGotten, err := semethodEstimate.FindAllSegTenderEstimateDetailsByIdTender(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	semethodEstimateEloMethod := models.SegRatingElo{}
	semethodEstimateEloMethodGotten := &models.SegRatingElo{}

	semethodEstimateEloInovation := models.SegRatingElo{}
	semethodEstimateEloInovationGotten := &models.SegRatingElo{}

	semethodEstimateEloVe := models.SegRatingElo{}
	semethodEstimateEloVeGotten := &models.SegRatingElo{}

	semethodEstimateEloFe := models.SegRatingElo{}
	semethodEstimateEloFeGotten := &models.SegRatingElo{}

	dRatingTenderEstimateDetail := RatingTenderEstimateDetail{}
	arrRatingTenderEstimateDetail := []RatingTenderEstimateDetail{}

	dRatingTenderEstimateMethodItem := RatingTenderEstimateMethodItem{}
	arrRatingTenderEstimateMethodItem := []RatingTenderEstimateMethodItem{}

	dRatingTenderEstimateMethod := RatingTenderEstimateMethod{}

	/////////////////////////////////////////////////////////////////////////////////////////

	dRatingTenderEstimateInovationItem := RatingTenderEstimateInovationItem{}
	arrRatingTenderEstimateInovationItem := []RatingTenderEstimateInovationItem{}

	dRatingTenderEstimateInovation := RatingTenderEstimateInovation{}

	/////////////////////////////////////////////////////////////////////////////////////////

	dRatingTenderEstimateVeItem := RatingTenderEstimateVeItem{}
	arrRatingTenderEstimateVeItem := []RatingTenderEstimateVeItem{}

	dRatingTenderEstimateVe := RatingTenderEstimateVe{}

	////////////////////////////////////////////////////////////////////////////////////////

	dRatingTenderEstimateFeItem := RatingTenderEstimateFeItem{}
	arrRatingTenderEstimateFeItem := []RatingTenderEstimateFeItem{}

	dRatingTenderEstimateFe := RatingTenderEstimateFe{}

	dSegAnalisaMethodDetail := models.SegAnalisaMethodDetail{}

	for i, _ := range semethodEstimateGotten { // loop through the files one by one
		dRatingTenderEstimateDetail.Id_barang = semethodEstimateGotten[i].Id_barang
		dRatingTenderEstimateDetail.Item_name = semethodEstimateGotten[i].Nama_barang
		dRatingTenderEstimateDetail.Price = semethodEstimateGotten[i].Price

		dRatingTenderEstimateMethod.Select_method = 0
		dRatingTenderEstimateInovation.Select_innovation = 0
		dRatingTenderEstimateVe.Select_value_e = 0
		dRatingTenderEstimateFe.Select_finance = 0

		semethodEstimate := models.SegTenderEstimateDetail{}
		semethodEstimateGotten, err := semethodEstimate.FindAllSegTenderEstimateDetailsByIdTenderIdBarang(server.DB, uint32(uid), semethodEstimateGotten[i].Id_barang)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		dRatingTenderEstimateMethodItem = RatingTenderEstimateMethodItem{}
		arrRatingTenderEstimateMethodItem = []RatingTenderEstimateMethodItem{}

		dRatingTenderEstimateInovationItem = RatingTenderEstimateInovationItem{}
		arrRatingTenderEstimateInovationItem = []RatingTenderEstimateInovationItem{}

		dRatingTenderEstimateVeItem = RatingTenderEstimateVeItem{}
		arrRatingTenderEstimateVeItem = []RatingTenderEstimateVeItem{}

		dRatingTenderEstimateFeItem = RatingTenderEstimateFeItem{}
		arrRatingTenderEstimateFeItem = []RatingTenderEstimateFeItem{}

		for i, _ := range semethodEstimateGotten { // loop through the files one by one
			dRatingTenderEstimateMethodItem.Id_analisa_method = semethodEstimateGotten[i].Method
			dRatingTenderEstimateMethodItem.Item_method = semethodEstimateGotten[i].Method_name
			dRatingTenderEstimateMethodItem.Item_value = semethodEstimateGotten[i].Method_koefisien
			dRatingTenderEstimateMethodItem.Rekomendasi = 0
			dRatingTenderEstimateMethodItem.T_efficiency_result = semethodEstimateGotten[i].Method_koefisien * dRatingTenderEstimateDetail.Price

			arrRatingTenderEstimateMethodItem = append(arrRatingTenderEstimateMethodItem, dRatingTenderEstimateMethodItem)

			semethodEstimateEloMethod = models.SegRatingElo{}
			semethodEstimateEloMethodGotten, err = semethodEstimateEloMethod.FindAllSegRatingEloRecommend(server.DB, 1, semethodEstimateGotten[i].Id_barang)

			if semethodEstimateEloMethodGotten.Id != 0 {
				if semethodEstimateGotten[i].Method != semethodEstimateEloMethodGotten.Id_analisa {
					arrRatingTenderEstimateMethodItem[0].Rekomendasi = 0

					dRatingTenderEstimateMethodItem.Id_analisa_method = semethodEstimateEloMethodGotten.Id_analisa
					dRatingTenderEstimateMethodItem.Item_method = semethodEstimateEloMethodGotten.Analisa_name
					dRatingTenderEstimateMethodItem.Item_value = semethodEstimateEloMethodGotten.Koefisien
					dRatingTenderEstimateMethodItem.Rekomendasi = 1
					dRatingTenderEstimateMethodItem.T_efficiency_result = semethodEstimateEloMethodGotten.Koefisien * dRatingTenderEstimateDetail.Price

					arrRatingTenderEstimateMethodItem = append(arrRatingTenderEstimateMethodItem, dRatingTenderEstimateMethodItem)
				} else {
					arrRatingTenderEstimateMethodItem[0].Rekomendasi = 1
				}
			} else {
				arrRatingTenderEstimateMethodItem[0].Rekomendasi = 1
			}

			///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

			dRatingTenderEstimateInovationItem.Id_analisa_inovasi = semethodEstimateGotten[i].Innovation
			dRatingTenderEstimateInovationItem.Innovation_item = semethodEstimateGotten[i].Innovation_name
			dRatingTenderEstimateInovationItem.Innovation_value = semethodEstimateGotten[i].Innovation_koefisien
			dRatingTenderEstimateInovationItem.Rekomendasi = 0
			dRatingTenderEstimateMethodItem.T_efficiency_result = semethodEstimateGotten[i].Innovation_koefisien * dRatingTenderEstimateDetail.Price

			arrRatingTenderEstimateInovationItem = append(arrRatingTenderEstimateInovationItem, dRatingTenderEstimateInovationItem)

			semethodEstimateEloInovation = models.SegRatingElo{}
			semethodEstimateEloInovationGotten, err = semethodEstimateEloInovation.FindAllSegRatingEloRecommend(server.DB, 2, semethodEstimateGotten[i].Id_barang)

			if semethodEstimateEloInovationGotten.Id != 0 {
				if semethodEstimateGotten[i].Innovation != semethodEstimateEloInovationGotten.Id_analisa {
					arrRatingTenderEstimateInovationItem[0].Rekomendasi = 0

					dRatingTenderEstimateInovationItem.Id_analisa_inovasi = semethodEstimateEloInovationGotten.Id_analisa
					dRatingTenderEstimateInovationItem.Innovation_item = semethodEstimateEloInovationGotten.Analisa_name
					dRatingTenderEstimateInovationItem.Innovation_value = semethodEstimateEloInovationGotten.Koefisien
					dRatingTenderEstimateInovationItem.Rekomendasi = 1
					dRatingTenderEstimateMethodItem.T_efficiency_result = semethodEstimateEloInovationGotten.Koefisien * dRatingTenderEstimateDetail.Price

					arrRatingTenderEstimateInovationItem = append(arrRatingTenderEstimateInovationItem, dRatingTenderEstimateInovationItem)
				} else {
					arrRatingTenderEstimateInovationItem[0].Rekomendasi = 1
				}
			} else {
				arrRatingTenderEstimateInovationItem[0].Rekomendasi = 1
			}

			///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

			dRatingTenderEstimateVeItem.Id_analisa_ve = semethodEstimateGotten[i].Value_enginering
			dRatingTenderEstimateVeItem.Valuee_item = semethodEstimateGotten[i].Value_enginering_name
			dRatingTenderEstimateVeItem.Valuee_value = semethodEstimateGotten[i].Value_enginering_koefisien
			dRatingTenderEstimateVeItem.Rekomendasi = 0
			dRatingTenderEstimateMethodItem.T_efficiency_result = semethodEstimateGotten[i].Value_enginering_koefisien * dRatingTenderEstimateDetail.Price

			arrRatingTenderEstimateVeItem = append(arrRatingTenderEstimateVeItem, dRatingTenderEstimateVeItem)

			semethodEstimateEloVe = models.SegRatingElo{}
			semethodEstimateEloVeGotten, err = semethodEstimateEloVe.FindAllSegRatingEloRecommend(server.DB, 3, semethodEstimateGotten[i].Id_barang)

			if semethodEstimateEloVeGotten.Id != 0 {
				if semethodEstimateGotten[i].Value_enginering != semethodEstimateEloVeGotten.Id_analisa {
					arrRatingTenderEstimateVeItem[0].Rekomendasi = 0

					dRatingTenderEstimateVeItem.Id_analisa_ve = semethodEstimateEloVeGotten.Id_analisa
					dRatingTenderEstimateVeItem.Valuee_item = semethodEstimateEloVeGotten.Analisa_name
					dRatingTenderEstimateVeItem.Valuee_value = semethodEstimateEloVeGotten.Koefisien
					dRatingTenderEstimateVeItem.Rekomendasi = 1
					dRatingTenderEstimateMethodItem.T_efficiency_result = semethodEstimateEloVeGotten.Koefisien * dRatingTenderEstimateDetail.Price

					arrRatingTenderEstimateVeItem = append(arrRatingTenderEstimateVeItem, dRatingTenderEstimateVeItem)
				} else {
					arrRatingTenderEstimateVeItem[0].Rekomendasi = 1
				}
			} else {
				arrRatingTenderEstimateVeItem[0].Rekomendasi = 1
			}

			///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

			dRatingTenderEstimateFeItem.Id_analisa_fe = semethodEstimateGotten[i].Finance_enginering
			dRatingTenderEstimateFeItem.Finance_item = semethodEstimateGotten[i].Finance_enginering_name
			dRatingTenderEstimateFeItem.Finance_value = semethodEstimateGotten[i].Finance_enginering_koefisien
			dRatingTenderEstimateFeItem.Rekomendasi = 0
			dRatingTenderEstimateMethodItem.T_efficiency_result = semethodEstimateGotten[i].Finance_enginering_koefisien * dRatingTenderEstimateDetail.Price

			arrRatingTenderEstimateFeItem = append(arrRatingTenderEstimateFeItem, dRatingTenderEstimateFeItem)

			semethodEstimateEloFe = models.SegRatingElo{}
			semethodEstimateEloFeGotten, err = semethodEstimateEloFe.FindAllSegRatingEloRecommend(server.DB, 4, semethodEstimateGotten[i].Id_barang)

			if semethodEstimateEloFeGotten.Id != 0 {
				if semethodEstimateGotten[i].Finance_enginering != semethodEstimateEloFeGotten.Id_analisa {
					arrRatingTenderEstimateFeItem[0].Rekomendasi = 0

					dRatingTenderEstimateFeItem.Id_analisa_fe = semethodEstimateEloFeGotten.Id_analisa
					dRatingTenderEstimateFeItem.Finance_item = semethodEstimateEloFeGotten.Analisa_name
					dRatingTenderEstimateFeItem.Finance_value = semethodEstimateEloFeGotten.Koefisien
					dRatingTenderEstimateFeItem.Rekomendasi = 1
					dRatingTenderEstimateMethodItem.T_efficiency_result = semethodEstimateEloFeGotten.Koefisien * dRatingTenderEstimateDetail.Price

					arrRatingTenderEstimateFeItem = append(arrRatingTenderEstimateFeItem, dRatingTenderEstimateFeItem)
				} else {
					arrRatingTenderEstimateFeItem[0].Rekomendasi = 1
				}
			} else {
				arrRatingTenderEstimateFeItem[0].Rekomendasi = 1
			}

			///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		}

		dRatingTenderEstimateMethod.Item = arrRatingTenderEstimateMethodItem
		dRatingTenderEstimateInovation.Item = arrRatingTenderEstimateInovationItem
		dRatingTenderEstimateVe.Item = arrRatingTenderEstimateVeItem
		dRatingTenderEstimateFe.Item = arrRatingTenderEstimateFeItem

		dRatingTenderEstimateDetail.Method = dRatingTenderEstimateMethod
		dRatingTenderEstimateDetail.Innovation = dRatingTenderEstimateInovation
		dRatingTenderEstimateDetail.Value_e = dRatingTenderEstimateVe
		dRatingTenderEstimateDetail.Finance_e = dRatingTenderEstimateFe

		arrRatingTenderEstimateDetail = append(arrRatingTenderEstimateDetail, dRatingTenderEstimateDetail)

		semethodEstimateBoost := models.SegTenderEstimateProyekBoost{}
		semethodEstimateBoostGotten, err := semethodEstimateBoost.FindAllSegTenderEstimateProyekBoostsByIdTender(server.DB, uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		dRatingTenderEstimateProjectBoost := RatingTenderEstimateProjectBoost{}
		arrRatingTenderEstimateProjectBoost := []RatingTenderEstimateProjectBoost{}

		for i, _ := range semethodEstimateBoostGotten { // loop through the files one by one
			dRatingTenderEstimateProjectBoost.Select_object = semethodEstimateBoostGotten[i].Analisa_type
			dRatingTenderEstimateProjectBoost.Value_object = semethodEstimateBoostGotten[i].Id_analisa
			dRatingTenderEstimateProjectBoost.Select_name = semethodEstimateBoostGotten[i].Select_name
			dRatingTenderEstimateProjectBoost.Value_name = semethodEstimateBoostGotten[i].Value_name

			semethodEstimate := models.SegTenderEstimateDetail{}
			semethodEstimateGotten, err := semethodEstimate.FindAllSegTenderEstimateDetailsByIdTender(server.DB, uint32(uid))
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			t_efficiency_project_boost := float32(0)
			t_efficiency_project_boost_result := float32(0)

			for i, _ := range semethodEstimateGotten { // loop through the files one by one
				id_barang := semethodEstimateGotten[i].Id_barang

				dSegAnalisaMethodDetail = models.SegAnalisaMethodDetail{}
				dSegAnalisaMethodDetailGotten, err := dSegAnalisaMethodDetail.FindSegAnalisaMethodDetailByAnalisaIdbarang(server.DB, uint32(dRatingTenderEstimateProjectBoost.Value_object), id_barang)
				if err != nil {

				}

				t_efficiency_project_boost = t_efficiency_project_boost + dSegAnalisaMethodDetailGotten.Eficiency
				t_efficiency_project_boost_result = t_efficiency_project_boost * dSegAnalisaMethodDetailGotten.Price

			}

			dRatingTenderEstimateProjectBoost.T_efficiency_project_boost = t_efficiency_project_boost
			dRatingTenderEstimateProjectBoost.T_efficiency_project_boost_result = t_efficiency_project_boost_result

			arrRatingTenderEstimateProjectBoost = append(arrRatingTenderEstimateProjectBoost, dRatingTenderEstimateProjectBoost)
		}

		dRatingTenderEstimate.Item_tabel = arrRatingTenderEstimateDetail
		dRatingTenderEstimate.Project_boost = arrRatingTenderEstimateProjectBoost
	}

	responses.JSON(w, http.StatusOK, dRatingTenderEstimate)
}

func (server *Server) SegTenderEstimateCalibrateRatingFirst(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parsing form", err)
		return
	}

	/*calVal, err := strconv.ParseFloat(r.FormValue("calibrate_value"), 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}	*/

	calVal := 1000

	dRatingElo := models.SegRatingElo{}
	arrRatingElo := []models.SegRatingElo{}

	semethodRatElo := models.SegRatingElo{}
	semethodRatEloGotten, err := semethodRatElo.FindAllSegRatingElosCalibrate(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	for i, _ := range semethodRatEloGotten { // loop through the files one by one
		dRatingElo = models.SegRatingElo{}

		dRatingElo.Id_analisa_type = semethodRatEloGotten[i].Id_analisa_type
		dRatingElo.Id_analisa = semethodRatEloGotten[i].Id_analisa
		dRatingElo.Id_barang = semethodRatEloGotten[i].Id_barang
		dRatingElo.Koefisien = semethodRatEloGotten[i].Koefisien
		dRatingElo.Rating = float32(semethodRatEloGotten[i].Koefisien) * float32(calVal)

		_, err = dRatingElo.SaveSegRatingElo(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		arrRatingElo = append(arrRatingElo, dRatingElo)
	}

	responses.JSON(w, http.StatusOK, arrRatingElo)
}

func (server *Server) SegTenderEstimateResultSave(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println(err)
	}

	rtes := RatingTenderEstimate{}
	err = json.Unmarshal(body, &rtes)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	rtesresult := RatingTenderEstimateResult{}
	rtesresult.Id = rtes.Id
	rtesresult.Project = rtes.Project
	rtesresult.Location = rtes.Location
	rtesresult.Const_type = rtes.Const_type
	rtesresult.BuiLding = rtes.BuiLding
	rtesresult.Class_project = rtes.Class_project

	rtesdetailresult := RatingTenderEstimateDetailResult{}
	arrtesdetailresult := []RatingTenderEstimateDetailResult{}

	rtesmethodresult := RatingTenderEstimateMethodResult{}
	rtesinovationresult := RatingTenderEstimateInovationResult{}
	rtesveresult := RatingTenderEstimateVeResult{}
	rtesferesult := RatingTenderEstimateFeResult{}

	for i, _ := range rtes.Item_tabel { // loop through the files one by one
		tenderEstimateDetail := models.SegTenderEstimateDetailResult{}

		tenderEstimateDetail.Id_tender_estimate = rtes.Id
		tenderEstimateDetail.Id_barang = rtes.Item_tabel[i].Id_barang
		tenderEstimateDetail.Price = rtes.Item_tabel[i].Price
		tenderEstimateDetail.Method = rtes.Item_tabel[i].Method.Select_method
		tenderEstimateDetail.Innovation = rtes.Item_tabel[i].Innovation.Select_innovation
		tenderEstimateDetail.Value_enginering = rtes.Item_tabel[i].Value_e.Select_value_e
		tenderEstimateDetail.Finance_enginering = rtes.Item_tabel[i].Finance_e.Select_finance

		rtesdetailresult = RatingTenderEstimateDetailResult{}
		rtesdetailresult.Id_barang = rtes.Item_tabel[i].Id_barang
		rtesdetailresult.Item_name = rtes.Item_tabel[i].Item_name
		rtesdetailresult.Price = rtes.Item_tabel[i].Price
		/*Method
		Innovation
		Value_e
		Finance_e*/

		_, err = tenderEstimateDetail.SaveSegTenderEstimateDetailResult(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		/////////////////////////////////////////// METHOD /////////////////////////////////////////////////////////////////

		id_barang := rtes.Item_tabel[i].Id_barang
		select_item := rtes.Item_tabel[i].Method.Select_method
		select_item_a_id := rtes.Item_tabel[i].Method.Item[0].Id_analisa_method

		if len(rtes.Item_tabel[i].Method.Item) == 1 {
			dRatingTenderEstimateMethodItem := RatingTenderEstimateMethodItem{}
			dRatingTenderEstimateMethodItem.Id_analisa_method = rtes.Item_tabel[i].Method.Item[0].Id_analisa_method
			dRatingTenderEstimateMethodItem.Item_method = rtes.Item_tabel[i].Method.Item[0].Item_method
			dRatingTenderEstimateMethodItem.Item_value = rtes.Item_tabel[i].Method.Item[0].Item_value
			dRatingTenderEstimateMethodItem.Rekomendasi = rtes.Item_tabel[i].Method.Item[0].Rekomendasi
			dRatingTenderEstimateMethodItem.T_efficiency_result = rtes.Item_tabel[i].Method.Item[0].Item_value * rtesdetailresult.Price

			rtes.Item_tabel[i].Method.Item = append(rtes.Item_tabel[i].Method.Item, dRatingTenderEstimateMethodItem)
		}

		select_item_b_id := rtes.Item_tabel[i].Method.Item[1].Id_analisa_method

		rtesmethodresult = RatingTenderEstimateMethodResult{}

		rtesmethodresult.Id_analisa_method = 0
		rtesmethodresult.Item_method = ""
		rtesmethodresult.Item_value = 0
		rtesmethodresult.Rekomendasi = 0
		rtesmethodresult.T_efficiency_result = 0

		if rtes.Item_tabel[i].Method.Select_method == rtes.Item_tabel[i].Method.Item[0].Id_analisa_method {
			rtesmethodresult.Id_analisa_method = rtes.Item_tabel[i].Method.Item[0].Id_analisa_method
			rtesmethodresult.Item_method = rtes.Item_tabel[i].Method.Item[0].Item_method
			rtesmethodresult.Item_value = rtes.Item_tabel[i].Method.Item[0].Item_value
			rtesmethodresult.Rekomendasi = rtes.Item_tabel[i].Method.Item[0].Rekomendasi
			rtesmethodresult.T_efficiency_result = rtes.Item_tabel[i].Method.Item[0].Item_value * rtesdetailresult.Price
		} else if rtes.Item_tabel[i].Method.Select_method == rtes.Item_tabel[i].Method.Item[1].Id_analisa_method {
			rtesmethodresult.Id_analisa_method = rtes.Item_tabel[i].Method.Item[1].Id_analisa_method
			rtesmethodresult.Item_method = rtes.Item_tabel[i].Method.Item[1].Item_method
			rtesmethodresult.Item_value = rtes.Item_tabel[i].Method.Item[1].Item_value
			rtesmethodresult.Rekomendasi = rtes.Item_tabel[i].Method.Item[1].Rekomendasi
			rtesmethodresult.T_efficiency_result = rtes.Item_tabel[i].Method.Item[0].Item_value * rtesdetailresult.Price
		}

		rtesdetailresult.Method = rtesmethodresult

		semethodRatEloA := models.SegRatingElo{}
		semethodRatEloAGotten, err := semethodRatEloA.FindAllSegByIdTypeAnalisaBarang(server.DB, 1, select_item_a_id, id_barang)

		semethodRatEloB := models.SegRatingElo{}
		semethodRatEloBGotten, err := semethodRatEloB.FindAllSegByIdTypeAnalisaBarang(server.DB, 1, select_item_b_id, id_barang)

		select_item_a := float64(semethodRatEloAGotten.Rating)
		select_item_b := float64(semethodRatEloBGotten.Rating)

		e_select_item_a := 1 / (1 + (math.Pow(10, (select_item_b-select_item_a)/400)))
		e_select_item_b := 1 / (1 + (math.Pow(10, (select_item_a-select_item_b)/400)))

		if select_item == select_item_a_id {
			// Update Rating use K-factor when A selected
			select_item_a_1, err := strconv.ParseFloat(FormatFloat(float64((16*(1-e_select_item_a))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_a, err = strconv.ParseFloat(FormatFloat(float64(select_item_a+select_item_a_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b_1, err := strconv.ParseFloat(FormatFloat(float64((16*(0-e_select_item_b))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b, err = strconv.ParseFloat(FormatFloat(float64(select_item_b+select_item_b_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
		} else if select_item == select_item_b_id {
			// Update Rating use K-factor when B selected
			select_item_a_1, err := strconv.ParseFloat(FormatFloat(float64((16*(0-e_select_item_a))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_a, err = strconv.ParseFloat(FormatFloat(float64(select_item_a+select_item_a_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b_1, err := strconv.ParseFloat(FormatFloat(float64((16*(1-e_select_item_b))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b, err = strconv.ParseFloat(FormatFloat(float64(select_item_b+select_item_b_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
		}

		// save rating pilihan user => this.objects[this.userSelect]['calrating'] = select_item_a;

		dSegRatingElo := models.SegRatingElo{}
		dSegRatingElo.Rating = float32(select_item_a)

		_, err = dSegRatingElo.UpdateASegRatingEloByIdTypeAnalisaBarang(server.DB, 1, select_item_a_id, id_barang)
		if err != nil {

		}

		// dont update same value if object A = object B
		if select_item_a_id != select_item_b_id {
			// save rating pilihan user => this.objects[this.systemRecommendation]['calrating'] = select_item_b;

			dSegRatingElo := models.SegRatingElo{}
			dSegRatingElo.Rating = float32(select_item_b)

			_, err = dSegRatingElo.UpdateASegRatingEloByIdTypeAnalisaBarang(server.DB, 1, select_item_b_id, id_barang)
			if err != nil {

			}
		}

		////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		/////////////////////////////////////////// Innovation /////////////////////////////////////////////////////////////////

		id_barang = rtes.Item_tabel[i].Id_barang
		select_item = rtes.Item_tabel[i].Innovation.Select_innovation
		select_item_a_id = rtes.Item_tabel[i].Innovation.Item[0].Id_analisa_inovasi

		if len(rtes.Item_tabel[i].Innovation.Item) == 1 {
			dRatingTenderEstimateInovationItem := RatingTenderEstimateInovationItem{}
			dRatingTenderEstimateInovationItem.Id_analisa_inovasi = rtes.Item_tabel[i].Innovation.Item[0].Id_analisa_inovasi
			dRatingTenderEstimateInovationItem.Innovation_item = rtes.Item_tabel[i].Innovation.Item[0].Innovation_item
			dRatingTenderEstimateInovationItem.Innovation_value = rtes.Item_tabel[i].Innovation.Item[0].Innovation_value
			dRatingTenderEstimateInovationItem.Rekomendasi = rtes.Item_tabel[i].Innovation.Item[0].Rekomendasi
			dRatingTenderEstimateInovationItem.T_efficiency_result = rtes.Item_tabel[i].Innovation.Item[0].Innovation_value * rtesdetailresult.Price

			rtes.Item_tabel[i].Innovation.Item = append(rtes.Item_tabel[i].Innovation.Item, dRatingTenderEstimateInovationItem)
		}

		select_item_b_id = rtes.Item_tabel[i].Innovation.Item[1].Id_analisa_inovasi

		rtesinovationresult = RatingTenderEstimateInovationResult{}

		rtesinovationresult.Id_analisa_inovasi = 0
		rtesinovationresult.Innovation_item = ""
		rtesinovationresult.Innovation_value = 0
		rtesinovationresult.Rekomendasi = 0
		rtesinovationresult.T_efficiency_result = 0

		if rtes.Item_tabel[i].Innovation.Select_innovation == rtes.Item_tabel[i].Innovation.Item[0].Id_analisa_inovasi {
			rtesinovationresult.Id_analisa_inovasi = rtes.Item_tabel[i].Innovation.Item[0].Id_analisa_inovasi
			rtesinovationresult.Innovation_item = rtes.Item_tabel[i].Innovation.Item[0].Innovation_item
			rtesinovationresult.Innovation_value = rtes.Item_tabel[i].Innovation.Item[0].Innovation_value
			rtesinovationresult.Rekomendasi = rtes.Item_tabel[i].Innovation.Item[0].Rekomendasi
			rtesinovationresult.T_efficiency_result = rtes.Item_tabel[i].Innovation.Item[0].Innovation_value * rtesdetailresult.Price
		} else if rtes.Item_tabel[i].Innovation.Select_innovation == rtes.Item_tabel[i].Innovation.Item[1].Id_analisa_inovasi {
			rtesinovationresult.Id_analisa_inovasi = rtes.Item_tabel[i].Innovation.Item[1].Id_analisa_inovasi
			rtesinovationresult.Innovation_item = rtes.Item_tabel[i].Innovation.Item[1].Innovation_item
			rtesinovationresult.Innovation_value = rtes.Item_tabel[i].Innovation.Item[1].Innovation_value
			rtesinovationresult.Rekomendasi = rtes.Item_tabel[i].Innovation.Item[1].Rekomendasi
			rtesinovationresult.T_efficiency_result = rtes.Item_tabel[i].Innovation.Item[1].Innovation_value * rtesdetailresult.Price
		}

		rtesdetailresult.Innovation = rtesinovationresult

		semethodRatEloA = models.SegRatingElo{}
		semethodRatEloAGotten, err = semethodRatEloA.FindAllSegByIdTypeAnalisaBarang(server.DB, 2, select_item_a_id, id_barang)

		semethodRatEloB = models.SegRatingElo{}
		semethodRatEloBGotten, err = semethodRatEloB.FindAllSegByIdTypeAnalisaBarang(server.DB, 2, select_item_b_id, id_barang)

		select_item_a = float64(semethodRatEloAGotten.Rating)
		select_item_b = float64(semethodRatEloBGotten.Rating)

		e_select_item_a = 1 / (1 + (math.Pow(10, (select_item_b-select_item_a)/400)))
		e_select_item_b = 1 / (1 + (math.Pow(10, (select_item_a-select_item_b)/400)))

		if select_item == select_item_a_id {
			// Update Rating use K-factor when A selected
			select_item_a_1, err := strconv.ParseFloat(FormatFloat(float64((16*(1-e_select_item_a))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_a, err = strconv.ParseFloat(FormatFloat(float64(select_item_a+select_item_a_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b_1, err := strconv.ParseFloat(FormatFloat(float64((16*(0-e_select_item_b))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b, err = strconv.ParseFloat(FormatFloat(float64(select_item_b+select_item_b_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
		} else if select_item == select_item_b_id {
			// Update Rating use K-factor when B selected
			select_item_a_1, err := strconv.ParseFloat(FormatFloat(float64((16*(0-e_select_item_a))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_a, err = strconv.ParseFloat(FormatFloat(float64(select_item_a+select_item_a_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b_1, err := strconv.ParseFloat(FormatFloat(float64((16*(1-e_select_item_b))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b, err = strconv.ParseFloat(FormatFloat(float64(select_item_b+select_item_b_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
		}

		// save rating pilihan user => this.objects[this.userSelect]['calrating'] = select_item_a;

		dSegRatingElo = models.SegRatingElo{}
		dSegRatingElo.Rating = float32(select_item_a)

		_, err = dSegRatingElo.UpdateASegRatingEloByIdTypeAnalisaBarang(server.DB, 2, select_item_a_id, id_barang)
		if err != nil {

		}

		// dont update same value if object A = object B
		if select_item_a_id != select_item_b_id {
			// save rating pilihan user => this.objects[this.systemRecommendation]['calrating'] = select_item_b;

			dSegRatingElo := models.SegRatingElo{}
			dSegRatingElo.Rating = float32(select_item_b)

			_, err = dSegRatingElo.UpdateASegRatingEloByIdTypeAnalisaBarang(server.DB, 2, select_item_b_id, id_barang)
			if err != nil {

			}
		}

		////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		/////////////////////////////////////////// VE /////////////////////////////////////////////////////////////////

		id_barang = rtes.Item_tabel[i].Id_barang
		select_item = rtes.Item_tabel[i].Value_e.Select_value_e
		select_item_a_id = rtes.Item_tabel[i].Value_e.Item[0].Id_analisa_ve

		if len(rtes.Item_tabel[i].Value_e.Item) == 1 {
			dRatingTenderEstimateVeItem := RatingTenderEstimateVeItem{}
			dRatingTenderEstimateVeItem.Id_analisa_ve = rtes.Item_tabel[i].Value_e.Item[0].Id_analisa_ve
			dRatingTenderEstimateVeItem.Valuee_item = rtes.Item_tabel[i].Value_e.Item[0].Valuee_item
			dRatingTenderEstimateVeItem.Valuee_value = rtes.Item_tabel[i].Value_e.Item[0].Valuee_value
			dRatingTenderEstimateVeItem.Rekomendasi = rtes.Item_tabel[i].Value_e.Item[0].Rekomendasi
			dRatingTenderEstimateVeItem.T_efficiency_result = rtes.Item_tabel[i].Value_e.Item[0].Valuee_value * rtesdetailresult.Price

			rtes.Item_tabel[i].Value_e.Item = append(rtes.Item_tabel[i].Value_e.Item, dRatingTenderEstimateVeItem)
		}

		select_item_b_id = rtes.Item_tabel[i].Value_e.Item[1].Id_analisa_ve

		rtesveresult = RatingTenderEstimateVeResult{}

		rtesveresult.Id_analisa_ve = 0
		rtesveresult.Valuee_item = ""
		rtesveresult.Valuee_value = 0
		rtesveresult.Rekomendasi = 0
		rtesveresult.T_efficiency_result = 0

		if rtes.Item_tabel[i].Value_e.Select_value_e == rtes.Item_tabel[i].Value_e.Item[0].Id_analisa_ve {
			rtesveresult.Id_analisa_ve = rtes.Item_tabel[i].Value_e.Item[0].Id_analisa_ve
			rtesveresult.Valuee_item = rtes.Item_tabel[i].Value_e.Item[0].Valuee_item
			rtesveresult.Valuee_value = rtes.Item_tabel[i].Value_e.Item[0].Valuee_value
			rtesveresult.Rekomendasi = rtes.Item_tabel[i].Value_e.Item[0].Rekomendasi
			rtesveresult.T_efficiency_result = rtes.Item_tabel[i].Value_e.Item[0].Valuee_value * rtesdetailresult.Price
		} else if rtes.Item_tabel[i].Value_e.Select_value_e == rtes.Item_tabel[i].Value_e.Item[1].Id_analisa_ve {
			rtesveresult.Id_analisa_ve = rtes.Item_tabel[i].Value_e.Item[1].Id_analisa_ve
			rtesveresult.Valuee_item = rtes.Item_tabel[i].Value_e.Item[1].Valuee_item
			rtesveresult.Valuee_value = rtes.Item_tabel[i].Value_e.Item[1].Valuee_value
			rtesveresult.Rekomendasi = rtes.Item_tabel[i].Value_e.Item[1].Rekomendasi
			rtesveresult.T_efficiency_result = rtes.Item_tabel[i].Value_e.Item[1].Valuee_value * rtesdetailresult.Price
		}

		rtesdetailresult.Value_e = rtesveresult

		semethodRatEloA = models.SegRatingElo{}
		semethodRatEloAGotten, err = semethodRatEloA.FindAllSegByIdTypeAnalisaBarang(server.DB, 3, select_item_a_id, id_barang)

		semethodRatEloB = models.SegRatingElo{}
		semethodRatEloBGotten, err = semethodRatEloB.FindAllSegByIdTypeAnalisaBarang(server.DB, 3, select_item_b_id, id_barang)

		select_item_a = float64(semethodRatEloAGotten.Rating)
		select_item_b = float64(semethodRatEloBGotten.Rating)

		e_select_item_a = 1 / (1 + (math.Pow(10, (select_item_b-select_item_a)/400)))
		e_select_item_b = 1 / (1 + (math.Pow(10, (select_item_a-select_item_b)/400)))

		if select_item == select_item_a_id {
			// Update Rating use K-factor when A selected
			select_item_a_1, err := strconv.ParseFloat(FormatFloat(float64((16*(1-e_select_item_a))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_a, err = strconv.ParseFloat(FormatFloat(float64(select_item_a+select_item_a_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b_1, err := strconv.ParseFloat(FormatFloat(float64((16*(0-e_select_item_b))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b, err = strconv.ParseFloat(FormatFloat(float64(select_item_b+select_item_b_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
		} else if select_item == select_item_b_id {
			// Update Rating use K-factor when B selected
			select_item_a_1, err := strconv.ParseFloat(FormatFloat(float64((16*(0-e_select_item_a))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_a, err = strconv.ParseFloat(FormatFloat(float64(select_item_a+select_item_a_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b_1, err := strconv.ParseFloat(FormatFloat(float64((16*(1-e_select_item_b))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b, err = strconv.ParseFloat(FormatFloat(float64(select_item_b+select_item_b_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
		}

		// save rating pilihan user => this.objects[this.userSelect]['calrating'] = select_item_a;

		dSegRatingElo = models.SegRatingElo{}
		dSegRatingElo.Rating = float32(select_item_a)

		_, err = dSegRatingElo.UpdateASegRatingEloByIdTypeAnalisaBarang(server.DB, 3, select_item_a_id, id_barang)
		if err != nil {

		}

		// dont update same value if object A = object B
		if select_item_a_id != select_item_b_id {
			// save rating pilihan user => this.objects[this.systemRecommendation]['calrating'] = select_item_b;

			dSegRatingElo := models.SegRatingElo{}
			dSegRatingElo.Rating = float32(select_item_b)

			_, err = dSegRatingElo.UpdateASegRatingEloByIdTypeAnalisaBarang(server.DB, 3, select_item_b_id, id_barang)
			if err != nil {

			}
		}

		////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		/////////////////////////////////////////// FE /////////////////////////////////////////////////////////////////

		id_barang = rtes.Item_tabel[i].Id_barang
		select_item = rtes.Item_tabel[i].Finance_e.Select_finance
		select_item_a_id = rtes.Item_tabel[i].Finance_e.Item[0].Id_analisa_fe

		if len(rtes.Item_tabel[i].Finance_e.Item) == 1 {
			dRatingTenderEstimateFeItem := RatingTenderEstimateFeItem{}
			dRatingTenderEstimateFeItem.Id_analisa_fe = rtes.Item_tabel[i].Finance_e.Item[0].Id_analisa_fe
			dRatingTenderEstimateFeItem.Finance_item = rtes.Item_tabel[i].Finance_e.Item[0].Finance_item
			dRatingTenderEstimateFeItem.Finance_value = rtes.Item_tabel[i].Finance_e.Item[0].Finance_value
			dRatingTenderEstimateFeItem.Rekomendasi = rtes.Item_tabel[i].Finance_e.Item[0].Rekomendasi
			dRatingTenderEstimateFeItem.T_efficiency_result = rtes.Item_tabel[i].Finance_e.Item[0].Finance_value * rtesdetailresult.Price

			rtes.Item_tabel[i].Finance_e.Item = append(rtes.Item_tabel[i].Finance_e.Item, dRatingTenderEstimateFeItem)
		}

		select_item_b_id = rtes.Item_tabel[i].Finance_e.Item[1].Id_analisa_fe

		rtesferesult = RatingTenderEstimateFeResult{}

		rtesferesult.Id_analisa_fe = 0
		rtesferesult.Finance_item = ""
		rtesferesult.Finance_value = 0
		rtesferesult.Rekomendasi = 0
		rtesferesult.T_efficiency_result = 0

		if rtes.Item_tabel[i].Finance_e.Select_finance == rtes.Item_tabel[i].Finance_e.Item[0].Id_analisa_fe {
			rtesferesult.Id_analisa_fe = rtes.Item_tabel[i].Finance_e.Item[0].Id_analisa_fe
			rtesferesult.Finance_item = rtes.Item_tabel[i].Finance_e.Item[0].Finance_item
			rtesferesult.Finance_value = rtes.Item_tabel[i].Finance_e.Item[0].Finance_value
			rtesferesult.Rekomendasi = rtes.Item_tabel[i].Finance_e.Item[0].Rekomendasi
			rtesferesult.T_efficiency_result = rtes.Item_tabel[i].Finance_e.Item[0].Finance_value * rtesdetailresult.Price
		} else if rtes.Item_tabel[i].Finance_e.Select_finance == rtes.Item_tabel[i].Finance_e.Item[1].Id_analisa_fe {
			rtesferesult.Id_analisa_fe = rtes.Item_tabel[i].Finance_e.Item[1].Id_analisa_fe
			rtesferesult.Finance_item = rtes.Item_tabel[i].Finance_e.Item[1].Finance_item
			rtesferesult.Finance_value = rtes.Item_tabel[i].Finance_e.Item[1].Finance_value
			rtesferesult.Rekomendasi = rtes.Item_tabel[i].Finance_e.Item[1].Rekomendasi
			rtesferesult.T_efficiency_result = rtes.Item_tabel[i].Finance_e.Item[1].Finance_value * rtesdetailresult.Price
		}

		rtesdetailresult.Finance_e = rtesferesult

		semethodRatEloA = models.SegRatingElo{}
		semethodRatEloAGotten, err = semethodRatEloA.FindAllSegByIdTypeAnalisaBarang(server.DB, 3, select_item_a_id, id_barang)

		semethodRatEloB = models.SegRatingElo{}
		semethodRatEloBGotten, err = semethodRatEloB.FindAllSegByIdTypeAnalisaBarang(server.DB, 3, select_item_b_id, id_barang)

		select_item_a = float64(semethodRatEloAGotten.Rating)
		select_item_b = float64(semethodRatEloBGotten.Rating)

		e_select_item_a = 1 / (1 + (math.Pow(10, (select_item_b-select_item_a)/400)))
		e_select_item_b = 1 / (1 + (math.Pow(10, (select_item_a-select_item_b)/400)))

		if select_item == select_item_a_id {
			// Update Rating use K-factor when A selected
			select_item_a_1, err := strconv.ParseFloat(FormatFloat(float64((16*(1-e_select_item_a))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_a, err = strconv.ParseFloat(FormatFloat(float64(select_item_a+select_item_a_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b_1, err := strconv.ParseFloat(FormatFloat(float64((16*(0-e_select_item_b))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b, err = strconv.ParseFloat(FormatFloat(float64(select_item_b+select_item_b_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
		} else if select_item == select_item_b_id {
			// Update Rating use K-factor when B selected
			select_item_a_1, err := strconv.ParseFloat(FormatFloat(float64((16*(0-e_select_item_a))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_a, err = strconv.ParseFloat(FormatFloat(float64(select_item_a+select_item_a_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b_1, err := strconv.ParseFloat(FormatFloat(float64((16*(1-e_select_item_b))), 0), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			select_item_b, err = strconv.ParseFloat(FormatFloat(float64(select_item_b+select_item_b_1), 4), 32)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
		}

		// save rating pilihan user => this.objects[this.userSelect]['calrating'] = select_item_a;

		dSegRatingElo = models.SegRatingElo{}
		dSegRatingElo.Rating = float32(select_item_a)

		_, err = dSegRatingElo.UpdateASegRatingEloByIdTypeAnalisaBarang(server.DB, 3, select_item_a_id, id_barang)
		if err != nil {

		}

		// dont update same value if object A = object B
		if select_item_a_id != select_item_b_id {
			// save rating pilihan user => this.objects[this.systemRecommendation]['calrating'] = select_item_b;

			dSegRatingElo := models.SegRatingElo{}
			dSegRatingElo.Rating = float32(select_item_b)

			_, err = dSegRatingElo.UpdateASegRatingEloByIdTypeAnalisaBarang(server.DB, 3, select_item_b_id, id_barang)
			if err != nil {

			}
		}

		////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		rtesdetailresult.T_efficiency = rtesmethodresult.Item_value + rtesinovationresult.Innovation_value + rtesveresult.Valuee_value + rtesferesult.Finance_value

		arrtesdetailresult = append(arrtesdetailresult, rtesdetailresult)
	}

	rtesresult.Item_tabel = arrtesdetailresult

	dRatingTenderEstimateProjectBoostResult := RatingTenderEstimateProjectBoostResult{}
	arrRatingTenderEstimateProjectBoostResult := []RatingTenderEstimateProjectBoostResult{}

	dSegAnalisaMethodDetail := models.SegAnalisaMethodDetail{}

	for i, _ := range rtes.Project_boost { // loop through the files one by one
		dRatingTenderEstimateProjectBoostResult = RatingTenderEstimateProjectBoostResult{}
		dRatingTenderEstimateProjectBoostResult.Select_object = rtes.Project_boost[i].Select_object
		dRatingTenderEstimateProjectBoostResult.Select_name = rtes.Project_boost[i].Select_name
		dRatingTenderEstimateProjectBoostResult.Value_object = rtes.Project_boost[i].Value_object
		dRatingTenderEstimateProjectBoostResult.Value_name = rtes.Project_boost[i].Value_name

		t_efficiency_project_boost := float32(0)
		t_efficiency_project_boost_result := float32(0)
		for i, _ := range rtes.Item_tabel { // loop through the files one by one
			id_barang := rtes.Item_tabel[i].Id_barang

			dSegAnalisaMethodDetail = models.SegAnalisaMethodDetail{}
			dSegAnalisaMethodDetailGotten, err := dSegAnalisaMethodDetail.FindSegAnalisaMethodDetailByAnalisaIdbarang(server.DB, uint32(dRatingTenderEstimateProjectBoostResult.Value_object), id_barang)
			if err != nil {

			}

			t_efficiency_project_boost = t_efficiency_project_boost + dSegAnalisaMethodDetailGotten.Eficiency
			t_efficiency_project_boost_result = t_efficiency_project_boost * dSegAnalisaMethodDetailGotten.Price
		}

		dRatingTenderEstimateProjectBoostResult.T_efficiency_project_boost = t_efficiency_project_boost
		dRatingTenderEstimateProjectBoostResult.T_efficiency_project_boost_result = t_efficiency_project_boost_result

		arrRatingTenderEstimateProjectBoostResult = append(arrRatingTenderEstimateProjectBoostResult, dRatingTenderEstimateProjectBoostResult)
	}

	rtesresult.Project_boost = arrRatingTenderEstimateProjectBoostResult

	responses.JSON(w, http.StatusOK, rtesresult)
}

func (server *Server) SegTenderEstimateDelete(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)

	stesproyekboost := models.SegTenderEstimateProyekBoost{}
	stesdetail := models.SegTenderEstimateDetail{}
	stesdetailres := models.SegTenderEstimateDetailResult{}
	stes := models.SegTenderEstimate{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = stesproyekboost.DeleteASegTenderEstimateProyekBoostByIdEstimate(server.DB, uint32(uid))
	if err != nil {

	}

	_, err = stesdetail.DeleteASegTenderEstimateDetailByIdEstimate(server.DB, uint32(uid))
	if err != nil {

	}

	_, err = stesdetailres.DeleteASegTenderEstimateDetailResultByIdEstimate(server.DB, uint32(uid))
	if err != nil {

	}

	_, err = stes.DeleteASegTenderEstimate(server.DB, uint32(uid))
	if err != nil {

	}

	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
