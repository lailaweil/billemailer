package services

import (
	"fmt"
	"github.com/lailaweil/billemailer/api/dao"
	"github.com/lailaweil/billemailer/api/domain"
	"github.com/lailaweil/billemailer/api/errors"
	"net/http"
)

type FolderService interface {
	CreateFolder(folder *domain.Folder) (*domain.Folder, *errors.Error)
	GetFolder(id string) (*domain.Folder, *errors.Error)
}

type folderService struct {
	Container dao.DBConnection
}

func NewFolderService(container dao.DBConnection) *folderService {
	service := &folderService{
		Container: container,
	}

	service.Container.Connect()

	return service
}

func (t folderService) CreateFolder(folder *domain.Folder) (*domain.Folder, *errors.Error) {
	_, err := t.Container.Insert(folder)

	if err != nil {
		fmt.Errorf("error inserting folder %s", err.Error())
		return nil, errors.NewError(http.StatusInternalServerError, "error inserting folder", err.Error())
	}

	return folder, nil
}

func (t folderService) GetFolder(id string) (*domain.Folder, *errors.Error) {
	var folder domain.Folder
	exist, err := t.Container.Get(id, &folder, "DefaultTemplate")

	if !exist {
		return nil, errors.NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound), "no folder found")
	}

	if err != nil {
		fmt.Errorf("error getting folder %s", err.Error())
		return nil, errors.NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
	}

	return &folder, nil
}
