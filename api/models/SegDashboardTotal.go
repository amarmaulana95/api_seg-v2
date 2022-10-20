package models

import (
	"errors"
	"html"
	"strings"
	"github.com/jinzhu/gorm"
)

type DashboardAnalisaTotal struct {
	Id 		uint32 `json:"id"`
	Name 	string `json:"name"`
	Total 	uint32 `json:"total"`
	Urutan 	uint32 `json:"urutan"`
}

func (DashboardAnalisaTotal) TableName() string {
    return "seg_analisa_total_by_type"
}

func (dashAT *DashboardAnalisaTotal) Prepare() {
	dashAT.Id = 0
	dashAT.Name = html.EscapeString(strings.TrimSpace(dashAT.Name))
	dashAT.Total = 0
	dashAT.Urutan = 0
}

func (dashAT *DashboardAnalisaTotal) FindAllDashboardAnalisaTotals(db *gorm.DB) (*[]DashboardAnalisaTotal, error) {
	var err error
	ddashAT := []DashboardAnalisaTotal{}
	err = db.Debug().Model(&DashboardAnalisaTotal{}).Limit(100).Order("id asc").Find(&ddashAT).Error
	if err != nil {
		return &[]DashboardAnalisaTotal{}, err
	}
	return &ddashAT, err
}

func (dashAT *DashboardAnalisaTotal) FindDashboardAnalisaTotalByID(db *gorm.DB, uid uint32) (*DashboardAnalisaTotal, error) {
	var err error
	err = db.Debug().Model(DashboardAnalisaTotal{}).Where("id = ?", uid).Take(&dashAT).Error
	if err != nil {
		return &DashboardAnalisaTotal{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &DashboardAnalisaTotal{}, errors.New("DashboardAnalisaTotal Not Found")
	}
	return dashAT, err
}
