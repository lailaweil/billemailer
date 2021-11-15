package domain

import (
	"errors"
)

type Template struct {
	ID      uint     `json:"id" gorm:"primaryKey" `
	Body    string   `json:"body" `
	Subject string   `json:"subject"`
	Folders []Folder `gorm:"ForeignKey:DefaultTemplateID"`
}

func (t Template) Validate() error {
	if t.Subject == "" {
		return errors.New("subject can't be empty")
	}

	if t.Body == "" {
		return errors.New("body can't be empty")
	}

	return nil
}
