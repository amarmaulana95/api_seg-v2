package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SegTenderEstimate struct {
	Id                   uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Project_name         string    `gorm:"size:200;" json:"project_name"`
	Location             string    `gorm:"size:120;" json:"location"`
	City                 string    `gorm:"size:120;" json:"city"`
	Construction_type    string    `gorm:"size:200;" json:"construction_type"`
	Building_designation string    `gorm:"size:200;" json:"building_designation"`
	Class                string    `gorm:"size:200;" json:"class"`
	Created_user         string    `gorm:"size:120;" json:"created_user"`
	Updated_user         string    `gorm:"size:120;" json:"updated_user"`
	Deleted_user         string    `gorm:"size:120;" json:"deleted_user"`
	Created_at           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted_at           time.Time `gorm:"default:NULL" json:"deleted_at"`
}

func (stestimate *SegTenderEstimate) Prepare() {
	stestimate.Id = 0
	stestimate.Project_name = html.EscapeString(strings.TrimSpace(stestimate.Project_name))
	stestimate.Location = html.EscapeString(strings.TrimSpace(stestimate.Location))
	stestimate.City = html.EscapeString(strings.TrimSpace(stestimate.City))
	stestimate.Construction_type = html.EscapeString(strings.TrimSpace(stestimate.Construction_type))
	stestimate.Building_designation = html.EscapeString(strings.TrimSpace(stestimate.Building_designation))
	stestimate.Class = html.EscapeString(strings.TrimSpace(stestimate.Class))
	stestimate.Created_user = html.EscapeString(strings.TrimSpace(stestimate.Created_user))
	stestimate.Updated_user = html.EscapeString(strings.TrimSpace(stestimate.Updated_user))
	stestimate.Created_at = time.Now()
	stestimate.Updated_at = time.Now()
}

func (stestimate *SegTenderEstimate) SaveSegTenderEstimate(db *gorm.DB) (*SegTenderEstimate, error) {

	var err error
	err = db.Debug().Create(&stestimate).Error
	if err != nil {
		return &SegTenderEstimate{}, err
	}
	return stestimate, nil
}

func (stestimate *SegTenderEstimate) FindAllSegTenderEstimates(db *gorm.DB) (*[]SegTenderEstimate, error) {
	var err error
	dstestimate := []SegTenderEstimate{}
	err = db.Debug().Model(&SegTenderEstimate{}).Limit(100).Find(&dstestimate).Error
	if err != nil {
		return &[]SegTenderEstimate{}, err
	}
	return &dstestimate, err
}

func (stestimate *SegTenderEstimate) FindSegTenderEstimateByID(db *gorm.DB, uid uint32) (*SegTenderEstimate, error) {
	var err error
	err = db.Debug().Model(SegTenderEstimate{}).Where("id = ?", uid).Take(&stestimate).Error
	if err != nil {
		return &SegTenderEstimate{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegTenderEstimate{}, errors.New("SegTenderEstimate Not Found")
	}
	return stestimate, err
}

func (stestimate *SegTenderEstimate) UpdateASegTenderEstimateAttacthment(db *gorm.DB, uid uint32) (*SegTenderEstimate, error) {
	var err error
	db = db.Debug().Model(&SegTenderEstimate{}).Where("id = ?", uid).Take(&SegTenderEstimate{}).UpdateColumns(
		map[string]interface{}{
			"id":                   stestimate.Id,
			"project_name":         stestimate.Project_name,
			"location":             stestimate.Location,
			"city":                 stestimate.City,
			"construction_type":    stestimate.Construction_type,
			"building_designation": stestimate.Building_designation,
			"class":                stestimate.Class,
			"updated_user":         stestimate.Updated_user,
			"updated_at":           stestimate.Updated_at,
		},
	)
	if db.Error != nil {
		return &SegTenderEstimate{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegTenderEstimate{}).Where("id = ?", uid).Take(&stestimate).Error
	if err != nil {
		return &SegTenderEstimate{}, err
	}
	return stestimate, nil
}

func (stestimate *SegTenderEstimate) DeleteASegTenderEstimate(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegTenderEstimate{}).Where("id = ?", uid).Take(&SegTenderEstimate{}).Delete(&SegTenderEstimate{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
