package domain

import (
	"errors"
)

type Folder struct {
	ID                uint      `json:"id" gorm:"primaryKey" `
	Name              string    `json:"name"`
	DefaultTemplateID *uint     `json:"default_template_id,omitempty" gorm:"column:default_template_id"`
	DefaultTemplate   *Template `json:"default_template,omitempty"`
	//Bills             []*Bill   `json:"bills,omitempty"`
}

func (f Folder) Validate() error {
	if f.Name == "" {
		return errors.New("name can't be empty")
	}

	return nil
}
