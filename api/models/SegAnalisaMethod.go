package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SegAnalisaMethod struct {
	Id                  uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Id_analisa_type     uint32    `json:"id_analisa_type"`
	Name                string    `gorm:"size:120;" json:"name"`
	Description         string    `gorm:"size:120;" json:"description"`
	Location            string    `gorm:"size:120;" json:"location"`
	Location_name       string    `gorm:"size:120;" json:"location_name"`
	Status_proyek_boost uint32    `json:"status_proyek_boost"`
	Created_user        string    `gorm:"size:120;" json:"created_user"`
	Updated_user        string    `gorm:"size:120;" json:"updated_user"`
	Deleted_user        string    `gorm:"size:120;" json:"deleted_user"`
	Created_at          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted_at          time.Time `gorm:"default:NULL" json:"deleted_at"`
}

func (samethod *SegAnalisaMethod) Prepare() {
	samethod.Id = 0
	samethod.Id_analisa_type = 0
	samethod.Name = html.EscapeString(strings.TrimSpace(samethod.Name))
	samethod.Description = html.EscapeString(strings.TrimSpace(samethod.Description))
	samethod.Location = html.EscapeString(strings.TrimSpace(samethod.Location))
	samethod.Location_name = html.EscapeString(strings.TrimSpace(samethod.Location_name))
	samethod.Status_proyek_boost = 0
	samethod.Created_user = html.EscapeString(strings.TrimSpace(samethod.Created_user))
	samethod.Updated_user = html.EscapeString(strings.TrimSpace(samethod.Updated_user))
	samethod.Created_at = time.Now()
	samethod.Updated_at = time.Now()
}

func (samethod *SegAnalisaMethod) SaveSegAnalisaMethod(db *gorm.DB) (*SegAnalisaMethod, error) {

	var err error
	err = db.Debug().Create(&samethod).Error
	if err != nil {
		return &SegAnalisaMethod{}, err
	}
	return samethod, nil
}

func (samethod *SegAnalisaMethod) FindAllSegAnalisaMethods(db *gorm.DB, search string) (*[]SegAnalisaMethod, error) {
	search_str := "%" + search + "%"

	var err error
	dsamethod := []SegAnalisaMethod{}
	err = db.Debug().Model(&SegAnalisaMethod{}).Where("lower(name) like lower(?)", search_str).Limit(100).Find(&dsamethod).Error
	if err != nil {
		return &[]SegAnalisaMethod{}, err
	}
	return &dsamethod, err
}

func (samethod *SegAnalisaMethod) FindAllSegAnalisaMethodsFull(db *gorm.DB, antipe uint32, search string, limit uint64, offset uint64) ([]SegAnalisaMethod, error) {
	search_str := "%" + search + "%"

	var err error
	dsamethod := []SegAnalisaMethod{}
	err = db.Debug().Model(&SegAnalisaMethod{}).Select("seg_analisa_methods.id, seg_analisa_methods.id_analisa_type, seg_analisa_methods.name, seg_analisa_methods.description, seg_analisa_methods.location, seg_analisa_methods.status_proyek_boost, b.location_name").Joins(" join seg_method_lokasi b on seg_analisa_methods.id = b.id").Where("id_analisa_type = ? and lower(name) like lower(?)", antipe, search_str).Limit(limit).Offset(offset).Find(&dsamethod).Error
	if err != nil {
		return []SegAnalisaMethod{}, err
	}
	return dsamethod, err
}

func (samethod *SegAnalisaMethod) FindAllSegAnalisaMethodsTotal(db *gorm.DB, antipe uint32, search string) uint64 {
	search_str := "%" + search + "%"

	// var err error
	dsamethod := []SegAnalisaMethod{}
	/*err = db.Debug().Model(&SegAnalisaMethod{}).Where("lower(name) like lower(?)", search_str).Find(&dsamethod).Error
	if err != nil {
		return &[]SegAnalisaMethod{}, err
	}*/

	result := db.Where("id_analisa_type = ? and lower(name) like lower(?)", antipe, search_str).Find(&dsamethod)

	return uint64(result.RowsAffected)
}

