package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewBadRequestError() error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: "bad request error",
	}
}

func NewVaildationError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewUnauthorizedError() error {
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: "unauthenticated error",
	}
}
