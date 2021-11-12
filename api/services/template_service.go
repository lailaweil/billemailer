package services

import (
	"github.com/lailaweil/billemailer/api/domain"
	"github.com/lailaweil/billemailer/api/errors"
	"log"
	"os"
	textTemplate "text/template"
)

type TemplateService interface {
	CreateTemplate(template domain.Template) *errors.Error
}

type templateService struct {
	//TODO: Database
}

func NewTemplateService() *templateService {
	return &templateService{}
}

func (t templateService) CreateTemplate(template domain.Template) *errors.Error {
	type Bill struct {
		Date, Sum, Reason, Name string
	}
	
	var bill = Bill{
		Date:   "10/11/2021",
		Sum:    "10055,6",
		Reason: "Expensas Septiembre 2020",
		Name:   "Laila Weil",
	}
	// Create a new template and parse the body into it.
	email := textTemplate.Must(textTemplate.New("email").Parse(template.Body))

	// Execute the template for each bill.
	err := email.Execute(os.Stdout, bill)
	if err != nil {
		log.Println("executing template:", err)
	}

	return nil
}


