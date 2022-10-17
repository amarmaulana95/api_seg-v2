package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type DashboardAnalisaEficiency struct {
	T_eficiency_result     float32 `json:"t_eficiency_result"`
	T_non_eficiency_result float32 `json:"t_non_eficiency_result"`
	Percent_eficiency      float32 `json:"percent_eficiency"`
}

func (DashboardAnalisaEficiency) TableName() string {
	return "seg_dashboard_eficiency_percent"
}

func (dashAE *DashboardAnalisaEficiency) Prepare() {
	dashAE.T_eficiency_result = 0
	dashAE.T_non_eficiency_result = 0
	dashAE.Percent_eficiency = 0
}

func (dashAE *DashboardAnalisaEficiency) FindAllDashboardAnalisaEficiencys(db *gorm.DB) (*[]DashboardAnalisaEficiency, error) {
	var err error
	ddashAE := []DashboardAnalisaEficiency{}
	err = db.Debug().Model(&DashboardAnalisaEficiency{}).Limit(100).Find(&ddashAE).Error // .Order("id asc")
	if err != nil {
		return &[]DashboardAnalisaEficiency{}, err
	}
	return &ddashAE, err
}

func (dashAE *DashboardAnalisaEficiency) FindDashboardAnalisaEficiencyByID(db *gorm.DB) (*DashboardAnalisaEficiency, error) { // , uid uint32
	var err error
	err = db.Debug().Model(DashboardAnalisaEficiency{}).Take(&dashAE).Error // .Where("id = ?", uid)
	if err != nil {
		return &DashboardAnalisaEficiency{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &DashboardAnalisaEficiency{}, errors.New("DashboardAnalisaEficiency Not Found")
	}
	return dashAE, err
}
