package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type SegAnalisaType struct {
	Id          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Urutan      uint32 `json:"urutan"`
}

func (satype *SegAnalisaType) Prepare() {
	satype.Id = 0
	satype.Name = html.EscapeString(strings.TrimSpace(satype.Name))
	satype.Description = html.EscapeString(strings.TrimSpace(satype.Description))
	satype.Urutan = 0
}

func (satype *SegAnalisaType) SaveSegAnalisaType(db *gorm.DB) (*SegAnalisaType, error) {

	var err error
	err = db.Debug().Create(&satype).Error
	if err != nil {
		return &SegAnalisaType{}, err
	}
	return satype, nil
}

func (satype *SegAnalisaType) FindAllSegAnalisaTypes(db *gorm.DB) (*[]SegAnalisaType, error) {
	var err error
	dsatype := []SegAnalisaType{}
	err = db.Debug().Model(&SegAnalisaType{}).Limit(100).Find(&dsatype).Error
	if err != nil {
		return &[]SegAnalisaType{}, err
	}
	return &dsatype, err
}

func (satype *SegAnalisaType) FindSegAnalisaTypeByID(db *gorm.DB, uid uint32) (*SegAnalisaType, error) {
	var err error
	err = db.Debug().Model(SegAnalisaType{}).Where("id = ?", uid).Take(&satype).Error
	if err != nil {
		return &SegAnalisaType{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegAnalisaType{}, errors.New("SegAnalisaType Not Found")
	}
	return satype, err
}

func (satype *SegAnalisaType) UpdateASegAnalisaType(db *gorm.DB, uid uint32) (*SegAnalisaType, error) {
	var err error
	db = db.Debug().Model(&SegAnalisaType{}).Where("id = ?", uid).Take(&SegAnalisaType{}).UpdateColumns(
		map[string]interface{}{
			"name":        satype.Name,
			"description": satype.Description,
			"urutan":      satype.Urutan,
		},
	)
	if db.Error != nil {
		return &SegAnalisaType{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegAnalisaType{}).Where("id = ?", uid).Take(&satype).Error
	if err != nil {
		return &SegAnalisaType{}, err
	}
	return satype, nil
}

func (satype *SegAnalisaType) DeleteASegAnalisaType(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegAnalisaType{}).Where("id = ?", uid).Take(&SegAnalisaType{}).Delete(&SegAnalisaType{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
