//pkg/errors/common.go
package errors

import (
	"{{ .Mod }}/core"
)

const (
	ErrCodeInvalidParam ErrCode = iota + ErrCodeTypeCommon
	ErrCodeResourceNotFound
	ErrCodeUnknown
	ErrCodeMethodNotImplemented
	ErrCodeInternal
	ErrCodeMethodDeprecated
)

type ErrResource struct {
	ErrBase
	ResourceType string `json:"resource_type"`
	ResourceId   string `json:"resource_id"`
}

type ErrResourceNotFound struct {
	ErrResource
}

type ErrUnknown struct {
	ErrBase
}

type ErrInvalidParam struct {
	ErrBase
}

type ErrMethodNotImplemented struct {
	ErrBase
}

type ErrInternal struct {
	ErrBase
}

type ErrMethodDeprecated struct {
	ErrBase
}

func NewErrInvalidParam(msg string) IError {
	e := new(ErrInvalidParam)
	if msg == "" {
		e.Message = "参数错误"
	} else {
		e.Message = msg
	}
	e.Init(ErrCodeInvalidParam, e.Message)
	return e
}

func NewErrUnknown() IError {
	e := new(ErrUnknown)
	e.Init(ErrCodeUnknown, "未知错误")
	return e
}

func NewErrResourceNotFound(resourceType, resourceId string) IError {
	e := new(ErrResourceNotFound)
	e.Init(ErrCodeResourceNotFound, core.GetResourceTypeName(resourceType)+"不存在")
	e.ResourceType = resourceType
	e.ResourceId = resourceId
	return e
}

func NewErrMethodNotImplemented() IError {
	e := new(ErrMethodNotImplemented)
	e.Init(ErrCodeMethodNotImplemented, "方法未实现")
	return e
}

func NewErrInternal() IError {
	e := new(ErrInternal)
	e.Init(ErrCodeInternal, "内部错误")
	return e
}

func NewErrMethodDeprecated() IError {
	e := new(ErrMethodDeprecated)
	e.Init(ErrCodeMethodDeprecated, "该方法已经弃用，请调用其它方法")
	return e
}
