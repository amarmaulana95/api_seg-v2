package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SegAnalisaMethodDetail struct {
	Id             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Id_analisa     uint32    `json:"id_analisa"`
	Id_barang      string    `json:"id_barang"`
	Eficiency      float32   `json:"analisa_exception_eficiency"`
	Barang         string    `gorm:"size:220;" json:"analisa_exception_label"`
	Eficiency_type uint32    `json:"eficiency_type"`
	Price          float32   `json:"price"`
	Created_user   string    `gorm:"size:120;" json:"created_user"`
	Updated_user   string    `gorm:"size:120;" json:"updated_user"`
	Deleted_user   string    `gorm:"size:120;" json:"deleted_user"`
	Created_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted_at     time.Time `gorm:"default:NULL" json:"deleted_at"`
}

func (samethoddet *SegAnalisaMethodDetail) Prepare() {
	samethoddet.Id = 0
	samethoddet.Id_analisa = 0
	samethoddet.Id_barang = html.EscapeString(strings.TrimSpace(samethoddet.Id_barang))
	samethoddet.Eficiency = 0
	samethoddet.Barang = html.EscapeString(strings.TrimSpace(samethoddet.Barang))
	samethoddet.Eficiency_type = 0
	samethoddet.Price = 0
	samethoddet.Created_user = html.EscapeString(strings.TrimSpace(samethoddet.Created_user))
	samethoddet.Updated_user = html.EscapeString(strings.TrimSpace(samethoddet.Updated_user))
	samethoddet.Created_at = time.Now()
	samethoddet.Updated_at = time.Now()
}

func (samethoddet *SegAnalisaMethodDetail) SaveSegAnalisaMethodDetail(db *gorm.DB) (*SegAnalisaMethodDetail, error) {

	var err error
	err = db.Debug().Create(&samethoddet).Error
	if err != nil {
		return &SegAnalisaMethodDetail{}, err
	}
	return samethoddet, nil
}

func (samethoddet *SegAnalisaMethodDetail) FindAllSegAnalisaMethodDetails(db *gorm.DB) (*[]SegAnalisaMethodDetail, error) {
	var err error
	dsamethoddet := []SegAnalisaMethodDetail{}
	err = db.Debug().Model(&SegAnalisaMethodDetail{}).Limit(100).Find(&dsamethoddet).Error
	if err != nil {
		return &[]SegAnalisaMethodDetail{}, err
	}
	return &dsamethoddet, err
}

