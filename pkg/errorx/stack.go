package errorx

import (
	"log/slog"

	"github.com/go-errors/errors"
)

type stackFrame struct {
	Func   string `json:"func"`
	Source string `json:"source"`
	Line   int    `json:"line"`
}

// ReplaceAttr replace value of slog.Any.
func ReplaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Value.Kind() != slog.KindAny {
		return a
	}

	if v, ok := a.Value.Any().(error); ok {
		a.Value = fmtErr(v)
	}

	return a
}

// marshalStack extracts stack frames from the error.
func marshalStack(err error) []stackFrame {
	sErr := isStackErr(err)
	if sErr == nil {
		return nil
	}

	frames := sErr.StackFrames()

	out := make([]stackFrame, len(frames))

	for index, frame := range frames {
		source, scErr := frame.SourceLine()
		if scErr != nil {
			slog.Warn("[!] stack frames get source line error", slog.Any("err", scErr))
			continue
		}

		out[index] = stackFrame{
			Func:   frame.Name,
			Source: source,
			Line:   frame.LineNumber,
		}
	}

	return out
}

func isStackErr(err error) *errors.Error {
	var sErr *errors.Error
	if ok := As(err, &sErr); ok {
		return sErr
	}

	if iErr := IsInternalErr[any](err); iErr != nil {
		return iErr.StackError()
	}

	return nil
}

// fmtErr returns a slog.Value with keys `msg` and `trace`. If the error
// does not implement interface { StackFrames() errors.StackFrames }, the `trace`
// key is omitted.
func fmtErr(err error) slog.Value {
	var groupValues []slog.Attr
	groupValues = append(groupValues, slog.String("msg", err.Error()))

	frames := marshalStack(err)

	if frames != nil {
		groupValues = append(groupValues,
			slog.Any("trace", frames),
		)
	}

	return slog.GroupValue(groupValues...)
}
