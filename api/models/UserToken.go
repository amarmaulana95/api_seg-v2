package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type UserToken struct {
	User_id        uint32 `json:"id"`
	User_name      string `json:"user_name"`
	User_nip       string `json:"user_nip"`
	Remember_token string `json:"remember_token"`
}

func (UserToken) TableName() string {
	return "tbl_dm_user"
}

func (utok *UserToken) Prepare() {
	utok.User_id = 0
	utok.User_name = html.EscapeString(strings.TrimSpace(utok.User_name))
	utok.User_nip = html.EscapeString(strings.TrimSpace(utok.User_nip))
	utok.Remember_token = html.EscapeString(strings.TrimSpace(utok.Remember_token))
}

func (utok *UserToken) FindUserTokenByID(db *gorm.DB, token string) (*UserToken, error) {
	var err error
	err = db.Debug().Model(UserToken{}).Where("remember_token = ?", token).Take(&utok).Error
	if err != nil {
		return &UserToken{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserToken{}, errors.New("UserToken Not Found")
	}
	return utok, err
}
