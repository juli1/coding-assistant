package searchtext

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/tmc/langchaingo/tools"
)

//go:embed description.txt
var toolDescription string

const ToolName = "SearchText"

var _ tools.Tool = &SearchText{}

type SearchText struct {
	RepositoryDirectory string
	Debug               bool
}

func (st *SearchText) Name() string {
	return ToolName
}

func (st *SearchText) Description() string {
	return toolDescription
}

func (st *SearchText) Call(ctx context.Context, input string) (string, error) {
	re, err := regexp.Compile(input)
	if err != nil {
		return "", fmt.Errorf("invalid regex pattern: %w", err)
	}

	var matchingFiles []string

	err = filepath.Walk(st.RepositoryDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return nil // Continue walking even if a file can't be opened
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if re.Match(scanner.Bytes()) {
					matchingFiles = append(matchingFiles, path)
					break // Move to the next file once a match is found
				}
			}
		}
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("error walking directory: %w", err)
	}

	var output string
	for _, file := range matchingFiles {
		output += file + "\n"
	}

	return output, nil
}
