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
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseList struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	NextCursor string      `json:"next_cursor"`
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

func ErrorResponse(err error, message string) (statusCode int, output interface{}) {
	if message == "" {
		message = err.Error()
	}
	output = ResponseError{
		Message: message,
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
	return
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrBadParamInput:
		return 40000
	case ErrUnauthorized:
		return 40001
	case ErrDuplicateLogin:
		return 40002
	case ErrConflict:
		return 40003
	case ErrNotFound:
		return 40004
	case ErrInternalServerError:
		return 40005
	case ErrForbiden:
		return 40006

	default:
		return http.StatusInternalServerError
	}
}
