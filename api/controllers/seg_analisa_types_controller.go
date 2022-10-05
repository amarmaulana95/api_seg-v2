package controllers

import (
	"net/http"

	"github.com/amarmaulana95/api_seg-v2/api/models"
	"github.com/amarmaulana95/api_seg-v2/api/responses"
	/*"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"*/)

func (server *Server) GetSegAnalisaTypes(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	setype := models.SegAnalisaType{}

	dsetype, err := setype.FindAllSegAnalisaTypes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dsetype)
}
