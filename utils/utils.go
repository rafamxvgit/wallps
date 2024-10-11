package utils

import (
	"os"
	"path"
)

// returns the path to the executable's directory
func ProgramDirectory() string {
	exe := Expect(os.Executable())
	exeDirPath := path.Dir(exe)
	return exeDirPath
}

func ProgramPath() string {
	exe := Expect(os.Executable())
	return exe
}

// returns the value and consumes the error
func Expect[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
