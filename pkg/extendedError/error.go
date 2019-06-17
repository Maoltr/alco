package extendedError

import "net/http"

func NewWithStatus(status int, message string) error {
	return Error{Status: status, Message: message}
}

func New(message string) error {
	return Error{Message: message, Status: http.StatusInternalServerError}
}

type Error struct {
	Status int
	Message string
}

func (e Error) Error() string {
	return e.Message
}
