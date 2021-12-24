package utilities

import (
	"encoding/base64"
	"net/http"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	if encodedTime == "" {
		encodedTime = EncodeCursor(time.Time{})
	}
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	ErrorCode int         `json:"error_code"`
}

type ResponseList struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	NextCursor string      `json:"next_cursor"`
}

type ResponseListOffset struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	TotalItem int64       `json:"total_item"`
}

type ResponseStandard struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func StandardResponse(input interface{}, message string) (statusCode int, output interface{}) {
	output = ResponseStandard{
		Data:    input,
		Message: message,
	}
	statusCode = http.StatusOK
	return
}

func ErrorResponse(err error, message string, errorCode int) (statusCode int, output interface{}) {
	if message == "" {
		message = err.Error()
	}
	output = ResponseError{
		Message:   message,
		ErrorCode: errorCode,
	}
	statusCode = int(getStatusCode(err))
	return
}

func ErrorResponseWithData(err error, message string, errorCode int, data interface{}) (statusCode int, output interface{}) {
	if message == "" {
		message = err.Error()
	}
	output = ResponseError{
		Message:   message,
		Data:      data,
		ErrorCode: errorCode,
	}
	statusCode = int(getStatusCode(err))
	return
}

func ListResponse(input interface{}, nextCursor string, message string) (statusCode int, output interface{}) {
	output = ResponseList{
		Data:       input,
		NextCursor: nextCursor,
		Message:    message,
	}
	statusCode = http.StatusOK

	return
}

func ListResponseWithOffsetPagging(input interface{}, TotalItem int64, message string) (statusCode int, output interface{}) {
	output = ResponseListOffset{
		Data:      input,
		TotalItem: TotalItem,
		Message:   message,
	}
	statusCode = http.StatusOK

	return
}
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrBadParamInput:
		return 400
	case ErrUnauthorized:
		return 401
	case ErrDuplicateLogin:
		return 419
	case ErrConflict:
		return 409
	case ErrNotFound:
		return 404
	case ErrInternalServerError:
		return 500
	case ErrForbiden:
		return 403
	case ErrNeedUpdate:
		return 401

	default:
		return http.StatusInternalServerError
	}
}
