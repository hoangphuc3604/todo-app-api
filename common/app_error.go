package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type AppErr struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewFullErrorResponse(statusCode int, rootErr error, message string, log string, key string) *AppErr {
	return &AppErr{
		StatusCode: statusCode,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(rootErr error, message string, log string, key string) *AppErr {
	return &AppErr{
		StatusCode: http.StatusBadRequest,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func (e *AppErr) RootError() error {
	if err, ok := e.RootErr.(*AppErr); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppErr) Error() string {
	return e.RootError().Error()
}

func NewCustomError(root error, message string, key string) *AppErr {
	if root != nil {
		return NewErrorResponse(root, message, root.Error(), key)
	}

	return NewErrorResponse(errors.New(message), message, message, key)
}

func ErrorDB(err error) *AppErr {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrorInvalidRequest(err error) *AppErr {
	return NewCustomError(err, "Invalid request", "INVALID_REQUEST")
}

func ErrorInternalServer(err error) *AppErr {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Internal server error", err.Error(), "INTERNAL_SERVER_ERROR")
}

func ErrorCanNotCreateEntity(entityName string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not create %s", strings.ToLower(entityName)),
		fmt.Sprintf("CAN_NOT_CREATE_%s", strings.ToUpper(entityName)),
	)
}

func ErrorCanNotUpdateEntity(entityName string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not update %s", strings.ToLower(entityName)),
		fmt.Sprintf("CAN_NOT_UPDATE_%s", strings.ToUpper(entityName)),
	)
}

func ErrorCanNotDeleteEntity(entityName string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not delete %s", strings.ToLower(entityName)),
		fmt.Sprintf("CAN_NOT_DELETE_%s", strings.ToUpper(entityName)),
	)
}

func ErrorCanNotGetEntity(entityName string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not get %s", strings.ToLower(entityName)),
		fmt.Sprintf("CAN_NOT_GET_%s", strings.ToUpper(entityName)),
	)
}

func ErrorCanNotListEntity(entityName string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not list %s", strings.ToLower(entityName)),
		fmt.Sprintf("CAN_NOT_LIST_%s", strings.ToUpper(entityName)),
	)
}

func ErrorEntityDeleted(entityName string, err error) *AppErr {
	return NewCustomError(
		err,
		fmt.Sprintf("%s has been deleted", strings.ToLower(entityName)),
		fmt.Sprintf("%s_HAS_BEEN_DELETED", strings.ToUpper(entityName)),
	)
}