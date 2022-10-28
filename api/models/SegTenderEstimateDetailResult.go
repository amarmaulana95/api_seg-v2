package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SegTenderEstimateDetailResult struct {
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

func (stestimatedetres *SegTenderEstimateDetailResult) Prepare() {
	stestimatedetres.Id = 0
	stestimatedetres.Id_tender_estimate = 0
	stestimatedetres.Id_barang = html.EscapeString(strings.TrimSpace(stestimatedetres.Id_barang))
	stestimatedetres.Price = 0
	stestimatedetres.Method = 0
	stestimatedetres.Innovation = 0
	stestimatedetres.Value_enginering = 0
	stestimatedetres.Finance_enginering = 0
	stestimatedetres.Created_user = html.EscapeString(strings.TrimSpace(stestimatedetres.Created_user))
	stestimatedetres.Updated_user = html.EscapeString(strings.TrimSpace(stestimatedetres.Updated_user))
	stestimatedetres.Created_at = time.Now()
	stestimatedetres.Updated_at = time.Now()
}

func (stestimatedetres *SegTenderEstimateDetailResult) SaveSegTenderEstimateDetailResult(db *gorm.DB) (*SegTenderEstimateDetailResult, error) {

	var err error
	err = db.Debug().Create(&stestimatedetres).Error
	if err != nil {
		return &SegTenderEstimateDetailResult{}, err
	}
	return stestimatedetres, nil
}

func (stestimatedetres *SegTenderEstimateDetailResult) FindAllSegTenderEstimateDetailResults(db *gorm.DB) (*[]SegTenderEstimateDetailResult, error) {
	var err error
	dstestimatedetres := []SegTenderEstimateDetailResult{}
	err = db.Debug().Model(&SegTenderEstimateDetailResult{}).Limit(100).Find(&dstestimatedetres).Error
	if err != nil {
		return &[]SegTenderEstimateDetailResult{}, err
	}
	return &dstestimatedetres, err
}

func (stestimatedetres *SegTenderEstimateDetailResult) FindSegTenderEstimateDetailResultByID(db *gorm.DB, uid uint32) (*SegTenderEstimateDetailResult, error) {
	var err error
	err = db.Debug().Model(SegTenderEstimateDetailResult{}).Where("id = ?", uid).Take(&stestimatedetres).Error
	if err != nil {
		return &SegTenderEstimateDetailResult{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegTenderEstimateDetailResult{}, errors.New("SegTenderEstimateDetailResult Not Found")
	}
	return stestimatedetres, err
}

func (stestimatedetres *SegTenderEstimateDetailResult) UpdateASegTenderEstimateDetailResultAttacthment(db *gorm.DB, uid uint32) (*SegTenderEstimateDetailResult, error) {
	var err error
	db = db.Debug().Model(&SegTenderEstimateDetailResult{}).Where("id = ?", uid).Take(&SegTenderEstimateDetailResult{}).UpdateColumns(
		map[string]interface{}{
			"id":                 stestimatedetres.Id,
			"id_tender_estimate": stestimatedetres.Id_tender_estimate,
			"id_barang":          stestimatedetres.Id_barang,
			"price":              stestimatedetres.Price,
			"method":             stestimatedetres.Method,
			"innovation":         stestimatedetres.Innovation,
			"value_enginering":   stestimatedetres.Value_enginering,
			"finance_enginering": stestimatedetres.Finance_enginering,
			"updated_user":       stestimatedetres.Updated_user,
			"updated_at":         stestimatedetres.Updated_at,
		},
	)
	if db.Error != nil {
		return &SegTenderEstimateDetailResult{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegTenderEstimateDetailResult{}).Where("id = ?", uid).Take(&stestimatedetres).Error
	if err != nil {
		return &SegTenderEstimateDetailResult{}, err
	}
	return stestimatedetres, nil
}

func (stestimatedetres *SegTenderEstimateDetailResult) DeleteASegTenderEstimateDetailResult(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegTenderEstimateDetailResult{}).Where("id = ?", uid).Take(&SegTenderEstimateDetailResult{}).Delete(&SegTenderEstimateDetailResult{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (stestimatedetres *SegTenderEstimateDetailResult) FindAllSegTenderEstimateDetailResultsByIdTender(db *gorm.DB, uid uint32) ([]SegTenderEstimateDetailResult, error) {
	var err error
	dstestimatedetres := []SegTenderEstimateDetailResult{}
	err = db.Debug().Model(&SegTenderEstimateDetailResult{}).Select(" seg_tender_estimate_detail_results.id_barang, master_pekerjaan_seg.tahap_nama_kendali as nama_barang, master_pekerjaan_seg.harga_rp as price, method, coalesce(a.name,'') as method_name, coalesce(e.eficiency,0) as method_koefisien, innovation, coalesce(b.name,'') as innovation_name, coalesce(f.eficiency,0) as innovation_koefisien, value_enginering, coalesce(c.name,'') as value_enginering_name, coalesce(g.eficiency,0) as value_enginering_koefisien, finance_enginering, coalesce(d.name,'') as finance_enginering_name, coalesce(h.eficiency,0) as finance_enginering_koefisien ").Joins(" left join master_pekerjaan_seg on (master_pekerjaan_seg.id_proyek::text || master_pekerjaan_seg.zona::text || '-' || master_pekerjaan_seg.tahap_kode_kendali) = seg_tender_estimate_detail_results.id_barang left join seg_analisa_methods a on seg_tender_estimate_detail_results.method = a.id left join seg_analisa_methods b on seg_tender_estimate_detail_results.innovation = b.id left join seg_analisa_methods c on seg_tender_estimate_detail_results.value_enginering = c.id left join seg_analisa_methods d on seg_tender_estimate_detail_results.finance_enginering = d.id left join seg_analisa_method_details e on e.id_analisa = a.id and e.id_barang = seg_tender_estimate_detail_results.id_barang  left join seg_analisa_method_details f on f.id_analisa = b.id and f.id_barang = seg_tender_estimate_detail_results.id_barang  left join seg_analisa_method_details g on g.id_analisa = c.id and g.id_barang = seg_tender_estimate_detail_results.id_barang  left join seg_analisa_method_details h on h.id_analisa = d.id and h.id_barang = seg_tender_estimate_detail_results.id_barang  ").Where("id_tender_estimate = ?", uid).Find(&dstestimatedetres).Error
	if err != nil {
		return []SegTenderEstimateDetailResult{}, err
	}
	return dstestimatedetres, err
}

func (stestimatedetres *SegTenderEstimateDetailResult) FindAllSegTenderEstimateDetailResultsByIdTenderIdBarang(db *gorm.DB, uid uint32, id_barang string) ([]SegTenderEstimateDetailResult, error) {
	var err error
	dstestimatedetres := []SegTenderEstimateDetailResult{}
	err = db.Debug().Model(&SegTenderEstimateDetailResult{}).Select(" seg_tender_estimate_detail_results.id_barang, master_pekerjaan_seg.tahap_nama_kendali as nama_barang, master_pekerjaan_seg.harga_rp as price, method, coalesce(a.name,'') as method_name, coalesce(e.eficiency,0) as method_koefisien, innovation, coalesce(b.name,'') as innovation_name, coalesce(f.eficiency,0) as innovation_koefisien, value_enginering, coalesce(c.name,'') as value_enginering_name, coalesce(g.eficiency,0) as value_enginering_koefisien, finance_enginering, coalesce(d.name,'') as finance_enginering_name, coalesce(h.eficiency,0) as finance_enginering_koefisien ").Joins(" left join master_pekerjaan_seg on (master_pekerjaan_seg.id_proyek::text || master_pekerjaan_seg.zona::text || '-' || master_pekerjaan_seg.tahap_kode_kendali) = seg_tender_estimate_detail_results.id_barang left join seg_analisa_methods a on seg_tender_estimate_detail_results.method = a.id left join seg_analisa_methods b on seg_tender_estimate_detail_results.innovation = b.id left join seg_analisa_methods c on seg_tender_estimate_detail_results.value_enginering = c.id left join seg_analisa_methods d on seg_tender_estimate_detail_results.finance_enginering = d.id left join seg_analisa_method_details e on e.id_analisa = a.id and e.id_barang = seg_tender_estimate_detail_results.id_barang  left join seg_analisa_method_details f on f.id_analisa = b.id and f.id_barang = seg_tender_estimate_detail_results.id_barang  left join seg_analisa_method_details g on g.id_analisa = c.id and g.id_barang = seg_tender_estimate_detail_results.id_barang  left join seg_analisa_method_details h on h.id_analisa = d.id and h.id_barang = seg_tender_estimate_detail_results.id_barang  ").Where("id_tender_estimate = ? and seg_tender_estimate_detail_results.id_barang = ?", uid, id_barang).Find(&dstestimatedetres).Error
	if err != nil {
		return []SegTenderEstimateDetailResult{}, err
	}
	return dstestimatedetres, err
}

func (stestimatedetres *SegTenderEstimateDetailResult) DeleteASegTenderEstimateDetailResultByIdEstimate(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegTenderEstimateDetailResult{}).Where("id_tender_estimate = ?", uid).Take(&SegTenderEstimateDetailResult{}).Delete(&SegTenderEstimateDetailResult{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
