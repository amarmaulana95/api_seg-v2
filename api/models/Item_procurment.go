package models

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type ItemProcurment struct {
	Id           string  `json:"id"`
	Id_barang    string  `json:"id_barang"`
	Barang       string  `json:"barang"`
	Harga_satuan float32 `json:"price"`
	Nama_proyek  string  `json:"nama_proyek"`
}

func (ItemProcurment) TableName() string {
	return "master_pekerjaan_seg"
}

func (lok *ItemProcurment) Prepare() {
	lok.Id = html.EscapeString(strings.TrimSpace(lok.Id))
	lok.Id_barang = html.EscapeString(strings.TrimSpace(lok.Id_barang))
	lok.Barang = html.EscapeString(strings.TrimSpace(lok.Barang))
	lok.Harga_satuan = 0
	lok.Nama_proyek = html.EscapeString(strings.TrimSpace(lok.Barang))
}

func (lok *ItemProcurment) FindAllItemProcurments(db *gorm.DB, search string, id_provinsi string) (*[]ItemProcurment, error) {
	search_str := "%" + search + "%"

	id_provinsi_data := strings.Split(id_provinsi, "-")

	fmt.Println(id_provinsi_data)

	var err error
	dlok := []ItemProcurment{}

	if id_provinsi == "" {
		err = db.Debug().Model(&ItemProcurment{}).Select("id_proyek::text || zona::text || '-' || tahap_kode_kendali as id, id_proyek::text || zona::text || '-' || tahap_kode_kendali as id_barang, tahap_nama_kendali || ' (' || coalesce(nickname,'') || ')' as barang, coalesce(nickname,'') as nama_proyek,harga_rp as harga_satuan").Joins("LEFT JOIN tbl_spro_proyek on master_pekerjaan_seg.id_proyek = tbl_spro_proyek.proyek_id ").Where("lower(tahap_nama_kendali || ' (' || coalesce(nickname,'') || ')') LIKE lower(?)", search_str).Limit(80).Order("tahap_nama_kendali || ' (' || coalesce(nickname,'') || ')' asc").Find(&dlok).Error //
	} else {
		err = db.Debug().Model(&ItemProcurment{}).Select("id_proyek::text || zona::text || '-' || tahap_kode_kendali as id, id_proyek::text || zona::text || '-' || tahap_kode_kendali as id_barang, tahap_nama_kendali || ' (' || coalesce(nickname,'') || ')' as barang, coalesce(nickname,'') as nama_proyek,harga_rp as harga_satuan").Joins("LEFT JOIN tbl_spro_proyek on master_pekerjaan_seg.id_proyek = tbl_spro_proyek.proyek_id ").Where("lower(tahap_nama_kendali || ' (' || coalesce(nickname,'') || ')') LIKE lower(?) and location in (?)", search_str, id_provinsi_data).Limit(80).Order("tahap_nama_kendali || ' (' || coalesce(nickname,'') || ')' asc").Find(&dlok).Error //
	}

	if err != nil {
		return &[]ItemProcurment{}, err
	}
	return &dlok, err
}

func (lok *ItemProcurment) FindItemProcurmentByID(db *gorm.DB, uid uint32) (*ItemProcurment, error) {
	var err error
	err = db.Debug().Model(ItemProcurment{}).Select("id_proyek::text || zona::text || '-' || tahap_kode_kendali as id, id_proyek::text || zona::text || '-' || tahap_kode_kendali as id_barang, tahap_nama_kendali || ' (' || coalesce(nickname,'') || ')' as barang, coalesce(nickname,'') as nama_proyek,harga_rp as harga_satuan").Joins("LEFT JOIN tbl_spro_proyek on master_pekerjaan_seg.id_proyek = tbl_spro_proyek.proyek_id ").Where("id = ?", uid).Take(&lok).Error
	if err != nil {
		return &ItemProcurment{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &ItemProcurment{}, errors.New("ItemProcurment Not Found")
	}
	return lok, err
}