func (samethod *SegAnalisaMethod) FindSegAnalisaMethodByID(db *gorm.DB, uid uint32) (*SegAnalisaMethod, error) {
	var err error
	err = db.Debug().Model(SegAnalisaMethod{}).Select("seg_analisa_methods.id, seg_analisa_methods.id_analisa_type, seg_analisa_methods.name, seg_analisa_methods.description, seg_analisa_methods.location, seg_analisa_methods.status_proyek_boost, b.location_name").Joins(" join seg_method_lokasi b on seg_analisa_methods.id = b.id").Where("id = ?", uid).Take(&samethod).Error
	if err != nil {
		return &SegAnalisaMethod{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegAnalisaMethod{}, errors.New("SegAnalisaMethod Not Found")
	}
	return samethod, err
}

func (samethod *SegAnalisaMethod) FindAllSegAnalisaMethodsByType(db *gorm.DB, id_type uint32, search string, id_barang string, id_provinsi string) (*[]SegAnalisaMethod, error) { //
	search_str := "%" + search + "%"

	var err error
	dsamethod := []SegAnalisaMethod{}

	if id_provinsi == "" {
		err = db.Debug().Model(&SegAnalisaMethod{}).Select("seg_analisa_methods.id, seg_analisa_methods.id_analisa_type, seg_analisa_methods.name, seg_analisa_methods.description, seg_analisa_methods.location, seg_analisa_methods.status_proyek_boost").Joins("left join seg_analisa_method_details a on seg_analisa_methods.id = a.id_analisa left join seg_method_lokasi_split b on seg_analisa_methods.id = b.id").Group("seg_analisa_methods.id, seg_analisa_methods.id_analisa_type, seg_analisa_methods.name, seg_analisa_methods.description, seg_analisa_methods.location, seg_analisa_methods.status_proyek_boost").Where("seg_analisa_methods.id_analisa_type = ? and lower(seg_analisa_methods.name) like lower(?) and case when a.id_barang = '0' then '00-' || a.id::text else a.id_barang end = ?", id_type, search_str, id_barang).Limit(100).Find(&dsamethod).Error //
	} else {
		err = db.Debug().Model(&SegAnalisaMethod{}).Select("seg_analisa_methods.id, seg_analisa_methods.id_analisa_type, seg_analisa_methods.name, seg_analisa_methods.description, seg_analisa_methods.location, seg_analisa_methods.status_proyek_boost").Joins("left join seg_analisa_method_details a on seg_analisa_methods.id = a.id_analisa left join seg_method_lokasi_split b on seg_analisa_methods.id = b.id").Group("seg_analisa_methods.id, seg_analisa_methods.id_analisa_type, seg_analisa_methods.name, seg_analisa_methods.description, seg_analisa_methods.location, seg_analisa_methods.status_proyek_boost").Where("seg_analisa_methods.id_analisa_type = ? and lower(seg_analisa_methods.name) like lower(?) and case when a.id_barang = '0' then '00-' || a.id::text else a.id_barang end = ? and b.location = ?", id_type, search_str, id_barang, id_provinsi).Limit(100).Find(&dsamethod).Error //
	}

	if err != nil {
		return &[]SegAnalisaMethod{}, err
	}
	return &dsamethod, err
}

func (samethod *SegAnalisaMethod) FindAllSegAnalisaMethodsByPBoost(db *gorm.DB, search string, type_analisa uint32, id_provinsi string) (*[]SegAnalisaMethod, error) {
	search_str := "%" + search + "%"

	var err error
	dsamethod := []SegAnalisaMethod{}

	if id_provinsi == "" {
		err = db.Debug().Model(&SegAnalisaMethod{}).Joins("left join seg_method_lokasi_split b on seg_analisa_methods.id = b.id").Where("status_proyek_boost = 1 and lower(name) like lower(?) and id_analisa_type = ?", search_str, type_analisa).Limit(100).Find(&dsamethod).Error
	} else {
		err = db.Debug().Model(&SegAnalisaMethod{}).Joins("left join seg_method_lokasi_split b on seg_analisa_methods.id = b.id").Where("status_proyek_boost = 1 and lower(name) like lower(?) and id_analisa_type = ? and b.location = ?", search_str, type_analisa, id_provinsi).Limit(100).Find(&dsamethod).Error
	}

	if err != nil {
		return &[]SegAnalisaMethod{}, err
	}
	return &dsamethod, err
}

func (samethod *SegAnalisaMethod) UpdateSegAnalisaMethod(db *gorm.DB, uid uint32) (*SegAnalisaMethod, error) {
	var err error
	db = db.Debug().Model(&SegAnalisaMethod{}).Where("id = ?", uid).Take(&SegAnalisaMethod{}).UpdateColumns(
		map[string]interface{}{
			"id":                  samethod.Id,
			"id_analisa_type":     samethod.Id_analisa_type,
			"name":                samethod.Name,
			"description":         samethod.Description,
			"location":            samethod.Location,
			"location_name":       samethod.Location_name,
			"status_proyek_boost": samethod.Status_proyek_boost,
			"updated_user":        samethod.Updated_user,
			"updated_at":          samethod.Updated_at,
		},
	)
	if db.Error != nil {
		return &SegAnalisaMethod{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegAnalisaMethod{}).Where("id = ?", uid).Take(&samethod).Error
	if err != nil {
		return &SegAnalisaMethod{}, err
	}
	return samethod, nil
}

func (samethod *SegAnalisaMethod) DeleteSegAnalisaMethod(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegAnalisaMethod{}).Where("id = ?", uid).Delete(&SegAnalisaMethod{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (samethod *SegAnalisaMethod) FindAllSegAnalisaMethodsByException(db *gorm.DB, search string, type_analisa uint32) (*[]SegAnalisaMethod, error) {
	search_str := "%" + search + "%"

	var err error
	dsamethod := []SegAnalisaMethod{}
	err = db.Debug().Model(&SegAnalisaMethod{}).Where("lower(name) like lower(?) and id_analisa_type = ?", search_str, type_analisa).Limit(100).Find(&dsamethod).Error
	if err != nil {
		return &[]SegAnalisaMethod{}, err
	}
	return &dsamethod, err
}
