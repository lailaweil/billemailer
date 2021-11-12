package server

import (
	"github.com/lailaweil/billemailer/api/controllers"
	"github.com/lailaweil/billemailer/api/services"
)

func resolveBillController() controllers.BillController{
	return controllers.BillController{}
}

func resolveTemplateController() controllers.TemplateController{
	return controllers.NewTemplateController(services.NewTemplateService())
}