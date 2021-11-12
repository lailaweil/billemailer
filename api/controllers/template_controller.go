package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type TemplateController struct {

}
//TODO: implement swagger
//TODO: implement methods

//CreateTemplate creates a Template
func (c TemplateController) CreateTemplate(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusCreated)
}

//UpdateTemplate updates a given Template
func (c TemplateController) UpdateTemplate(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id: %v\n", vars["id"])
}

//DeleteTemplate deletes a given Template
func (c TemplateController) DeleteTemplate(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id: %v\n", vars["id"])
}

//GetTemplate returns a given Template
func (c TemplateController) GetTemplate(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id: %v\n", vars["id"])
}