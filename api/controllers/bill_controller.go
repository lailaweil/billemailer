package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type BillController struct {
}
//TODO: implement swagger
//TODO: implement methods

//CreateBill creates a bill
func (c BillController) CreateBill(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusCreated)
}

//UpdateBill updates a given bill
func (c BillController) UpdateBill(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id: %v\n", vars["id"])
}

//DeleteBill deletes a given bill
func (c BillController) DeleteBill(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id: %v\n", vars["id"])
}

//GetBill returns a given bill
func (c BillController) GetBill(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id: %v\n", vars["id"])
}