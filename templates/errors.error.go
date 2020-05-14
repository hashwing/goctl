//pkg/errors/error.go
package errors

import (
	"encoding/json"
)

type ErrCode int16

const (
	ErrCodeTypeCommon ErrCode = 10000 + iota*1000
)

type IError interface {
	Error() string
	String() string
	SetErrDetail(detail string)
}

type ErrBase struct {
	ErrID     string  `json:"err_id"`
	ErrCode   ErrCode `json:"err_code"`
	ErrDetail string  `json:"err_detail"`
	Message   string  `json:"message"`
}

// Init intial error
func (e *ErrBase) Init(errCode ErrCode, message string) {
	e.ErrCode = errCode
	e.Message = message
}

// String return error message as string
func (e *ErrBase) String() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// Error return error message
func (e *ErrBase) Error() string {
	return e.String()
}

// SetTrack set error detail
func (e *ErrBase) SetErrDetail(detail string) {
	e.ErrDetail = detail
}
