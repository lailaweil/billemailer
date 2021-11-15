package services

import (
	"encoding/json"
	errors2 "errors"
	"fmt"
	"github.com/lailaweil/billemailer/api/dao"
	"github.com/lailaweil/billemailer/api/domain"
	"github.com/lailaweil/billemailer/api/errors"
	mocks "github.com/lailaweil/billemailer/api/mocks/dao"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestTemplateService_GetTemplate(t *testing.T) {
	cases := []struct {
		name   string
		id     string
		result *domain.Template
		exists bool
		err    *errors.Error
		errDB  error
	}{
		{
			name: "OK",
			id:   "1",
			result: &domain.Template{
				ID:      1,
				Body:    "test template body",
				Subject: "test template subject",
			},
			exists: true,
		},
		{
			name:   "NOT_FOUND",
			id:     "2",
			result: nil,
			exists: false,
			err:    errors.NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound), "no template found"),
		},
		{
			name:   "ERR_DB",
			id:     "2",
			result: nil,
			exists: true,
			err:    errors.NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "err db"),
			errDB:  errors2.New("err db"),
		},
	}

	for _, c := range cases {
		service := &templateService{}

		mockDB := new(mocks.GenericDB)
		mockDB.On("Get", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
			if c.result != nil {
				bytes, _ := json.Marshal(c.result)
				json.Unmarshal(bytes, args.Get(1))
			}
		}).Return(c.exists, c.errDB)
		service.Container = dao.NewDBConnection(mockDB)

		result, err := service.GetTemplate(c.id)

		if c.err != nil {
			assert.Equal(t, c.err, err, fmt.Sprintf("CASE:%s", c.name))
		} else {
			assert.Equal(t, c.result, result, fmt.Sprintf("CASE:%s", c.name))
		}

		mockDB.AssertExpectations(t)
	}
}

func TestTemplateService_GetAllTemplates(t *testing.T) {
	cases := []struct {
		name   string
		result []*domain.Template
		err    *errors.Error
		errDB  error
	}{
		{
			name: "OK",
			result: []*domain.Template{
				&domain.Template{
					ID:      1,
					Body:    "test template body",
					Subject: "test template subject",
				},
			},
		},
		{
			name:   "NOT_FOUND",
			result: []*domain.Template{},
			err:    errors.NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound), "no templates found"),
			errDB:  nil,
		},
		{
			name:   "ERR_DB",
			result: []*domain.Template{},
			err:    errors.NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "err db"),
			errDB:  errors2.New("err db"),
		},
	}

	for _, c := range cases {
		service := &templateService{}

		mockDB := new(mocks.GenericDB)
		mockDB.On("GetAll", mock.Anything).Run(func(args mock.Arguments) {
			if c.result != nil {
				bytes, _ := json.Marshal(c.result)
				json.Unmarshal(bytes, args.Get(0))
			}
		}).Return(c.errDB)
		service.Container = dao.NewDBConnection(mockDB)

		result, err := service.GetAllTemplates()

		if c.err != nil {
			assert.Equal(t, c.err, err, fmt.Sprintf("CASE:%s", c.name))
		} else {
			assert.Equal(t, c.result, result, fmt.Sprintf("CASE:%s", c.name))
		}

		mockDB.AssertExpectations(t)
	}
}

func TestTemplateService_CreateTemplate(t *testing.T) {
	cases := []struct {
		name     string
		template domain.Template
		err      *errors.Error
		errDB    error
	}{
		{
			name: "OK",
			template: domain.Template{
				ID:      1,
				Body:    "test template body",
				Subject: "test template subject",
			},
		},
		{
			name:     "ERR",
			template: domain.Template{},
			err:      errors.NewError(http.StatusInternalServerError, "error inserting template", "err db"),
			errDB:    errors2.New("err db"),
		},
	}

	for _, c := range cases {
		service := &templateService{}

		mockDB := new(mocks.GenericDB)
		mockDB.On("Insert", mock.Anything).Run(func(args mock.Arguments) {
			if c.template.ID == 0 {
				bytes, _ := json.Marshal(c.template)
				json.Unmarshal(bytes, args.Get(0))
			}
		}).Return(c.template, c.errDB)
		service.Container = dao.NewDBConnection(mockDB)

		result, err := service.CreateTemplate(&c.template)

		if c.err != nil {
			assert.Equal(t, c.err, err, fmt.Sprintf("CASE:%s", c.name))
		} else {
			assert.Equal(t, &c.template, result, fmt.Sprintf("CASE:%s", c.name))
		}

	}
}
