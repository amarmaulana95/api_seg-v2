package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type DashboardAnalisaTotalBulanan struct {
	Bulan uint32 `json:"bulan"`
	Tahun uint32 `json:"tahun"`
	Total uint32 `json:"total"`
}

func (DashboardAnalisaTotalBulanan) TableName() string {
	return "seg_dashboard_estimate_bulanan"
}

func (dashATBul *DashboardAnalisaTotalBulanan) Prepare() {
	dashATBul.Bulan = 0
	dashATBul.Tahun = 0
	dashATBul.Total = 0
}

func (dashATBul *DashboardAnalisaTotalBulanan) FindAllDashboardAnalisaTotalBulanans(db *gorm.DB) (*[]DashboardAnalisaTotalBulanan, error) {
	var err error
	ddashATBul := []DashboardAnalisaTotalBulanan{}
	err = db.Debug().Model(&DashboardAnalisaTotalBulanan{}).Limit(100).Order("tahun asc, bulan asc").Find(&ddashATBul).Error
	if err != nil {
		return &[]DashboardAnalisaTotalBulanan{}, err
	}
	return &ddashATBul, err
}

func (dashATBul *DashboardAnalisaTotalBulanan) FindDashboardAnalisaTotalBulananByID(db *gorm.DB, tahun uint32, bulan uint32) (*DashboardAnalisaTotalBulanan, error) {
	var err error
	err = db.Debug().Model(DashboardAnalisaTotalBulanan{}).Where("tahun = ? and bulan = ?", tahun, bulan).Take(&dashATBul).Error
	if err != nil {
		return &DashboardAnalisaTotalBulanan{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &DashboardAnalisaTotalBulanan{}, errors.New("DashboardAnalisaTotalBulanan Not Found")
	}
	return dashATBul, err
}
