package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SegTenderEstimateDetail struct {
	Id                           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Id_tender_estimate           uint32    `json:"id_tender_estimate"`
	Id_barang                    string    `json:"id_barang"`
	Nama_barang                  string    `json:"nama_barang"`
	Price                        float32   `json:"price"`
	Method                       uint32    `json:"method"`
	Method_name                  string    `json:"method_name"`
	Method_koefisien             float32   `json:"method_koefisien"`
	Innovation                   uint32    `json:"innovation"`
	Innovation_name              string    `json:"innovation_name"`
	Innovation_koefisien         float32   `json:"innovation_koefisien"`
	Value_enginering             uint32    `json:"value_enginering"`
	Value_enginering_name        string    `json:"value_enginering_name"`
	Value_enginering_koefisien   float32   `json:"value_enginering_koefisien"`
	Finance_enginering           uint32    `json:"finance_enginering"`
	Finance_enginering_name      string    `json:"finance_enginering_name"`
	Finance_enginering_koefisien float32   `json:"finance_enginering_koefisien"`
	Created_user                 string    `gorm:"size:120;" json:"created_user"`
	Updated_user                 string    `gorm:"size:120;" json:"updated_user"`
	Deleted_user                 string    `gorm:"size:120;" json:"deleted_user"`
	Created_at                   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at                   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted_at                   time.Time `gorm:"default:NULL" json:"deleted_at"`
}

func (stestimatedet *SegTenderEstimateDetail) Prepare() {
	stestimatedet.Id = 0
	stestimatedet.Id_tender_estimate = 0
	stestimatedet.Id_barang = html.EscapeString(strings.TrimSpace(stestimatedet.Id_barang))
	stestimatedet.Price = 0
	stestimatedet.Method = 0
	stestimatedet.Innovation = 0
	stestimatedet.Value_enginering = 0
	stestimatedet.Finance_enginering = 0
	stestimatedet.Created_user = html.EscapeString(strings.TrimSpace(stestimatedet.Created_user))
	stestimatedet.Updated_user = html.EscapeString(strings.TrimSpace(stestimatedet.Updated_user))
	stestimatedet.Created_at = time.Now()
	stestimatedet.Updated_at = time.Now()
}

func (stestimatedet *SegTenderEstimateDetail) SaveSegTenderEstimateDetail(db *gorm.DB) (*SegTenderEstimateDetail, error) {

	var err error
	err = db.Debug().Create(&stestimatedet).Error
	if err != nil {
		return &SegTenderEstimateDetail{}, err
	}
	return stestimatedet, nil
}

func (stestimatedet *SegTenderEstimateDetail) FindAllSegTenderEstimateDetails(db *gorm.DB) (*[]SegTenderEstimateDetail, error) {
	var err error
	dstestimatedet := []SegTenderEstimateDetail{}
	err = db.Debug().Model(&SegTenderEstimateDetail{}).Limit(100).Find(&dstestimatedet).Error
	if err != nil {
		return &[]SegTenderEstimateDetail{}, err
	}
	return &dstestimatedet, err
}

