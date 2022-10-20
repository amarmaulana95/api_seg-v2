package controllers

import "github.com/amarmaulana95/api_seg-v2/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/lokasi", middlewares.SetMiddlewareJSON(s.GetLokasis)).Methods("GET")
	s.Router.HandleFunc("/lokasi/{prop_id}", middlewares.SetMiddlewareJSON(s.GetLokasi)).Methods("GET")
	s.Router.HandleFunc("/get_analisa_exception", middlewares.SetMiddlewareJSON(s.GetSegAnalisaException)).Methods("GET")
	s.Router.HandleFunc("/get_barang", middlewares.SetMiddlewareJSON(s.GetItemProcurments)).Methods("GET") //-!- ?q={param}&id_provinsi={id} -!-//
	s.Router.HandleFunc("/seg_analisa_types", middlewares.SetMiddlewareJSON(s.GetSegAnalisaTypes)).Methods("GET")
	s.Router.HandleFunc("/seg_analisa_types/{id}", middlewares.SetMiddlewareJSON(s.GetSegAnalisaType)).Methods("GET")
	s.Router.HandleFunc("/seg_analisa_types/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareJSON(s.UpdateSegAnalisaType))).Methods("PUT")
	s.Router.HandleFunc("/seg_analisa_types/{id}", middlewares.SetMiddlewareJSON(s.DeleteSegAnalisaType)).Methods("DELETE")

	s.Router.HandleFunc("/analisa_value_enginering", middlewares.SetMiddlewareJSON(s.SegAnalisaValueEngineringAll)).Methods("GET")
	s.Router.HandleFunc("/analisa_method", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodAll)).Methods("GET")
	s.Router.HandleFunc("/analisa_inovasi", middlewares.SetMiddlewareJSON(s.SegAnalisaInovasiAll)).Methods("GET")
	s.Router.HandleFunc("/analisa_finance_enginering", middlewares.SetMiddlewareJSON(s.SegAnalisaFinanceEngineringAll)).Methods("GET")

	s.Router.HandleFunc("/analisa_method", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodInsert)).Methods("POST")

	s.Router.HandleFunc("/analisa_method/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodSingle)).Methods("GET")

	s.Router.HandleFunc("/analisa_method_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_method_detail_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodDetailDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_method_exception_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodExceptionDelete)).Methods("POST")

	s.Router.HandleFunc("/analisa_inovasi_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaInovasiDelete)).Methods("POST")

	//----------------------------------------------VALUE ENGINEERING------------------------------------------------------------//
	s.Router.HandleFunc("/analisa_value_enginering_update/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaValueEngineringUpdate)).Methods("POST")
	s.Router.HandleFunc("/analisa_value_enginering_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaValueEngineringDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_value_enginering_attachment_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaValueEngineringAttachmentDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_value_enginering_detail_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaValueEngineringDetailDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_value_enginering_exception_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaValueEngineringExceptionDelete)).Methods("POST")

	//------------------------------------FINANCE ENGINEERING---------------------------------------------------------------------------------------//
	s.Router.HandleFunc("/analisa_finance_enginering_update/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaFinanceEngineringUpdate)).Methods("POST")
	s.Router.HandleFunc("/analisa_finance_enginering_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaFinanceEngineringDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_finance_enginering_attachment_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaFinanceEngineringAttachmentDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_finance_enginering_detail_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaFinanceEngineringDetailDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_finance_enginering_exception_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaFinanceEngineringExceptionDelete)).Methods("POST")

	//-------------------------------------------DASHBOARD----------------------------------------------------------------------------------------//
	s.Router.HandleFunc("/dashboard_total_eficiency", middlewares.SetMiddlewareJSON(s.GetSegDashboardEficiency)).Methods("GET")
	s.Router.HandleFunc("/dashboard_total_bulanan_analisa", middlewares.SetMiddlewareJSON(s.GetSegDashboardTotalBulananAnalisa)).Methods("GET")
	s.Router.HandleFunc("/dashboard_total_analisa", middlewares.SetMiddlewareJSON(s.GetSegDashboardTotalAnalisa)).Methods("GET")
}
