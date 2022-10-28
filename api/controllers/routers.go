package controllers

import "github.com/amarmaulana95/api_seg-v2/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateUser))).Methods("POST") //OK

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUsers))).Methods("GET") // OK
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.GetUser)).Methods("GET")                            // PL
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("POST")

	// s.Router.HandleFunc("/register_token", middlewares.SetMiddlewareJSON(s.ValidasiToken)).Methods("GET")

	s.Router.HandleFunc("/lokasi", middlewares.SetMiddlewareJSON(s.GetLokasis)).Methods("GET")
	s.Router.HandleFunc("/lokasi/{prop_id}", middlewares.SetMiddlewareJSON(s.GetLokasi)).Methods("GET")
	s.Router.HandleFunc("/get_analisa_exception", middlewares.SetMiddlewareJSON(s.GetSegAnalisaException)).Methods("GET")
	s.Router.HandleFunc("/get_barang", middlewares.SetMiddlewareJSON(s.GetItemProcurments)).Methods("GET") //-!- ?q={param}&id_provinsi={id} -!-//

	s.Router.HandleFunc("/seg_analisa_types", middlewares.SetMiddlewareAuthentication(s.CreateSegAnalisaType)).Methods("POST")
	s.Router.HandleFunc("/seg_analisa_types", middlewares.SetMiddlewareJSON(s.GetSegAnalisaTypes)).Methods("GET")
	s.Router.HandleFunc("/seg_analisa_types/{id}", middlewares.SetMiddlewareAuthentication(s.GetSegAnalisaType)).Methods("GET")
	s.Router.HandleFunc("/seg_analisa_types/{id}", middlewares.SetMiddlewareAuthentication(middlewares.SetMiddlewareAuthentication(s.UpdateSegAnalisaType))).Methods("POST")
	s.Router.HandleFunc("/seg_analisa_types/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteSegAnalisaType)).Methods("DELETE")

	s.Router.HandleFunc("/analisa_inovasi", middlewares.SetMiddlewareJSON(s.SegAnalisaInovasiAll)).Methods("GET")

	//===================================================METODE=============================================================//
	s.Router.HandleFunc("/analisa_method", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodAll)).Methods("GET")
	s.Router.HandleFunc("/analisa_method", middlewares.SetMiddlewareAuthentication(s.SegAnalisaMethodInsert)).Methods("POST")
	s.Router.HandleFunc("/analisa_method/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaMethodSingle)).Methods("GET")
	s.Router.HandleFunc("/analisa_method_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaMethodDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_method_detail_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaMethodDetailDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_method_exception_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaMethodExceptionDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_inovasi_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaInovasiDelete)).Methods("POST")

	//----------------------------------------------VALUE ENGINEERING------------------------------------------------------------//

	s.Router.HandleFunc("/analisa_value_enginering", middlewares.SetMiddlewareJSON(s.SegAnalisaValueEngineringAll)).Methods("GET")
	s.Router.HandleFunc("/analisa_value_enginering", middlewares.SetMiddlewareAuthentication(s.SegAnalisaValueEngineringAll)).Methods("POST")
	s.Router.HandleFunc("/analisa_value_enginering_update/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaValueEngineringUpdate)).Methods("POST")
	s.Router.HandleFunc("/analisa_value_enginering_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaValueEngineringDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_value_enginering_attachment_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaValueEngineringAttachmentDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_value_enginering_detail_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaValueEngineringDetailDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_value_enginering_exception_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaValueEngineringExceptionDelete)).Methods("POST")

	//------------------------------------FINANCE ENGINEERING---------------------------------------------------------------------------------------//
	s.Router.HandleFunc("/analisa_finance_enginering", middlewares.SetMiddlewareJSON(s.SegAnalisaFinanceEngineringAll)).Methods("GET")
	s.Router.HandleFunc("/analisa_finance_enginering", middlewares.SetMiddlewareAuthentication(s.SegAnalisaFinanceEngineringInsert)).Methods("POST")
	s.Router.HandleFunc("/analisa_finance_enginering_update/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaFinanceEngineringUpdate)).Methods("POST")
	s.Router.HandleFunc("/analisa_finance_enginering_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaFinanceEngineringDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_finance_enginering_attachment_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaFinanceEngineringAttachmentDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_finance_enginering_detail_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaFinanceEngineringDetailDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_finance_enginering_exception_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegAnalisaFinanceEngineringExceptionDelete)).Methods("POST")

	//-------------------------------------------DASHBOARD----------------------------------------------------------------------------------------//
	s.Router.HandleFunc("/dashboard_total_eficiency", middlewares.SetMiddlewareJSON(s.GetSegDashboardEficiency)).Methods("GET")
	s.Router.HandleFunc("/dashboard_total_bulanan_analisa", middlewares.SetMiddlewareJSON(s.GetSegDashboardTotalBulananAnalisa)).Methods("GET")
	s.Router.HandleFunc("/dashboard_total_analisa", middlewares.SetMiddlewareJSON(s.GetSegDashboardTotalAnalisa)).Methods("GET")

	//-----------------------------------------ESTIMATE-----------------------------------------------------------------//
	s.Router.HandleFunc("/estimate_method", middlewares.SetMiddlewareJSON(s.SegAnalisaMethod)).Methods("GET")                        // SegEstimateMethod q=&id_barang=&id_provinsi=1
	s.Router.HandleFunc("/estimate_inovasi", middlewares.SetMiddlewareJSON(s.SegAnalisaInovasi)).Methods("GET")                      // SegEstimateInovasi
	s.Router.HandleFunc("/estimate_value_enginering", middlewares.SetMiddlewareJSON(s.SegAnalisaValueEnginering)).Methods("GET")     // SegEstimateValueENginering
	s.Router.HandleFunc("/estimate_finance_enginering", middlewares.SetMiddlewareJSON(s.SegAnalisaFinanceENginering)).Methods("GET") // SegEstimateFinanceENginering
	s.Router.HandleFunc("/estimate_project_boost", middlewares.SetMiddlewareJSON(s.SegAnalisaProjectBoost)).Methods("GET")           // SegEstimateProjectBoost

	s.Router.HandleFunc("/seg_tender_estimate", middlewares.SetMiddlewareAuthentication(s.SegTenderEstimateCalculation)).Methods("POST")
	s.Router.HandleFunc("/seg_tender_estimate_rating/{id}", middlewares.SetMiddlewareJSON(s.SegTenderEstimateCalculationRating)).Methods("GET")
	s.Router.HandleFunc("/seg_tender_estimate_rating_calibrate_firts", middlewares.SetMiddlewareJSON(s.SegTenderEstimateCalibrateRatingFirst)).Methods("GET")
	s.Router.HandleFunc("/seg_tender_estimate_rating_save", middlewares.SetMiddlewareAuthentication(s.SegTenderEstimateResultSave)).Methods("POST")
	s.Router.HandleFunc("/seg_tender_estimate_delete/{id}", middlewares.SetMiddlewareAuthentication(s.SegTenderEstimateDelete)).Methods("POST")

}
