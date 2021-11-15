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
	GetAllTemplates() ([]*domain.Template, *errors.Error)
	UpdateTemplate(newTemplate *domain.Template) (*domain.Template, *errors.Error)
	DeleteTemplate(id string) (*domain.Template, *errors.Error)
}

type templateService struct {
	Container dao.DBConnection
}

func NewTemplateService(container dao.DBConnection) *templateService {
	service := &templateService{
		Container: container,
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
	exist, err := t.Container.Get(id, &template)

	if !exist {
		return nil, errors.NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound), "no template found")
	}

	if err != nil {
		fmt.Errorf("error getting template %s", err.Error())
		return nil, errors.NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
	}

	return &template, nil
}

func (t templateService) GetAllTemplates() ([]*domain.Template, *errors.Error) {
	var templates []*domain.Template
	err := t.Container.GetAll(&templates)

	if err != nil {
		fmt.Errorf("error getting templates %s", err.Error())
		return nil, errors.NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
	}

	if len(templates) == 0 {
		return nil, errors.NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound), "no templates found")
	}

	return templates, nil
}

func (t templateService) UpdateTemplate(newTemplate *domain.Template) (*domain.Template, *errors.Error) {
	_, err := t.Container.Update(newTemplate)

	if err != nil {
		fmt.Errorf("error updating template %s", err.Error())
		return nil, errors.NewError(http.StatusInternalServerError, "error updating template", err.Error())
	}

	return newTemplate, nil
}

func (t templateService) DeleteTemplate(id string) (*domain.Template, *errors.Error) {
	template, err := t.GetTemplate(id)

	if err != nil {
		return nil, err
	}

	if template.ID == 0 {
		return nil, errors.NewError(http.StatusNotFound, "error deleting template", "template not found")
	}

	var deletedTemplate domain.Template
	_, errDelete := t.Container.Delete(&deletedTemplate, id)

	if errDelete != nil {
		fmt.Errorf("error deleting template %s", errDelete.Error())
		return nil, errors.NewError(http.StatusInternalServerError, "error deleting template", errDelete.Error())
	}

	return &deletedTemplate, nil
}
