package tools

import (
	"testing"
)

// FileObj is a mock implementation of the FileObj interface for testing
type MockFileObj struct {
	name string
}

func (m MockFileObj) Name() string {
	return m.name
}

func TestGetSortedFileNames(t *testing.T) {
	tests := []struct {
		name     string
		input    []FileObj
		expected []string
	}{
		{
			name: "sorted input",
			input: []FileObj{
				MockFileObj{name: "a.txt"},
				MockFileObj{name: "b.txt"},
				MockFileObj{name: "c.txt"},
			},
			expected: []string{"a.txt", "b.txt", "c.txt"},
		},
		{
			name: "unsorted input",
			input: []FileObj{
				MockFileObj{name: "b.txt"},
				MockFileObj{name: "c.txt"},
				MockFileObj{name: "a.txt"},
			},
			expected: []string{"a.txt", "b.txt", "c.txt"},
		},
		{
			name:     "empty input",
			input:    []FileObj{},
			expected: []string{},
		},
		// Add more test cases as necessary...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetSortedFileNames(tt.input)
			if !equalSlices(result, tt.expected) {
				t.Errorf("GetSortedFileNames(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// equalSlices checks if two slices of strings are equal
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
