package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SegAnalisaMethodException struct {
	Id                   uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Id_analisa           uint32    `json:"id_analisa"`
	Analisa_type         uint32    `json:"analisa_type"`
	Id_analisa_exception uint32    `json:"id_analisa_exception"`
	Label_type           string    `json:"label_type"`
	Label_exception      string    `json:"label_exception"`
	Created_user         string    `gorm:"size:120;" json:"created_user"`
	Updated_user         string    `gorm:"size:120;" json:"updated_user"`
	Deleted_user         string    `gorm:"size:120;" json:"deleted_user"`
	Created_at           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted_at           time.Time `gorm:"default:NULL" json:"deleted_at"`
}

func (samethodexcp *SegAnalisaMethodException) Prepare() {
	samethodexcp.Id = 0
	samethodexcp.Id_analisa = 0
	samethodexcp.Analisa_type = 0
	samethodexcp.Id_analisa_exception = 0
	samethodexcp.Label_type = html.EscapeString(strings.TrimSpace(samethodexcp.Label_type))
	samethodexcp.Label_exception = html.EscapeString(strings.TrimSpace(samethodexcp.Label_exception))
	samethodexcp.Created_user = html.EscapeString(strings.TrimSpace(samethodexcp.Created_user))
	samethodexcp.Updated_user = html.EscapeString(strings.TrimSpace(samethodexcp.Updated_user))
	samethodexcp.Created_at = time.Now()
	samethodexcp.Updated_at = time.Now()
}

func (samethodexcp *SegAnalisaMethodException) SaveSegAnalisaMethodException(db *gorm.DB) (*SegAnalisaMethodException, error) {

	var err error
	err = db.Debug().Create(&samethodexcp).Error
	if err != nil {
		return &SegAnalisaMethodException{}, err
	}
	return samethodexcp, nil
}

func (samethodexcp *SegAnalisaMethodException) FindAllSegAnalisaMethodExceptions(db *gorm.DB) (*[]SegAnalisaMethodException, error) {
	var err error
	dsamethodexcp := []SegAnalisaMethodException{}
	err = db.Debug().Model(&SegAnalisaMethodException{}).Limit(100).Find(&dsamethodexcp).Error
	if err != nil {
		return &[]SegAnalisaMethodException{}, err
	}
	return &dsamethodexcp, err
}

func (samethodexcp *SegAnalisaMethodException) FindSegAnalisaMethodExceptionByID(db *gorm.DB, uid uint32) (*SegAnalisaMethodException, error) {
	var err error
	err = db.Debug().Model(SegAnalisaMethodException{}).Where("id = ?", uid).Take(&samethodexcp).Error
	if err != nil {
		return &SegAnalisaMethodException{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegAnalisaMethodException{}, errors.New("SegAnalisaMethodException Not Found")
	}
	return samethodexcp, err
}

func (samethodexcp *SegAnalisaMethodException) UpdateSegAnalisaMethodException(db *gorm.DB, uid uint32) (*SegAnalisaMethodException, error) {
	var err error
	db = db.Debug().Model(&SegAnalisaMethodException{}).Where("id = ?", uid).Take(&SegAnalisaMethodException{}).UpdateColumns(
		map[string]interface{}{
			"id":                   samethodexcp.Id,
			"id_analisa":           samethodexcp.Id_analisa,
			"analisa_type":         samethodexcp.Analisa_type,
			"id_analisa_exception": samethodexcp.Id_analisa_exception,
			"label_type":           samethodexcp.Label_type,
			"label_exception":      samethodexcp.Label_exception,
			"updated_user":         samethodexcp.Updated_user,
			"updated_at":           samethodexcp.Updated_at,
		},
	)
	if db.Error != nil {
		return &SegAnalisaMethodException{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegAnalisaMethodException{}).Where("id = ?", uid).Take(&samethodexcp).Error
	if err != nil {
		return &SegAnalisaMethodException{}, err
	}
	return samethodexcp, nil
}

func (samethodattach *SegAnalisaMethodException) DeleteSegAnalisaMethodException(db *gorm.DB, uid uint32, analisa_type uint32) (int64, error) {

	// db = db.Debug().Model(&SegAnalisaMethodException{}).Joins("LEFT JOIN seg_analisa_methods ON seg_analisa_method_attachments.id_analisa = seg_analisa_methods.id").Where("seg_analisa_method_attachments.id = ? and seg_analisa_methods.id_analisa_type = ?", uid, analisa_type).Take(&SegAnalisaMethodException{}).Delete(&SegAnalisaMethodException{})
	db = db.Debug().Model(&SegAnalisaMethodException{}).Where("id = ?", uid).Take(&SegAnalisaMethodException{}).Delete(&SegAnalisaMethodException{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (samethodexcp *SegAnalisaMethodException) DeleteSegAnalisaMethodExceptionByIdAnalisa(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegAnalisaMethodException{}).Where("id_analisa = ?", uid).Delete(&SegAnalisaMethodException{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (samethodexcp *SegAnalisaMethodException) FindAllSegAnalisaMethodExceptionsByAnalisa(db *gorm.DB, uid uint32) ([]SegAnalisaMethodException, error) {
	var err error
	dsamethodexcp := []SegAnalisaMethodException{}
	err = db.Debug().Model(&SegAnalisaMethodException{}).Where("id_analisa = ?", uid).Find(&dsamethodexcp).Error
	if err != nil {
		return []SegAnalisaMethodException{}, err
	}
	return dsamethodexcp, err
}
