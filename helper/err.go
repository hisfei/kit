package h

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func Wrap(err error, message string) error {
	return errors.Wrap(err, "==> "+printCallerNameAndLine()+message)
}

func WithMessage(err error, message ...string) error {
	return errors.WithMessage(err, "==> "+printCallerNameAndLine()+strings.Join(message, ";"))
}

func printCallerNameAndLine() string {
	pc, _, line, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name() + "()@" + strconv.Itoa(line) + ": "
}
