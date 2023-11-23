package tools

import (
	"os"
	"sort"
)

type FileObj interface {
	Name() string
}

func Mv(old string, new string) error {
	err := os.Rename(old, new)
	if err != nil {
		return err
	}
	return nil
}

func GetSortedFileNames(files []FileObj) []string {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	sortedNames := make([]string, len(files))
	for i, file := range files {
		sortedNames[i] = file.Name()
	}
	return sortedNames
}
