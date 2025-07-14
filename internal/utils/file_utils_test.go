package utils

import (
	"coding-assistant/internal/model"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestGetLanguageForFile(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		expected model.Language
	}{
		{"Java file", "test.java", model.LanguageJava},
		{"Go file", "main.go", model.LanguageGo},
		{"Python file", "script.py", model.LanguagePython},
		{"Python 3 file", "script.py3", model.LanguagePython},
		{"Shell script", "run.sh", model.LanguageShell},
		{"Ruby file", "app.rb", model.LanguageRuby},
		{"PHP file", "index.php", model.LanguagePHP},
		{"PHP 4 file", "server.php4", model.LanguagePHP},
		{"JavaScript file", "script.js", model.LanguageJavaScript},
		{"TypeScript file", "component.ts", model.LanguageTypeScript},
		{"HTML file", "index.html", model.LanguageHTML},
		{"HTML5 file", "page.html5", model.LanguageHTML},
		{"CSS file", "style.css", model.LanguageCSS},
		{"Unknown extension", "README.md", model.LanguageUnknown},
		{"No extension", "filewithnoextension", model.LanguageUnknown},
		{"Dotfile", ".bashrc", model.LanguageUnknown},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetLanguageForFile(tc.filePath)
			if actual != tc.expected {
				t.Errorf("For file '%s', expected language '%s', but got '%s'", tc.filePath, tc.expected, actual)
			}
		})
	}
}

func TestGetFilesInDirectory(t *testing.T) {
	// Create a temporary directory structure for testing
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create some files and directories
	files := []string{
		filepath.Join(tmpDir, "file1.txt"),
		filepath.Join(tmpDir, "dir1", "file2.txt"),
		filepath.Join(tmpDir, "dir1", "dir2", "file3.txt"),
	}

	for _, file := range files {
		if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
			t.Fatal(err)
		}
		if _, err := os.Create(file); err != nil {
			t.Fatal(err)
		}
	}

	// Test the function
	actualFiles, err := GetFilesInDirectory(tmpDir)
	if err != nil {
		t.Fatalf("GetFilesInDirectory failed: %v", err)
	}

	// Sort both slices for consistent comparison
	sort.Strings(files)
	sort.Strings(actualFiles)

	if len(actualFiles) != len(files) {
		t.Errorf("Expected %d files, but got %d", len(files), len(actualFiles))
	}

	for i := range files {
		if files[i] != actualFiles[i] {
			t.Errorf("Expected file '%s', but got '%s'", files[i], actualFiles[i])
		}
	}
}
