package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Lokasi struct {
	Prop_id   uint32 `gorm:"primary_key;auto_increment" json:"id_provinsi"`
	Prop_nama string `json:"provinsi"`
}

func (Lokasi) TableName() string {
	return "tbl_provinsi"
}

func (lok *Lokasi) Prepare() {
	lok.Prop_id = 0
	lok.Prop_nama = html.EscapeString(strings.TrimSpace(lok.Prop_nama))
}

func (lok *Lokasi) FindAllLokasis(db *gorm.DB) (*[]Lokasi, error) {
	var err error
	dlok := []Lokasi{}
	err = db.Debug().Model(&Lokasi{}).Limit(100).Order("prop_id asc").Find(&dlok).Error
	if err != nil {
		return &[]Lokasi{}, err
	}
	return &dlok, err
}

func (lok *Lokasi) FindLokasiByID(db *gorm.DB, uid uint32) (*Lokasi, error) {
	var err error
	err = db.Debug().Model(Lokasi{}).Where("prop_id = ?", uid).Take(&lok).Error
	if err != nil {
		return &Lokasi{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Lokasi{}, errors.New("Lokasi Not Found")
	}
	return lok, err
}
