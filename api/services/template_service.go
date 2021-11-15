package services

import (
	"fmt"
	"github.com/lailaweil/billemailer/api/dao"
	"github.com/lailaweil/billemailer/api/domain"
	"github.com/lailaweil/billemailer/api/errors"
	"net/http"
)

type TemplateService interface {
	CreateTemplate(template *domain.Template) (*domain.Template, *errors.Error)
	GetTemplate(id string) (*domain.Template, *errors.Error)
}

type templateService struct {
	Container dao.DBConnection
}

func NewTemplateService() *templateService {
	service := &templateService{
		Container: dao.NewDBConnection(&dao.PostgresConnection{}),
	}

	service.Container.Connect()

	return service
}

func (t templateService) CreateTemplate(template *domain.Template) (*domain.Template, *errors.Error) {
	_, err := t.Container.Insert(template)

	if err != nil {
		fmt.Errorf("error inserting template %s", err.Error())
		return nil, errors.NewError(http.StatusInternalServerError, "error inserting template", err.Error())
	}

	return template, nil
}

func (t templateService) GetTemplate(id string) (*domain.Template, *errors.Error) {
	var template domain.Template
	err := t.Container.Get(id, &template)

	if err != nil {
		fmt.Errorf("error getting template %s", err.Error())
		return nil, errors.NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
	}

	if template.ID == 0 {
		return nil, errors.NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound), "no template found")
	}

	return &template, nil
}
