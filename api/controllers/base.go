package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	// server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

// func getDir() string {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		fmt.Printf(err.Error())
// 	}

// 	dir_uploads := dir + "/uploads/"

// 	return dir_uploads
// }

// func isError(err error) bool {
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	return (err != nil)
// }

type Token struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Meta     MyData `json:"meta"`
}

type MyData struct {
	Tokens string `json:"token"`
}

type DashEfi struct {
	Data_a [2]float32 `json:"data_a"`
	Persen string     `json:"persen"`
}

type DataDetail struct {
	Id             uint32  `json:"detail_id"`
	Id_barang      string  `json:"id_barang"`
	Label_barang   string  `json:"analisa_exception_label"`
	Eficiency      float32 `json:"analisa_exception_eficiency"`
	Eficiency_type float32 `json:"eficiency_type"`
	Price          float32 `json:"price"`
}

type DataAttachment struct {
	Id             uint32 `json:"file_id"`
	File_name      string `json:"file_name"`
	Attachment     string `json:"attachment"`
	Path_file_name string `json:"path_attachment"`
}

type DataException struct {
	Id                   uint32 `json:"exception_id"`
	Analisa_type         uint32 `json:"id_analisa_type"`
	Label_type           string `json:"label_type"`
	Id_analisa_exception uint32 `json:"id_analisa_exception"`
	Label_exception      string `json:"label_exception"`
}

type ResponAnalisa struct {
	Id                  uint32      `json:"id"`
	Id_analisa_type     uint32      `json:"id_analisa_type"`
	Name                string      `json:"name"`
	Description         string      `json:"description"`
	Location            string      `json:"location"`
	Location_name       string      `json:"location_name"`
	Status_proyek_boost uint32      `json:"status_proyek_boost"`
	DataAttachment      interface{} `json:"document"`
	DataDetail          interface{} `json:"eficiency"`
	DataException       interface{} `json:"combine"`
}

type ArrDataException struct {
	DataArr []DataException
}