func (samethoddet *SegAnalisaMethodDetail) FindSegAnalisaMethodDetailByID(db *gorm.DB, uid uint32) (*SegAnalisaMethodDetail, error) {
	var err error
	err = db.Debug().Model(SegAnalisaMethodDetail{}).Where("id = ?", uid).Take(&samethoddet).Error
	if err != nil {
		return &SegAnalisaMethodDetail{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegAnalisaMethodDetail{}, errors.New("SegAnalisaMethodDetail Not Found")
	}
	return samethoddet, err
}

func (samethoddet *SegAnalisaMethodDetail) UpdateSegAnalisaMethodDetail(db *gorm.DB, uid uint32) (*SegAnalisaMethodDetail, error) {
	var err error
	db = db.Debug().Model(&SegAnalisaMethodDetail{}).Where("id = ?", uid).Take(&SegAnalisaMethodDetail{}).UpdateColumns(
		map[string]interface{}{
			"id":             samethoddet.Id,
			"id_analisa":     samethoddet.Id_analisa,
			"id_barang":      samethoddet.Id_barang,
			"eficiency":      samethoddet.Eficiency,
			"barang":         samethoddet.Barang,
			"eficiency_type": samethoddet.Eficiency_type,
			"price":          samethoddet.Price,
			"updated_user":   samethoddet.Updated_user,
			"updated_at":     samethoddet.Updated_at,
		},
	)
	if db.Error != nil {
		return &SegAnalisaMethodDetail{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegAnalisaMethodDetail{}).Where("id = ?", uid).Take(&samethoddet).Error
	if err != nil {
		return &SegAnalisaMethodDetail{}, err
	}
	return samethoddet, nil
}

func (samethodattach *SegAnalisaMethodDetail) DeleteSegAnalisaMethodDetail(db *gorm.DB, uid uint32, analisa_type uint32) (int64, error) {
	db = db.Debug().Model(&SegAnalisaMethodDetail{}).Where("id = ?", uid).Take(&SegAnalisaMethodDetail{}).Delete(&SegAnalisaMethodDetail{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (samethoddet *SegAnalisaMethodDetail) DeleteSegAnalisaMethodDetailByIdAnalisa(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegAnalisaMethodDetail{}).Where("id_analisa = ?", uid).Delete(&SegAnalisaMethodDetail{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (samethoddet *SegAnalisaMethodDetail) FindAllSegAnalisaMethodDetailByIdAnalisa(db *gorm.DB, uid uint32) ([]SegAnalisaMethodDetail, error) {
	var err error
	dsamethoddet := []SegAnalisaMethodDetail{}
	// err = db.Debug().Model(&SegAnalisaMethodDetail{}).Select("seg_analisa_method_details.id, seg_analisa_method_details.id_analisa, seg_analisa_method_details.id_barang, seg_analisa_method_details.eficiency,  seg_analisa_method_details.eficiency_type, coalesce(master_pekerjaan_seg.harga_rp,coalesce(seg_analisa_method_details.price,0)) as price, (master_pekerjaan_seg.tahap_nama_kendali || ' (' || coalesce(tbl_spro_proyek.nickname,'') || ')') as barang").Joins("LEFT JOIN master_pekerjaan_seg ON case when seg_analisa_method_details.id_barang = '0' then '00-' || seg_analisa_method_details.id::text else seg_analisa_method_details.id_barang end = (master_pekerjaan_seg.id_proyek::text || master_pekerjaan_seg.zona::text || '-' || master_pekerjaan_seg.tahap_kode_kendali) left join tbl_spro_proyek on tbl_spro_proyek.proyek_id = master_pekerjaan_seg.id_proyek").Where("id_analisa = ?", uid).Find(&dsamethoddet).Error
	// if err != nil {
	// 	return []SegAnalisaMethodDetail{}, err
	// }

	err = db.Debug().Model(&SegAnalisaMethodDetail{}).Select("seg_analisa_method_details.id, seg_analisa_method_details.id_analisa,seg_analisa_method_details.barang, seg_analisa_method_details.eficiency_type, seg_analisa_method_details.id_barang, seg_analisa_method_details.eficiency, seg_analisa_method_details.price").Joins("LEFT JOIN seg_analisa_methods ON seg_analisa_methods.id = seg_analisa_method_details.id_analisa").Where("seg_analisa_methods.id = ?", uid).Find(&dsamethoddet).Error
	if err != nil {
		return []SegAnalisaMethodDetail{}, err
	}
	return dsamethoddet, err
}

func (samethoddet *SegAnalisaMethodDetail) FindSegAnalisaMethodDetailByAnalisaIdbarang(db *gorm.DB, id_analisa uint32, id_barang string) (*SegAnalisaMethodDetail, error) {
	var err error
	err = db.Debug().Model(SegAnalisaMethodDetail{}).Select("seg_analisa_method_details.id, seg_analisa_method_details.id_analisa, seg_analisa_method_details.id_barang, seg_analisa_method_details.eficiency, seg_analisa_method_details.eficiency_type, coalesce(master_pekerjaan_seg.harga_rp,coalesce(seg_analisa_method_details.price,0)) as price, (master_pekerjaan_seg.tahap_nama_kendali || ' (' || coalesce(tbl_spro_proyek.nickname,'') || ')') as barang").Joins("LEFT JOIN master_pekerjaan_seg ON case when seg_analisa_method_details.id_barang = '0' then '00-' || seg_analisa_method_details.id::text else seg_analisa_method_details.id_barang end = (master_pekerjaan_seg.id_proyek::text || master_pekerjaan_seg.zona::text || '-' || master_pekerjaan_seg.tahap_kode_kendali) left join tbl_spro_proyek on tbl_spro_proyek.proyek_id = master_pekerjaan_seg.id_proyek").Where("id_analisa = ? and id_barang = ?", id_analisa, id_barang).Take(&samethoddet).Error
	if err != nil {
		return &SegAnalisaMethodDetail{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegAnalisaMethodDetail{}, errors.New("SegAnalisaMethodDetail Not Found")
	}
	return samethoddet, err
}
