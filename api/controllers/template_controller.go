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

type TemplateController struct {
	service services.TemplateService
}

func NewTemplateController(service services.TemplateService) TemplateController {
	return TemplateController{service: service}
}

//TODO: implement swagger

//CreateTemplate creates a Template
func (c TemplateController) CreateTemplate(w http.ResponseWriter, r *http.Request) {
	// Declare a new Template.
	template := domain.Template{}

	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error decoding body", err.Error())
		utils.WriteResponse(w, errResponse, errResponse.Status)
		return
	}

	if err := template.Validate(); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error validating body", err.Error())
		utils.WriteResponse(w, errResponse, errResponse.Status)
		return
	}

	if template.ID != 0 {
		errResponse := errors.NewError(http.StatusBadRequest, "error validating body", "id must be empty")
		utils.WriteResponse(w, errResponse, errResponse.Status)
		return
	}

	//CreateTemplate
	result, errCreate := c.service.CreateTemplate(&template)

	if errCreate != nil {
		utils.WriteResponse(w, errCreate, errCreate.Status)
		return
	}

	utils.WriteResponse(w, result, http.StatusCreated)
}

//UpdateTemplate updates a given Template
func (c TemplateController) UpdateTemplate(w http.ResponseWriter, r *http.Request) {
	template := domain.Template{}

	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error decoding body", err.Error())
		utils.WriteResponse(w, errResponse, errResponse.Status)
		return
	}

	if err := template.Validate(); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error validating body", err.Error())
		utils.WriteResponse(w, errResponse, errResponse.Status)
		return
	}

	result, errUpdate := c.service.UpdateTemplate(&template)

	if errUpdate != nil {
		utils.WriteResponse(w, errUpdate, errUpdate.Status)
		return
	}

	utils.WriteResponse(w, result, http.StatusOK)
}

//GetTemplate returns a given Template
func (c TemplateController) GetTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	template, err := c.service.GetTemplate(vars["id"])

	if err != nil {
		utils.WriteResponse(w, err, err.Status)
		return
	}

	utils.WriteResponse(w, template, http.StatusOK)
}

//GetAllTemplate returns all Templates
//TODO: paging
func (c TemplateController) GetAllTemplate(w http.ResponseWriter, r *http.Request) {
	templates, err := c.service.GetAllTemplates()

	if err != nil {
		utils.WriteResponse(w, err, err.Status)
		return
	}

	utils.WriteResponse(w, templates, http.StatusOK)
}

//DeleteTemplate deleted a given template
func (c TemplateController) DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	deletedTemplate, err := c.service.DeleteTemplate(vars["id"])

	if err != nil {
		utils.WriteResponse(w, err, err.Status)
		return
	}

	utils.WriteResponse(w, deletedTemplate, http.StatusOK)
}
