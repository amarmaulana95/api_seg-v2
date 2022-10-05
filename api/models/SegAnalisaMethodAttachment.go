package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SegAnalisaMethodAttachment struct {
	Id             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Id_analisa     uint32    `json:"id_analisa"`
	File_name      string    `gorm:"size:120;" json:"name"`
	Path_file_name string    `gorm:"size:120;" json:"name"`
	Created_user   string    `gorm:"size:120;" json:"created_user"`
	Updated_user   string    `gorm:"size:120;" json:"updated_user"`
	Deleted_user   string    `gorm:"size:120;" json:"deleted_user"`
	Created_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted_at     time.Time `gorm:"default:NULL" json:"deleted_at"`
}

func (samethodattach *SegAnalisaMethodAttachment) Prepare() {
	samethodattach.Id = 0
	samethodattach.Id_analisa = 0
	samethodattach.File_name = html.EscapeString(strings.TrimSpace(samethodattach.File_name))
	samethodattach.Path_file_name = html.EscapeString(strings.TrimSpace(samethodattach.Path_file_name))
	samethodattach.Created_user = html.EscapeString(strings.TrimSpace(samethodattach.Created_user))
	samethodattach.Updated_user = html.EscapeString(strings.TrimSpace(samethodattach.Updated_user))
	samethodattach.Created_at = time.Now()
	samethodattach.Updated_at = time.Now()
}

func (samethodattach *SegAnalisaMethodAttachment) SaveSegAnalisaMethodAttachment(db *gorm.DB) (*SegAnalisaMethodAttachment, error) {

	var err error
	err = db.Debug().Create(&samethodattach).Error
	if err != nil {
		return &SegAnalisaMethodAttachment{}, err
	}
	return samethodattach, nil
}

func (samethodattach *SegAnalisaMethodAttachment) FindAllSegAnalisaMethodAttachments(db *gorm.DB) (*[]SegAnalisaMethodAttachment, error) {
	var err error
	dsamethodattach := []SegAnalisaMethodAttachment{}
	err = db.Debug().Model(&SegAnalisaMethodAttachment{}).Limit(100).Find(&dsamethodattach).Error
	if err != nil {
		return &[]SegAnalisaMethodAttachment{}, err
	}
	return &dsamethodattach, err
}

func (samethodattach *SegAnalisaMethodAttachment) FindSegAnalisaMethodAttachmentByID(db *gorm.DB, uid uint32) (*SegAnalisaMethodAttachment, error) {
	var err error
	err = db.Debug().Model(SegAnalisaMethodAttachment{}).Where("id = ?", uid).Take(&samethodattach).Error
	if err != nil {
		return &SegAnalisaMethodAttachment{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegAnalisaMethodAttachment{}, errors.New("SegAnalisaMethodAttachment Not Found")
	}
	return samethodattach, err
}

func (samethodattach *SegAnalisaMethodAttachment) UpdateSegAnalisaMethodAttachment(db *gorm.DB, uid uint32) (*SegAnalisaMethodAttachment, error) {
	var err error
	db = db.Debug().Model(&SegAnalisaMethodAttachment{}).Where("id = ?", uid).Take(&SegAnalisaMethodAttachment{}).UpdateColumns(
		map[string]interface{}{
			"id":             samethodattach.Id,
			"id_analisa":     samethodattach.Id_analisa,
			"file_name":      samethodattach.File_name,
			"path_file_name": samethodattach.Path_file_name,
			"updated_user":   samethodattach.Updated_user,
			"updated_at":     samethodattach.Updated_at,
		},
	)
	if db.Error != nil {
		return &SegAnalisaMethodAttachment{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegAnalisaMethodAttachment{}).Where("id = ?", uid).Take(&samethodattach).Error
	if err != nil {
		return &SegAnalisaMethodAttachment{}, err
	}
	return samethodattach, nil
}

func (samethodattach *SegAnalisaMethodAttachment) DeleteSegAnalisaMethodAttachment(db *gorm.DB, uid uint32, analisa_type uint32) (int64, error) {

	// db = db.Debug().Model(&SegAnalisaMethodAttachment{}).Joins("INNER JOIN seg_analisa_methods ON seg_analisa_method_attachments.id_analisa = seg_analisa_methods.id").Where("seg_analisa_method_attachments.id = ? and seg_analisa_methods.id_analisa_type = ?", uid, analisa_type).Take(&SegAnalisaMethodAttachment{}).Delete(&SegAnalisaMethodAttachment{})
	db = db.Debug().Model(&SegAnalisaMethodAttachment{}).Where("id = ?", uid).Take(&SegAnalisaMethodAttachment{}).Delete(&SegAnalisaMethodAttachment{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (samethodattach *SegAnalisaMethodAttachment) DeleteSegAnalisaMethodAttachmentByAnalisa(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegAnalisaMethodAttachment{}).Where("id_analisa = ?", uid).Delete(&SegAnalisaMethodAttachment{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (samethodattach *SegAnalisaMethodAttachment) FindAllSegAnalisaMethodAttachmentsByAnalisa(db *gorm.DB, uid uint32) ([]SegAnalisaMethodAttachment, error) {
	var err error
	dsamethodattach := []SegAnalisaMethodAttachment{}
	err = db.Debug().Model(&SegAnalisaMethodAttachment{}).Where("id_analisa = ?", uid).Find(&dsamethodattach).Error
	if err != nil {
		return []SegAnalisaMethodAttachment{}, err
	}
	return dsamethodattach, err
}
