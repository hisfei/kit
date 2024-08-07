package h

import (
	"errors"
	"io"
	"os"
	"syscall"
)

func WriteFile(filename, str string) error {
	var f *os.File
	var err error
	f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = io.WriteString(f, str)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}
func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	var pathErr *os.PathError
	ok := errors.As(err, &pathErr)
	if ok && errors.Is(syscall.ENOENT, pathErr.Err) {
		return false
	}
	return true
}
