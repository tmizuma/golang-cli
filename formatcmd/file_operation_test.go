package formatcmd

import (
	"golang-cli/tools"
	"strconv"
	"testing"
)

type MockFileObj struct {
	name string
}

func (m MockFileObj) Name() string {
	return m.name
}

func TestCreateNewFileInfoSuccess(t *testing.T) {
	files := []tools.FileObj{
		MockFileObj{"00001.png"},
		MockFileObj{"00002.png"},
		MockFileObj{"00003.png"},
		MockFileObj{"00004.png"},
		MockFileObj{"00005.png"},
		MockFileObj{"00006.png"},
		MockFileObj{"00007.png"},
		MockFileObj{"00008.png"},
		MockFileObj{"00009.png"},
	}
	path := "\\example\\path"
	chunkSize := 3
	newPathIndex := 10

	infos, err := createNewPathInfo(files, path, chunkSize, newPathIndex)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(infos) != len(files) {
		t.Errorf("Expected slice length to be %d, got: %d", len(files), len(infos))
	}

	for i, info := range infos {
		actualDestinationFolderName := OUTPUT_FOLDER_PREFIX + strconv.Itoa(i/chunkSize+newPathIndex)
		if info.DestinationFolderName != actualDestinationFolderName {
			t.Errorf("Expected destination folder name to be %s, got: %s", actualDestinationFolderName, info.DestinationFolderName)
		}

		actualCurrentFilePath := path + "\\" + ("0000" + strconv.Itoa(i+1)) + IMG_FILE_EXT
		if info.CurrentFilePath != actualCurrentFilePath {
			t.Errorf("Expected current file path to be %s, got: %s", actualCurrentFilePath, info.CurrentFilePath)
		}

		actualRenamedFileName := strconv.Itoa(i%chunkSize+1) + IMG_FILE_EXT
		if info.RenamedFileName != actualRenamedFileName {
			t.Errorf("Expected renamed file name to be %s, got: %s", actualRenamedFileName, info.RenamedFileName)
		}
	}
}

func TestCreateNewFileInfoFailed(t *testing.T) {
	path := "/example/path"

	tests := []struct {
		name         string
		files        []tools.FileObj
		chunkSize    int
		newPathIndex int
	}{
		{"invalid chunkSize", []tools.FileObj{
			MockFileObj{"00001.png"},
		}, 0, 10},
		{"invalid newPathIndex", []tools.FileObj{
			MockFileObj{"00001.png"},
		}, 1, 0},
		{"invalid conbination with file size and chunk size", []tools.FileObj{
			MockFileObj{"00001.png"},
		}, 5, 11},
	}

	for _, test := range tests {
		_, err := createNewPathInfo(test.files, path, test.chunkSize, test.newPathIndex)
		if err == nil {
			t.Errorf("Expected error %s, got: %v", test.name, err)
		}
	}

}
