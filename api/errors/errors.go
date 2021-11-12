package errors

import "fmt"

type Error struct {
	status int `json:"status"`
	message string `json:"message"`
	cause string `json:"cause"`
}

func (e Error) Status() int {
	return e.status
}

func (e Error) Message() string {
	return e.message
}

func (e Error) Cause() string {
	return e.cause
}

func NewError(status int, message string, cause string) Error {
	return Error{status: status, message: message, cause: cause}
}

func (e Error) Error() string {
	return fmt.Sprintf("Message: %s; Status: %d; Cause: %s", e.message, e.status, e.cause)
}

