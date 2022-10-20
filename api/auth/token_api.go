package auth

import (
	"net/http"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserToken struct {
	User_name 		string    `json:"user_name"`
	User_nip 		string    `json:"user_nip"`
	Remember_token 	string    `json:"remember_token"`
}

func (UserToken) TableName() string {
    return "tbl_dm_user"
}

func CekToken(db *gorm.DB, r *http.Request) error {

	var err error

	err = r.ParseForm()
    if err != nil {
        fmt.Println("error parsing form", err)
    }

    /*token := r.Header.Get("token")

    utok := UserToken{}

	err = db.Debug().Model(UserToken{}).Where("remember_token = ?", token).Take(&utok).Error

	if gorm.IsRecordNotFoundError(err) {
		return err
	}*/

	return nil
}