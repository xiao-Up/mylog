package mylog

import (
	"os"
)

func IsExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil && os.IsExist(err)
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	os.Chmod(path, os.ModePerm)
	return nil
}

func CreateFile(name string) (f *os.File, err error) {
	f, err = os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return f, err
}