func (stestimatedet *SegTenderEstimateDetail) FindSegTenderEstimateDetailByID(db *gorm.DB, uid uint32) (*SegTenderEstimateDetail, error) {
	var err error
	err = db.Debug().Model(SegTenderEstimateDetail{}).Where("id = ?", uid).Take(&stestimatedet).Error
	if err != nil {
		return &SegTenderEstimateDetail{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegTenderEstimateDetail{}, errors.New("SegTenderEstimateDetail Not Found")
	}
	return stestimatedet, err
}

func (stestimatedet *SegTenderEstimateDetail) UpdateASegTenderEstimateDetailAttacthment(db *gorm.DB, uid uint32) (*SegTenderEstimateDetail, error) {
	var err error
	db = db.Debug().Model(&SegTenderEstimateDetail{}).Where("id = ?", uid).Take(&SegTenderEstimateDetail{}).UpdateColumns(
		map[string]interface{}{
			"id":                 stestimatedet.Id,
			"id_tender_estimate": stestimatedet.Id_tender_estimate,
			"id_barang":          stestimatedet.Id_barang,
			"price":              stestimatedet.Price,
			"method":             stestimatedet.Method,
			"innovation":         stestimatedet.Innovation,
			"value_enginering":   stestimatedet.Value_enginering,
			"finance_enginering": stestimatedet.Finance_enginering,
			"updated_user":       stestimatedet.Updated_user,
			"updated_at":         stestimatedet.Updated_at,
		},
	)
	if db.Error != nil {
		return &SegTenderEstimateDetail{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegTenderEstimateDetail{}).Where("id = ?", uid).Take(&stestimatedet).Error
	if err != nil {
		return &SegTenderEstimateDetail{}, err
	}
	return stestimatedet, nil
}

func (stestimatedet *SegTenderEstimateDetail) DeleteASegTenderEstimateDetail(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegTenderEstimateDetail{}).Where("id = ?", uid).Take(&SegTenderEstimateDetail{}).Delete(&SegTenderEstimateDetail{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (stestimatedet *SegTenderEstimateDetail) FindAllSegTenderEstimateDetailsByIdTender(db *gorm.DB, uid uint32) ([]SegTenderEstimateDetail, error) {
	var err error
	dstestimatedet := []SegTenderEstimateDetail{}
	err = db.Debug().Model(&SegTenderEstimateDetail{}).Select(" seg_tender_estimate_details.id_barang, master_pekerjaan_seg.tahap_nama_kendali as nama_barang, master_pekerjaan_seg.harga_rp as price, method, coalesce(a.name,'') as method_name, coalesce(e.eficiency,0) as method_koefisien, innovation, coalesce(b.name,'') as innovation_name, coalesce(f.eficiency,0) as innovation_koefisien, value_enginering, coalesce(c.name,'') as value_enginering_name, coalesce(g.eficiency,0) as value_enginering_koefisien, finance_enginering, coalesce(d.name,'') as finance_enginering_name, coalesce(h.eficiency,0) as finance_enginering_koefisien ").Joins(" left join master_pekerjaan_seg on (master_pekerjaan_seg.id_proyek::text || master_pekerjaan_seg.zona::text || '-' || master_pekerjaan_seg.tahap_kode_kendali) = seg_tender_estimate_details.id_barang left join seg_analisa_methods a on seg_tender_estimate_details.method = a.id left join seg_analisa_methods b on seg_tender_estimate_details.innovation = b.id left join seg_analisa_methods c on seg_tender_estimate_details.value_enginering = c.id left join seg_analisa_methods d on seg_tender_estimate_details.finance_enginering = d.id left join seg_analisa_method_details e on e.id_analisa = a.id and e.id_barang = seg_tender_estimate_details.id_barang  left join seg_analisa_method_details f on f.id_analisa = b.id and f.id_barang = seg_tender_estimate_details.id_barang  left join seg_analisa_method_details g on g.id_analisa = c.id and g.id_barang = seg_tender_estimate_details.id_barang  left join seg_analisa_method_details h on h.id_analisa = d.id and h.id_barang = seg_tender_estimate_details.id_barang  ").Where("id_tender_estimate = ?", uid).Find(&dstestimatedet).Error
	if err != nil {
		return []SegTenderEstimateDetail{}, err
	}
	return dstestimatedet, err
}

func (stestimatedet *SegTenderEstimateDetail) FindAllSegTenderEstimateDetailsByIdTenderIdBarang(db *gorm.DB, uid uint32, id_barang string) ([]SegTenderEstimateDetail, error) {
	var err error
	dstestimatedet := []SegTenderEstimateDetail{}
	err = db.Debug().Model(&SegTenderEstimateDetail{}).Select(" seg_tender_estimate_details.id_barang, master_pekerjaan_seg.tahap_nama_kendali as nama_barang, master_pekerjaan_seg.harga_rp as price, method, coalesce(a.name,'') as method_name, coalesce(e.eficiency,0) as method_koefisien, innovation, coalesce(b.name,'') as innovation_name, coalesce(f.eficiency,0) as innovation_koefisien, value_enginering, coalesce(c.name,'') as value_enginering_name, coalesce(g.eficiency,0) as value_enginering_koefisien, finance_enginering, coalesce(d.name,'') as finance_enginering_name, coalesce(h.eficiency,0) as finance_enginering_koefisien ").Joins(" left join master_pekerjaan_seg on (master_pekerjaan_seg.id_proyek::text || master_pekerjaan_seg.zona::text || '-' || master_pekerjaan_seg.tahap_kode_kendali) = seg_tender_estimate_details.id_barang  left join seg_analisa_methods a on seg_tender_estimate_details.method = a.id  left join seg_analisa_methods b on seg_tender_estimate_details.innovation = b.id  left join seg_analisa_methods c on seg_tender_estimate_details.value_enginering = c.id  left join seg_analisa_methods d on seg_tender_estimate_details.finance_enginering = d.id  left join seg_analisa_method_details e on e.id_analisa = a.id and case when e.eficiency_type = 1 then '00-' || e.id else e.id_barang  end = seg_tender_estimate_details.id_barang   left join seg_analisa_method_details f on f.id_analisa = b.id and case when f.eficiency_type = 1 then '00-' || f.id else f.id_barang  end = seg_tender_estimate_details.id_barang   left join seg_analisa_method_details g on g.id_analisa = c.id and case when g.eficiency_type = 1 then '00-' || g.id else g.id_barang  end = seg_tender_estimate_details.id_barang   left join seg_analisa_method_details h on h.id_analisa = d.id and case when h.eficiency_type = 1 then '00-' || h.id else h.id_barang  end = seg_tender_estimate_details.id_barang ").Where("id_tender_estimate = ? and seg_tender_estimate_details.id_barang = ?", uid, id_barang).Find(&dstestimatedet).Error
	if err != nil {
		return []SegTenderEstimateDetail{}, err
	}
	return dstestimatedet, err
}

func (stestimatedet *SegTenderEstimateDetail) DeleteASegTenderEstimateDetailByIdEstimate(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegTenderEstimateDetail{}).Where("id_tender_estimate = ?", uid).Take(&SegTenderEstimateDetail{}).Delete(&SegTenderEstimateDetail{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
