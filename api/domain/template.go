package domain

import "errors"

type Template struct {
	tableName struct{} `pg:"template"`
	ID        int      `json:"id,omitempty" pg:"id,pk"`
	Body      string   `json:"body" pg:"body"`
	Subject   string   `json:"subject" pg:"subject"`
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
