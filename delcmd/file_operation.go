package delCmd

import (
	"fmt"
	"io/fs"
	"log"
)

type FileOperator interface {
	RemoveAll(string) error
	Remove(string) error
	Stat(name string) (fs.FileInfo, error)
	Mkdir(path string) error
}

func del(operator FileOperator, path string) error {
	fileInfo, err := operator.Stat(path)
	if err != nil {
		fmt.Printf("failed to get file info: %v\n", err)
		return nil // return nil to avoid exit status 1
	}
	if fileInfo.IsDir() {
		err = operator.RemoveAll(path)
		if err != nil {
			fmt.Printf("failed to delete file or path: %v\n", err)
			return err
		}

		log.Printf("try to recreate directory: %s\n", path)
		err = operator.Mkdir(path)
		if err != nil {
			fmt.Printf("failed to create directory: %v\n", err)
			return err
		}
	} else {
		err = operator.Remove(path)
	}
	if err != nil {
		fmt.Printf("failed to delete file or path: %v\n", err)
		return err
	}
	return nil
}
