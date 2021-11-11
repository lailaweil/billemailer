package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lailaweil/billemailer/api/controllers"
	"net/http"
)

func mapUrls(router *mux.Router, billController controllers.BillController)  {
	billSubrouter := router.PathPrefix("/email/bill").Subrouter()

	mapBillHandlers(billSubrouter, billController)
}


func mapBillHandlers(r *mux.Router, c controllers.BillController) {

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello world post ")
	}).Methods("POST")

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello world get ")
	}).Methods("GET")
}