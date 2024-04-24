package errorx

import (
	"fmt"

	"github.com/go-errors/errors"
)

type InternalError[T any] struct {
	code   ErrCode
	err    error
	params T
}

func (i *InternalError[T]) Code() ErrCode {
	return i.code
}

func (i *InternalError[T]) Error() string {
	if i.err != nil {
		return i.err.Error()
	}

	return fmt.Sprint(i.code)
}

func (i *InternalError[T]) Params() T {
	return i.params
}

func (i *InternalError[T]) WithError(err error) *InternalError[T] {
	i.err = err
	return i
}

func (i *InternalError[T]) WithParams(params T) *InternalError[T] {
	i.params = params
	return i
}

func (i *InternalError[T]) StackError() *errors.Error {
	if i.err == nil {
		return nil
	}

	var eErr *errors.Error
	if ok := As(i.err, &eErr); ok {
		return eErr
	}

	return nil
}

func NewInternalErr[T any](code ErrCode) *InternalError[T] {
	return &InternalError[T]{
		code: code,
	}
}

func IsInternalErr[T any](err error) *InternalError[T] {
	if err == nil {
		return nil
	}

	var iErr *InternalError[T]
	if ok := As(err, &iErr); ok {
		return iErr
	}

	return nil
}

func InternalCode[T any](err error) ErrCode {
	if iErr := IsInternalErr[T](err); iErr != nil {
		return iErr.Code()
	}

	return CodeUnknown
}
