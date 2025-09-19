package errs

import "net/http"

type ErrMessage interface {
	Message() string
	StatusCode() int
	Error() string
}

type ErrMessageData struct {
	ErrMessage    string `json:"message" example:"This is an error message"`
	ErrStatusCode int    `json:"status_code" example:"400"`
	ErrError      string `json:"error" example:"BAD_REQUEST"`
}

func (e *ErrMessageData) Message() string {
	return e.ErrMessage
}

func (e *ErrMessageData) StatusCode() int {
	return e.ErrStatusCode
}

func (e *ErrMessageData) Error() string {
	return e.ErrError
}

func NewInternalServerError(message string) ErrMessage {
	return &ErrMessageData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      "INTERNAL_SERVER_ERROR",
	}
}

func NewNotFound(message string) ErrMessage {
	return &ErrMessageData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusNotFound,
		ErrError:      "DATA_NOT_FOUND",
	}
}

func NewFound(message string) ErrMessage {
	return &ErrMessageData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusFound,
		ErrError:      "DATA_ALREADY_EXIST",
	}
}

func NewBadRequest(message string) ErrMessage {
	return &ErrMessageData{
		ErrStatusCode: http.StatusBadRequest,
		ErrMessage:    message,
		ErrError:      "BAD_REQUEST",
	}
}

func NewUnauthorized(message string) ErrMessage {
	return &ErrMessageData{
		ErrStatusCode: http.StatusUnauthorized,
		ErrMessage:    message,
		ErrError:      "UNAUTHORIZED_USER",
	}
}

func NewUnprocessableEntity(message string) ErrMessage {
	return &ErrMessageData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusUnprocessableEntity,
		ErrError:      "INVALID_REQUEST_BODY",
	}
}
