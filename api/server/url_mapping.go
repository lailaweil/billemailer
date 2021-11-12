package server

import (
	"github.com/gorilla/mux"
	"github.com/lailaweil/billemailer/api/controllers"
)

const (
	DELETE = "DELETE"
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
)

func mapUrls(router *mux.Router, billController controllers.BillController, templateController controllers.TemplateController)  {
	mapBillHandlers(router, billController)
	mapTemplateHandlers(router, templateController )
}


func mapBillHandlers(r *mux.Router, c controllers.BillController) {

	r.Methods(POST).Path("/bill").HandlerFunc(c.CreateBill)

	r.Methods(PUT).Path("/bill/{id:[0-9]+}").HandlerFunc(c.UpdateBill)

	r.Methods(DELETE).Path("/bill/{id:[0-9]+}").HandlerFunc( c.DeleteBill)

	r.Methods(GET).Path("/bill/{id:[0-9]+}").HandlerFunc(c.GetBill)
}

func mapTemplateHandlers(r *mux.Router, c controllers.TemplateController) {

	r.Methods(POST).Path("/template").HandlerFunc(c.CreateTemplate)

	r.Methods(PUT).Path("/template/{id:[0-9]+}").HandlerFunc(c.UpdateTemplate)

	r.Methods(DELETE).Path("/template/{id:[0-9]+}").HandlerFunc( c.DeleteTemplate)

	r.Methods(GET).Path("/bitemplatell/{id:[0-9]+}").HandlerFunc(c.GetTemplate)
}