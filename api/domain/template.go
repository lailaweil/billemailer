package domain

import "errors"

type Template struct {
	ID string `json:"id,omitempty"`
	Body string `json:"body"`
	Subject string `json:"subject"`
}

func (t Template) Validate() error {
	if t.ID != ""{
		return errors.New("id must be empty")
	}

	if t.Subject == ""{
		return errors.New("subject can't be empty")
	}

	if t.Body == ""{
		return errors.New("body can't be empty")
	}

	return nil
}