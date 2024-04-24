package errorx

import (
	"github.com/go-errors/errors"
)

// New makes an Error from the given value. If that value is already an
// error then it will be used directly, if not, it will be passed to
// fmt.Errorf("%v"). The stacktrace will point to the line of code that
// called New.
func New(text string) error {
	return errors.New(text)
}

// Is detects whether the error is equal to a given error. Errors
// are considered equal by this function if they are matched by errors.Is
// or if their contained errors are matched through errors.Is.
func Is(e error, original error) bool {
	return errors.Is(e, original)
}

// As finds the first error in err's tree that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
//
// For more information see stdlib errors.As.
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Wrap makes an Error from the given value. If that value is already an
// error then it will be used directly, if not, it will be passed to
// fmt.Errorf("%v"). The skip parameter indicates how far up the stack
// to start the stacktrace. 0 is from the current call, 1 from its caller, etc.
func Wrap(e any, skip int) error {
	return errors.Wrap(e, skip)
}

// WrapPrefix makes an Error from the given value. If that value is already an
// error then it will be used directly, if not, it will be passed to
// fmt.Errorf("%v"). The prefix parameter is used to add a prefix to the
// error message when calling Error(). The skip parameter indicates how far
// up the stack to start the stacktrace. 0 is from the current call,
// 1 from its caller, etc.
func WrapPrefix(e any, prefix string, skip int) error {
	return errors.WrapPrefix(e, prefix, skip)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
//
// Unwrap only calls a method of the form "Unwrap() error".
// In particular Unwrap does not unwrap errors returned by [Join].
//
// For more information see stdlib errors.Unwrap.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Errorf creates a new error with the given message. You can use it
// as a drop-in replacement for fmt.Errorf() to provide descriptive
// errors in return values.
func Errorf(format string, a ...any) error {
	return errors.Errorf(format, a...)
}
