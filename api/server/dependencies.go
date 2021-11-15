package server

import (
	"github.com/lailaweil/billemailer/api/controllers"
	"github.com/lailaweil/billemailer/api/dao"
	"github.com/lailaweil/billemailer/api/services"
)

func resolveBillController() controllers.BillController {
	return controllers.BillController{}
}

func resolveTemplateController(db dao.DBConnection) controllers.TemplateController {
	return controllers.NewTemplateController(services.NewTemplateService(db))
}

func resolveFolderController(db dao.DBConnection) controllers.FolderController {
	return controllers.NewFolderController(services.NewFolderService(db))
}
