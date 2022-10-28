package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SegTenderEstimateProyekBoost struct {
	Id                 uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Id_tender_estimate uint32    `json:"id_tender_estimate"`
	Analisa_type       uint32    `json:"analisa_type"`
	Id_analisa         uint32    `json:"id_analisa"`
	Select_name        string    `gorm:"size:120;" json:"select_name"`
	Value_name         string    `gorm:"size:120;" json:"value_name"`
	Created_user       string    `gorm:"size:120;" json:"created_user"`
	Updated_user       string    `gorm:"size:120;" json:"updated_user"`
	Deleted_user       string    `gorm:"size:120;" json:"deleted_user"`
	Created_at         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted_at         time.Time `gorm:"default:NULL" json:"deleted_at"`
}

func (stestimatepboost *SegTenderEstimateProyekBoost) Prepare() {
	stestimatepboost.Id = 0
	stestimatepboost.Id_tender_estimate = 0
	stestimatepboost.Analisa_type = 0
	stestimatepboost.Id_analisa = 0
	stestimatepboost.Select_name = html.EscapeString(strings.TrimSpace(stestimatepboost.Select_name))
	stestimatepboost.Value_name = html.EscapeString(strings.TrimSpace(stestimatepboost.Value_name))
	stestimatepboost.Created_user = html.EscapeString(strings.TrimSpace(stestimatepboost.Created_user))
	stestimatepboost.Updated_user = html.EscapeString(strings.TrimSpace(stestimatepboost.Updated_user))
	stestimatepboost.Created_at = time.Now()
	stestimatepboost.Updated_at = time.Now()
}

func (stestimatepboost *SegTenderEstimateProyekBoost) SaveSegTenderEstimateProyekBoost(db *gorm.DB) (*SegTenderEstimateProyekBoost, error) {

	var err error
	err = db.Debug().Create(&stestimatepboost).Error
	if err != nil {
		return &SegTenderEstimateProyekBoost{}, err
	}
	return stestimatepboost, nil
}

func (stestimatepboost *SegTenderEstimateProyekBoost) FindAllSegTenderEstimateProyekBoosts(db *gorm.DB) (*[]SegTenderEstimateProyekBoost, error) {
	var err error
	dstestimatepboost := []SegTenderEstimateProyekBoost{}
	err = db.Debug().Model(&SegTenderEstimateProyekBoost{}).Limit(100).Find(&dstestimatepboost).Error
	if err != nil {
		return &[]SegTenderEstimateProyekBoost{}, err
	}
	return &dstestimatepboost, err
}

func (stestimatepboost *SegTenderEstimateProyekBoost) FindSegTenderEstimateProyekBoostByID(db *gorm.DB, uid uint32) (*SegTenderEstimateProyekBoost, error) {
	var err error
	err = db.Debug().Model(SegTenderEstimateProyekBoost{}).Where("id = ?", uid).Take(&stestimatepboost).Error
	if err != nil {
		return &SegTenderEstimateProyekBoost{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegTenderEstimateProyekBoost{}, errors.New("SegTenderEstimateProyekBoost Not Found")
	}
	return stestimatepboost, err
}

func (stestimatepboost *SegTenderEstimateProyekBoost) UpdateASegTenderEstimateProyekBoostAttacthment(db *gorm.DB, uid uint32) (*SegTenderEstimateProyekBoost, error) {
	var err error
	db = db.Debug().Model(&SegTenderEstimateProyekBoost{}).Where("id = ?", uid).Take(&SegTenderEstimateProyekBoost{}).UpdateColumns(
		map[string]interface{}{
			"id":                 stestimatepboost.Id,
			"id_tender_estimate": stestimatepboost.Id_tender_estimate,
			"analisa_type":       stestimatepboost.Analisa_type,
			"id_analisa":         stestimatepboost.Id_analisa,
			"select_name":        stestimatepboost.Select_name,
			"value_name":         stestimatepboost.Value_name,
			"updated_user":       stestimatepboost.Updated_user,
			"updated_at":         stestimatepboost.Updated_at,
		},
	)
	if db.Error != nil {
		return &SegTenderEstimateProyekBoost{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegTenderEstimateProyekBoost{}).Where("id = ?", uid).Take(&stestimatepboost).Error
	if err != nil {
		return &SegTenderEstimateProyekBoost{}, err
	}
	return stestimatepboost, nil
}

func (stestimatepboost *SegTenderEstimateProyekBoost) DeleteASegTenderEstimateProyekBoost(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegTenderEstimateProyekBoost{}).Where("id = ?", uid).Take(&SegTenderEstimateProyekBoost{}).Delete(&SegTenderEstimateProyekBoost{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (stestimatepboost *SegTenderEstimateProyekBoost) FindAllSegTenderEstimateProyekBoostsByIdTender(db *gorm.DB, uid uint32) ([]SegTenderEstimateProyekBoost, error) {
	var err error
	dstestimatepboost := []SegTenderEstimateProyekBoost{}
	err = db.Debug().Model(&SegTenderEstimateProyekBoost{}).Select(" seg_tender_estimate_proyek_boosts.id, seg_tender_estimate_proyek_boosts.id_tender_estimate, seg_tender_estimate_proyek_boosts.analisa_type, seg_tender_estimate_proyek_boosts.id_analisa, a.name as select_name, b.name as value_name ").Joins("left join seg_analisa_types a on seg_tender_estimate_proyek_boosts.analisa_type = a.id left join seg_analisa_methods b on seg_tender_estimate_proyek_boosts.id_analisa = b.id").Where("id_tender_estimate = ?", uid).Find(&dstestimatepboost).Error
	if err != nil {
		return []SegTenderEstimateProyekBoost{}, err
	}
	return dstestimatepboost, err
}

func (stestimatepboost *SegTenderEstimateProyekBoost) DeleteASegTenderEstimateProyekBoostByIdEstimate(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegTenderEstimateProyekBoost{}).Where("id_tender_estimate = ?", uid).Take(&SegTenderEstimateProyekBoost{}).Delete(&SegTenderEstimateProyekBoost{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