type ResponStatus struct {
	Status  uint32      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponStatusData struct {
	Status  uint32      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponStatusDataView struct {
	Page        uint32      `json:"page"`
	Per_page    uint32      `json:"per_page"`
	Total_pages uint32      `json:"total_pages"`
	Total       uint32      `json:"total"`
	Data        interface{} `json:"data"`
}

type RawTenderEstimateCalculation struct {
	Project_name      string `json:"project_name"`
	Construction_type string `json:"construction_type"`
	Building          string `json:"building"`
	Province          string `json:"province"`
	Tipe_class        string `json:"tipe_class"`
	Component_item    []RawTEComponent_item
	Project_boost     []RawTEProject_boost
}

type RawTEComponent_item struct {
	Item        string  `json:"item"`
	Price       float32 `json:"price"`
	Method      uint32  `json:"method"`
	Innovation  uint32  `json:"innovation"`
	Value_eng   uint32  `json:"value_eng"`
	Finance_eng uint32  `json:"finance_eng"`
}

type RawTEProject_boost struct {
	Select_object uint32 `json:"select_object"`
	Name_object   uint32 `json:"name_object"`
}

type RatingTenderEstimate struct {
	Id            uint32                             `json:"id"`
	Project       string                             `json:"project"`
	Location      string                             `json:"location"`
	Const_type    string                             `json:"const_type"`
	BuiLding      string                             `json:"buiLding"`
	Class_project string                             `json:"class_project"`
	Item_tabel    []RatingTenderEstimateDetail       `json:"item_tabel"`
	Project_boost []RatingTenderEstimateProjectBoost `json:"project_boost"`
}

type RatingTenderEstimateDetail struct {
	Id_barang  string                        `json:"id_barang"`
	Item_name  string                        `json:"item_name"`
	Price      float32                       `json:"price"`
	Method     RatingTenderEstimateMethod    `json:"method"`
	Innovation RatingTenderEstimateInovation `json:"innovation"`
	Value_e    RatingTenderEstimateVe        `json:"value_e"`
	Finance_e  RatingTenderEstimateFe        `json:"finance_e"`
}

type RatingTenderEstimateMethod struct {
	Select_method uint32                           `json:"select_method"`
	Item          []RatingTenderEstimateMethodItem `json:"item"`
}

type RatingTenderEstimateMethodItem struct {
	Id_analisa_method   uint32  `json:"id_analisa_method"`
	Item_method         string  `json:"item_method"`
	Item_value          float32 `json:"item_value"`
	Rekomendasi         int32   `json:"rekomendasi"`
	T_efficiency_result float32 `json:"t_efficiency_result"`
}

type RatingTenderEstimateInovation struct {
	Select_innovation uint32                              `json:"select_innovation"`
	Item              []RatingTenderEstimateInovationItem `json:"item"`
}

type RatingTenderEstimateInovationItem struct {
	Id_analisa_inovasi  uint32  `json:"id_analisa_inovasi"`
	Innovation_item     string  `json:"innovation_item"`
	Innovation_value    float32 `json:"innovation_value"`
	Rekomendasi         int32   `json:"rekomendasi"`
	T_efficiency_result float32 `json:"t_efficiency_result"`
}

type RatingTenderEstimateVe struct {
	Select_value_e uint32                       `json:"select_value_e"`
	Item           []RatingTenderEstimateVeItem `json:"item"`
}

type RatingTenderEstimateVeItem struct {
	Id_analisa_ve       uint32  `json:"id_analisa_ve"`
	Valuee_item         string  `json:"valuee_item"`
	Valuee_value        float32 `json:"valuee_value"`
	Rekomendasi         int32   `json:"rekomendasi"`
	T_efficiency_result float32 `json:"t_efficiency_result"`
}

type RatingTenderEstimateFe struct {
	Select_finance uint32                       `json:"select_finance"`
	Item           []RatingTenderEstimateFeItem `json:"item"`
}

type RatingTenderEstimateFeItem struct {
	Id_analisa_fe       uint32  `json:"id_analisa_fe"`
	Finance_item        string  `json:"finance_item"`
	Finance_value       float32 `json:"finance_value"`
	Rekomendasi         int32   `json:"rekomendasi"`
	T_efficiency_result float32 `json:"t_efficiency_result"`
}

type RatingTenderEstimateProjectBoost struct {
	Select_object                     uint32  `json:"select_object"`
	Select_name                       string  `json:"select_name"`
	Value_object                      uint32  `json:"value_object"`
	Value_name                        string  `json:"value_name"`
	T_efficiency_project_boost        float32 `json:"t_efficiency_project_boost"`
	T_efficiency_project_boost_result float32 `json:"t_efficiency_project_boost_result"`
}

type RatingTenderEstimateResult struct {
	Id            uint32                                   `json:"id"`
	Project       string                                   `json:"project"`
	Location      string                                   `json:"location"`
	Const_type    string                                   `json:"const_type"`
	BuiLding      string                                   `json:"buiLding"`
	Class_project string                                   `json:"class_project"`
	Item_tabel    []RatingTenderEstimateDetailResult       `json:"item_tabel"`
	Project_boost []RatingTenderEstimateProjectBoostResult `json:"project_boost"`
}

type RatingTenderEstimateDetailResult struct {
	Id_barang    string                              `json:"id_barang"`
	Item_name    string                              `json:"item_name"`
	Price        float32                             `json:"price"`
	Method       RatingTenderEstimateMethodResult    `json:"method"`
	Innovation   RatingTenderEstimateInovationResult `json:"innovation"`
	Value_e      RatingTenderEstimateVeResult        `json:"value_e"`
	Finance_e    RatingTenderEstimateFeResult        `json:"finance_e"`
	T_efficiency float32                             `json:"t_efficiency"`
}

type RatingTenderEstimateMethodResult struct {
	Id_analisa_method   uint32  `json:"id_analisa_method"`
	Item_method         string  `json:"item_method"`
	Item_value          float32 `json:"item_value"`
	Rekomendasi         int32   `json:"rekomendasi"`
	T_efficiency_result float32 `json:"t_efficiency_result"`
}
type RatingTenderEstimateInovationResult struct {
	Id_analisa_inovasi  uint32  `json:"id_analisa_inovasi"`
	Innovation_item     string  `json:"innovation_item"`
	Innovation_value    float32 `json:"innovation_value"`
	Rekomendasi         int32   `json:"rekomendasi"`
	T_efficiency_result float32 `json:"t_efficiency_result"`
}
type RatingTenderEstimateVeResult struct {
	Id_analisa_ve       uint32  `json:"id_analisa_ve"`
	Valuee_item         string  `json:"valuee_item"`
	Valuee_value        float32 `json:"valuee_value"`
	Rekomendasi         int32   `json:"rekomendasi"`
	T_efficiency_result float32 `json:"t_efficiency_result"`
}
type RatingTenderEstimateFeResult struct {
	Id_analisa_fe       uint32  `json:"id_analisa_fe"`
	Finance_item        string  `json:"finance_item"`
	Finance_value       float32 `json:"finance_value"`
	Rekomendasi         int32   `json:"rekomendasi"`
	T_efficiency_result float32 `json:"t_efficiency_result"`
}

type RatingTenderEstimateProjectBoostResult struct {
	Select_object                     uint32  `json:"select_object"`
	Select_name                       string  `json:"select_name"`
	Value_object                      uint32  `json:"value_object"`
	Value_name                        string  `json:"value_name"`
	T_efficiency_project_boost        float32 `json:"t_efficiency_project_boost"`
	T_efficiency_project_boost_result float32 `json:"t_efficiency_project_boost_result"`
}

func FormatFloat(num float64, prc int) string {
	if num < 0 && prc == 0 {
		prc = prc + 1
	}

	var (
		zero, dot = "0", "."

		str = fmt.Sprintf("%."+strconv.Itoa(prc)+"f", num)
	)

	if str == "0" {
		str = "0.0"
	}

	return strings.TrimRight(strings.TrimRight(str, zero), dot)
}

func getDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf(err.Error())
	}

	dir_uploads := dir + "/uploads/"

	return dir_uploads
}
