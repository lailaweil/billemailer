package server

import (
	"github.com/gorilla/mux"
)

func bootstrap(router *mux.Router)  {
	billController := resolveBillController()
	templateController := resolveTemplateController()

	mapUrls(router, billController, templateController)
}
