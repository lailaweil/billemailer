package server

import (
	"github.com/gorilla/mux"
	"github.com/lailaweil/billemailer/api/dao"
)

func bootstrap(router *mux.Router) {
	emailerDB := dao.CreateEmailerDB()

	billController := resolveBillController()
	templateController := resolveTemplateController(emailerDB)
	folderController := resolveFolderController(emailerDB)

	mapUrls(router, billController, templateController, folderController)
}
