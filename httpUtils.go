package utils

import (
	"errors"
	"gorm.io/gorm"
)

type HTTPResponse struct {
	Status int                    `json:"status"`
	Data   any                    `json:"data"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
}

type HTTPError struct {
	Message string `json:"message"`
}

type HTTPErrorResponse struct {
	Status int       `json:"status"`
	Error  HTTPError `json:"error"`
}

func NewHTTPResponse(status int, data any, meta ...map[string]interface{}) (int, *HTTPResponse) {
	var metaData map[string]interface{}
	if meta != nil {
		metaData = meta[0]
	}
	return status, &HTTPResponse{
		Status: status,
		Data:   data,
		Meta:   metaData,
	}
}

func NewHTTPErrorResponse(err error, status ...int) (int, *HTTPErrorResponse) {
	httpStatus := 500
	if status != nil {
		httpStatus = status[0]
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		httpStatus = 404
	}
	return httpStatus, &HTTPErrorResponse{
		Status: httpStatus,
		Error: HTTPError{
			Message: err.Error(),
		},
	}
}
