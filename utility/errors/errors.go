package errors

import "net/http"

type RestErrors struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

func NewBadRequestError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusBadRequest,
	}

}

func NewBaInternalServerErrordRequestError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusInternalServerError,
	}

}
