package errors

import "fmt"

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func NewError(status int, message string, cause string) *Error {
	return &Error{Status: status, Message: message, Cause: cause}
}

func (e Error) Error() string {
	return fmt.Sprintf("Message: %s; Status: %d; Cause: %s", e.Message, e.Status, e.Cause)
}
