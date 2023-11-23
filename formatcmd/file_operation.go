package formatcmd

import (
	"fmt"
	"golang-cli/tools"
	"os"
	"strconv"
)

type FileTransferInfo struct {
	DestinationFolderName string // col1, col2, col3, ...
	CurrentFilePath       string // /example/path/00001.png
	RenamedFileName       string // 1.png
}

/*
ファイル転送後の情報(FileTransferInfo)の配列を返却する関数
FileTransferInfoには以下の情報が含まれる
  - DestinationFolderName : 転送先のフォルダ名. col1, col2, col3, ...
  - CurrentFilePath : 転送元ファイルのfullパス. /example/path/00001.png
  - RenamedFileName : 転送先のファイル名. 1.png
*/
func createNewPathInfo(files []tools.FileObj, path string, chunkSize int, newPathIndex int) ([]*FileTransferInfo, error) {

	if chunkSize < 1 {
		return nil, fmt.Errorf("chunk size must be greater than 1")
	}
	if newPathIndex <= 0 {
		return nil, fmt.Errorf("new path index must be greater than 0")
	}

	n := len(files)

	if n%chunkSize != 0 {
		return nil, fmt.Errorf("number of files must be divisible by chunk size, the number of files(%d) is not divisible by chunk size(%d)", n, chunkSize)
	}

	pathArray := make([]*FileTransferInfo, n)
	fileNames := tools.GetSortedFileNames(files)

	for i := 0; i < n; i++ {
		folderSuffixIndex := newPathIndex + i/chunkSize
		targetFolderName := OUTPUT_FOLDER_PREFIX + strconv.Itoa(folderSuffixIndex)
		filePath := path + "/" + fileNames[i]
		renamedFileName := strconv.Itoa((i%chunkSize)+1) + IMG_FILE_EXT
		pathArray[i] = &FileTransferInfo{targetFolderName, filePath, renamedFileName}

	}

	return pathArray, nil
}

func readFileNamesByPath(path string) ([]tools.FileObj, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// NameGetterインターフェースのスライスを作成
	nameGetters := make([]tools.FileObj, len(entries))
	for i, entry := range entries {
		nameGetters[i] = entry
	}

	return nameGetters, nil
}

func getMaxFolderIndex(destinationPath string) (int, error) {
	listFolder, err := os.ReadDir(destinationPath)
	if err != nil {
		return 0, err
	}
	max := -1
	for _, folder := range listFolder {
		folderName := folder.Name()
		index, err := strconv.Atoi(folderName[len(OUTPUT_FOLDER_PREFIX):])
		if err != nil {
			return 0, err
		}
		if index > max {
			max = index
		}
	}
	return max, nil
}

func mvPath(fileObj []*FileTransferInfo, destinationPath string) error {

	// create new folders
	var newFolders []string
	for _, file := range fileObj {
		newPath := file.DestinationFolderName
		if !containsString(newFolders, newPath) {
			newFolders = append(newFolders, newPath)
			err := os.Mkdir(destinationPath+"/"+newPath, 0755)
			if err != nil {
				return err
			}
		}
	}

	// move files
	for _, file := range fileObj {
		filePath := file.CurrentFilePath
		newPath := destinationPath + "/" + file.DestinationFolderName + "/" + file.RenamedFileName
		err := tools.Mv(filePath, newPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func containsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
