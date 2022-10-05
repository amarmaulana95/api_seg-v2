package controllers

import "github.com/amarmaulana95/api_seg-v2/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/lokasi", middlewares.SetMiddlewareJSON(s.GetLokasis)).Methods("GET")
	s.Router.HandleFunc("/lokasi/{prop_id}", middlewares.SetMiddlewareJSON(s.GetLokasi)).Methods("GET")
	s.Router.HandleFunc("/get_analisa_exception", middlewares.SetMiddlewareJSON(s.GetSegAnalisaException)).Methods("GET")
	s.Router.HandleFunc("/get_barang", middlewares.SetMiddlewareJSON(s.GetItemProcurments)).Methods("GET") //-!- ?q={param}&id_provinsi={id} -!-//
	s.Router.HandleFunc("/seg_analisa_types", middlewares.SetMiddlewareJSON(s.GetSegAnalisaTypes)).Methods("GET")

	s.Router.HandleFunc("/analisa_method", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodAll)).Methods("GET")
	s.Router.HandleFunc("/analisa_method", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodInsert)).Methods("POST")

	s.Router.HandleFunc("/analisa_method/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodSingle)).Methods("GET")

	s.Router.HandleFunc("/analisa_method_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_method_detail_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodDetailDelete)).Methods("POST")
	s.Router.HandleFunc("/analisa_method_exception_delete/{id}", middlewares.SetMiddlewareJSON(s.SegAnalisaMethodExceptionDelete)).Methods("POST")
}
