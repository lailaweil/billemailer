package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lailaweil/billemailer/api/domain"
	"github.com/lailaweil/billemailer/api/errors"
	"github.com/lailaweil/billemailer/api/services"
	"github.com/lailaweil/billemailer/api/utils"
	"net/http"
)

type FolderController struct {
	service services.FolderService
}

func NewFolderController(service services.FolderService) FolderController {
	return FolderController{service: service}
}

//TODO: implement swagger

//CreateFolder creates a Folder
func (c FolderController) CreateFolder(w http.ResponseWriter, r *http.Request) {
	// Declare a new Folder.
	folder := &domain.Folder{}

	if err := json.NewDecoder(r.Body).Decode(folder); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error decoding body", err.Error())
		utils.WriteResponse(w, errResponse, errResponse.Status)
		return
	}

	if err := folder.Validate(); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error validating body", err.Error())
		utils.WriteResponse(w, errResponse, errResponse.Status)
		return
	}

	if folder.ID != 0 {
		errResponse := errors.NewError(http.StatusBadRequest, "error validating body", "id must be empty")
		utils.WriteResponse(w, errResponse, errResponse.Status)
		return
	}

	//CreateFolder
	result, errCreate := c.service.CreateFolder(folder)

	if errCreate != nil {
		utils.WriteResponse(w, errCreate, errCreate.Status)
		return
	}

	utils.WriteResponse(w, result, http.StatusCreated)
}

//GetFolder returns a given Folder
func (c FolderController) GetFolder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	folder, err := c.service.GetFolder(vars["id"])

	if err != nil {
		utils.WriteResponse(w, err, err.Status)
		return
	}

	utils.WriteResponse(w, folder, http.StatusOK)
}
