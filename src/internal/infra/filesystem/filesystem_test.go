package filesystem

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestListFolder(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create test files and directories
	testFiles := []struct {
		name      string
		isDir     bool
		extension string
	}{
		{"file1.txt", false, "txt"},
		{"file2.log", false, "log"},
		{"file3.txt", false, "txt"},
		{".hiddenfile", false, ""},
		{"subdir", true, ""},
	}

	for _, file := range testFiles {
		if file.isDir {
			err := os.Mkdir(filepath.Join(tempDir, file.name), 0755)
			if err != nil {
				t.Fatalf("Failed to create test directory: %v", err)
			}
		} else {
			f, err := os.Create(filepath.Join(tempDir, file.name))
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}
			f.Close()
		}
	}

	// Test cases
	tests := []struct {
		name      string
		extension string
		expected  int
	}{
		{"All files", "", 3}, // Excludes hidden files and directories
		{"Only .txt files", "txt", 2},
		{"Only .log files", "log", 1},
		{"Non-existent extension", "csv", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, _ := ListFolder(tempDir, tt.extension)
			log.Printf("%v - %v ", tt.extension, files)
			if len(files) != tt.expected {
				t.Errorf("Expected %d files, got %d", tt.expected, len(files))
			}
		})
	}
}

func TestListFolder_ErrorOnReadDir(t *testing.T) {
	_, err := ListFolder("/invalid/path", "")
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
