package utils

import (
	"coding-assistant/internal/model"
	"io/fs"
	"path/filepath"
	"strings"
)

// GetLanguageForFile determines the programming language of a file based on its extension.
func GetLanguageForFile(path string) model.Language {
	extension := strings.ToLower(filepath.Ext(path))

	switch extension {
	case ".java":
		return model.LanguageJava
	case ".go":
		return model.LanguageGo
	case ".py", ".py3":
		return model.LanguagePython
	case ".sh":
		return model.LanguageShell
	case ".rb":
		return model.LanguageRuby
	case ".php", ".php4":
		return model.LanguagePHP
	case ".js":
		return model.LanguageJavaScript
	case ".ts":
		return model.LanguageTypeScript
	case ".html", ".html5":
		return model.LanguageHTML
	case ".css":
		return model.LanguageCSS
	default:
		return model.LanguageUnknown
	}
}

// GetFilesInDirectory recursively finds all files in a directory.
func GetFilesInDirectory(path string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(path, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, s)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
