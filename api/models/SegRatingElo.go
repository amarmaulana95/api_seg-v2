package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type SegRatingElo struct {
	Id              uint32  `gorm:"primary_key;auto_increment" json:"id"`
	Id_analisa_type uint32  `json:"id_analisa_type"`
	Id_analisa      uint32  `json:"id_analisa"`
	Id_barang       string  `json:"id_barang"`
	Koefisien       float32 `json:"koefisien"`
	Rating          float32 `json:"rating"`
	Analisa_name    string  `json:"analisa_name"`
}

func (srelo *SegRatingElo) Prepare() {
	srelo.Id = 0
	srelo.Id_analisa_type = 0
	srelo.Id_analisa = 0
	srelo.Id_barang = html.EscapeString(strings.TrimSpace(srelo.Id_barang))
	srelo.Koefisien = 0
	srelo.Rating = 0
}

func (srelo *SegRatingElo) SaveSegRatingElo(db *gorm.DB) (*SegRatingElo, error) {

	var err error
	err = db.Debug().Create(&srelo).Error
	if err != nil {
		return &SegRatingElo{}, err
	}
	return srelo, nil
}

func (srelo *SegRatingElo) FindAllSegRatingElos(db *gorm.DB) (*[]SegRatingElo, error) {
	var err error
	dsrelo := []SegRatingElo{}
	err = db.Debug().Model(&SegRatingElo{}).Limit(100).Find(&dsrelo).Error
	if err != nil {
		return &[]SegRatingElo{}, err
	}
	return &dsrelo, err
}

func (srelo *SegRatingElo) FindSegRatingEloByID(db *gorm.DB, uid uint32) (*SegRatingElo, error) {
	var err error
	err = db.Debug().Model(SegRatingElo{}).Where("id = ?", uid).Take(&srelo).Error
	if err != nil {
		return &SegRatingElo{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SegRatingElo{}, errors.New("SegRatingElo Not Found")
	}
	return srelo, err
}

func (srelo *SegRatingElo) UpdateASegRatingElo(db *gorm.DB, uid uint32) (*SegRatingElo, error) {
	var err error
	db = db.Debug().Model(&SegRatingElo{}).Where("id = ?", uid).Take(&SegRatingElo{}).UpdateColumns(
		map[string]interface{}{
			"id":              srelo.Id,
			"id_analisa_type": srelo.Id_analisa_type,
			"id_analisa":      srelo.Id_analisa,
			"id_barang":       srelo.Id_barang,
			"koefisien":       srelo.Koefisien,
			"rating":          srelo.Rating,
		},
	)
	if db.Error != nil {
		return &SegRatingElo{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegRatingElo{}).Where("id = ?", uid).Take(&srelo).Error
	if err != nil {
		return &SegRatingElo{}, err
	}
	return srelo, nil
}

func (srelo *SegRatingElo) DeleteASegRatingElo(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SegRatingElo{}).Where("id = ?", uid).Take(&SegRatingElo{}).Delete(&SegRatingElo{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (srelo *SegRatingElo) FindAllSegRatingElosCalibrate(db *gorm.DB) ([]SegRatingElo, error) {
	var err error
	dsrelo := []SegRatingElo{}
	// err = db.Debug().Model(&SegRatingElo{}).Select(" 0 as id, seg_analisa_methods.id_analisa_type as id_analisa_type, seg_analisa_methods.id as id_analisa, a.id_barang as id_barang, a.eficiency as koefisien, b.rating as rating ").Joins(" left join seg_analisa_method_details a on seg_analisa_methods.id = a.id_analisa  left join seg_rating_elos b on seg_analisa_methods.id_analisa_type = b.id_analisa_type and seg_analisa_methods.id = b.id_analisa and a.id_barang = b.id_barang  ").Where("?",pwhere).Find(&dsrelo).Error

	err = db.Debug().Raw("select seg_analisa_methods.id_analisa_type as id_analisa_type, seg_analisa_methods.id as id_analisa, a.id_barang as id_barang, a.eficiency as koefisien, b.rating as rating from seg_analisa_methods left join seg_analisa_method_details a on seg_analisa_methods.id = a.id_analisa left join seg_rating_elos b on seg_analisa_methods.id_analisa_type = b.id_analisa_type and seg_analisa_methods.id = b.id_analisa and a.id_barang = b.id_barang where a.id_barang is not null and b.rating is null order by seg_analisa_methods.id_analisa_type, seg_analisa_methods.id, a.id").Find(&dsrelo).Error
	if err != nil {
		return []SegRatingElo{}, err
	}
	return dsrelo, err
}

func (srelo *SegRatingElo) FindAllSegRatingEloRecommend(db *gorm.DB, idAnalisatype uint32, idBarang string) (*SegRatingElo, error) {
	var err error
	err = db.Debug().Raw("select * from seg_best_rating_name where id_analisa_type = ? and id_barang = ?", idAnalisatype, idBarang).Take(&srelo).Error
	if err != nil {
		return &SegRatingElo{}, err
	}
	return srelo, err
}

func (srelo *SegRatingElo) FindAllSegByIdTypeAnalisaBarang(db *gorm.DB, id_analisa_type uint32, id_analisa uint32, id_barang string) (*SegRatingElo, error) {
	var err error
	err = db.Debug().Where("id_analisa_type = ? and id_analisa = ? and id_barang = ?", id_analisa_type, id_analisa, id_barang).Take(&srelo).Error
	if err != nil {
		return &SegRatingElo{}, err
	}
	return srelo, err
}

func (srelo *SegRatingElo) UpdateASegRatingEloByIdTypeAnalisaBarang(db *gorm.DB, id_analisa_type uint32, id_analisa uint32, id_barang string) (*SegRatingElo, error) {
	var err error
	db = db.Debug().Model(&SegRatingElo{}).Where("id_analisa_type = ? and id_analisa = ? and id_barang = ?", id_analisa_type, id_analisa, id_barang).Take(&SegRatingElo{}).UpdateColumns(
		map[string]interface{}{
			"rating": srelo.Rating,
		},
	)
	if db.Error != nil {
		return &SegRatingElo{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SegRatingElo{}).Where("id_analisa_type = ? and id_analisa = ? and id_barang = ?", id_analisa_type, id_analisa, id_barang).Take(&srelo).Error
	if err != nil {
		return &SegRatingElo{}, err
	}
	return srelo, nil
}
