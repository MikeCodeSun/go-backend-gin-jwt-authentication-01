package util

import "net/http"

type Err struct {
	Message string
	StatusCode int
}

func CustomError(msg string) *Err {
	return &Err{
		Message: msg,
		StatusCode: http.StatusInternalServerError,
	}
}