package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lailaweil/billemailer/api/domain"
	"github.com/lailaweil/billemailer/api/errors"
	"github.com/lailaweil/billemailer/api/services"
	"net/http"
)

type TemplateController struct {
	service services.TemplateService
}

func NewTemplateController(service services.TemplateService) TemplateController {
	return TemplateController{service: service}
}

//TODO: implement swagger
//TODO: implement methods

//CreateTemplate creates a Template
func (c TemplateController) CreateTemplate(w http.ResponseWriter, r *http.Request) {
	// Declare a new Template.
	template := &domain.Template{}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(template); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error decoding body", err.Error())
		http.Error(w, errResponse.Error(), errResponse.Status())
		return
	}

	if err := template.Validate(); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error validating body", err.Error())
		http.Error(w, errResponse.Error(), errResponse.Status())
		return
	}

	if template.ID != 0 {
		errResponse := errors.NewError(http.StatusBadRequest, "error validating body", "id must be empty")
		http.Error(w, errResponse.Error(), errResponse.Status())
		return
	}

	//CreateTemplate
	result, errCreate := c.service.CreateTemplate(template)

	if errCreate != nil {
		http.Error(w, errCreate.Error(), errCreate.Status())
		return
	}
	response, errMarshal := json.Marshal(result)
	if errMarshal != nil {
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

//UpdateTemplate updates a given Template
func (c TemplateController) UpdateTemplate(w http.ResponseWriter, r *http.Request) {
	template := &domain.Template{}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(template); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error decoding body", err.Error())
		http.Error(w, errResponse.Error(), errResponse.Status())
		return
	}

	if err := template.Validate(); err != nil {
		errResponse := errors.NewError(http.StatusBadRequest, "error validating body", err.Error())
		http.Error(w, errResponse.Error(), errResponse.Status())
		return
	}

	result, errUpdate := c.service.UpdateTemplate(template)

	if errUpdate != nil {
		http.Error(w, errUpdate.Error(), errUpdate.Status())
		return
	}
	response, errMarshal := json.Marshal(result)
	if errMarshal != nil {
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

//GetTemplate returns a given Template
func (c TemplateController) GetTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	template, err := c.service.GetTemplate(vars["id"])

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), err.Status())
		return
	}

	response, errMarshal := json.Marshal(template)
	if errMarshal != nil {
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

//GetAllTemplate returns all Templates
//TODO: paging
func (c TemplateController) GetAllTemplate(w http.ResponseWriter, r *http.Request) {
	templates, err := c.service.GetAllTemplates()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), err.Status())
		return
	}

	response, errMarshal := json.Marshal(templates)
	if errMarshal != nil {
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

//DeleteTemplate deleted a given template
func (c TemplateController) DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	deletedTemplate, err := c.service.DeleteTemplate(vars["id"])

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), err.Status())
		return
	}

	response, errMarshal := json.Marshal(deletedTemplate)
	if errMarshal != nil {
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}
