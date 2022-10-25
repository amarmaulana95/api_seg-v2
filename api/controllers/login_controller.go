package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/amarmaulana95/api_seg-v2/api/auth"
	"github.com/amarmaulana95/api_seg-v2/api/models"
	"github.com/amarmaulana95/api_seg-v2/api/responses"
	"github.com/amarmaulana95/api_seg-v2/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	var response Token
	response.Username = user.Username
	response.Email = user.Email
	response.Meta.Tokens = token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	fmt.Println(user.Username)
	return auth.CreateToken(user.ID, user.Username, user.Email)
}

// func (server *Server) ValidasiToken(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	r.ParseForm()
// 	token := r.Header.Get("token")
// 	utok := models.UserToken{}
// 	utokGotten, err := utok.FindUserTokenByID(server.DB, token)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	token_result, err := auth.CreateToken(utokGotten.User_id)
// 	tokres := TokenResult{}
// 	tokres.Token_seg = token_result
// 	responses.JSON(w, http.StatusOK, tokres)
// }
