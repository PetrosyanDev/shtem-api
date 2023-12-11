// Erik Petrosyan ©
package domain

import "net/http"

var (
	ErrAccessDenied = NewError().SetStatus(http.StatusForbidden).SetMessage("ERR_ACCESS_DENIED")
	ErrBadRequest   = NewError().SetStatus(http.StatusBadRequest).SetMessage("ERR_BAD_REQUEST_PARAMS")
	ErrNoRows       = NewError().SetStatus(http.StatusBadRequest).SetMessage("ERR_BAD_REQUEST_PARAMS_1")
)

type Error interface {
	SetStatus(s int) Error
	SetMessage(m string) Error
	GetStatus() int
	GetMessage() string
	SetError(err error) Error
	RawError() error
}

type errorStr struct {
	error
	status  int
	message string
}

func (e *errorStr) Error() string {
	if e.error == nil {
		return string(e.message)
	}
	return e.error.Error()
}

func (e *errorStr) SetStatus(s int) Error {
	e.status = s
	return e
}

func (e *errorStr) SetMessage(m string) Error {
	e.message = m
	return e
}

func (e errorStr) SetError(err error) Error {
	e.error = err
	return &e
}

func (e *errorStr) GetStatus() int {
	if e.status == 0 {
		return http.StatusInternalServerError
	}
	return e.status
}

func (e *errorStr) GetMessage() string {
	if e.message == "" && e.error != nil {
		return e.error.Error()
	}
	return e.message
}

func (e *errorStr) RawError() error {
	return e.error
}

func NewError() Error {
	return &errorStr{nil, http.StatusInternalServerError, "ERR_INTERNAL_SEE_LOGS"}
}
