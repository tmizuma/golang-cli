package delCmd

import (
	"fmt"
	"io/fs"
)

type FileOperator interface {
	RemoveAll(string) error
	Remove(string) error
	Stat(name string) (fs.FileInfo, error)
}

func del(operator FileOperator, path string) error {
	fileInfo, err := operator.Stat(path)
	if err != nil {
		fmt.Printf("failed to get file info: %v\n", err)
		return nil // return nil to avoid exit status 1
	}
	if fileInfo.IsDir() {
		err = operator.RemoveAll(path)
	} else {
		err = operator.Remove(path)
	}
	if err != nil {
		fmt.Printf("failed to delete file or path: %v\n", err)
		return err
	}
	return nil
}
