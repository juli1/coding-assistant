package ls

import (
	"context"
	_ "embed"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/tools"
)

const (
	ToolName = "ls"
)

//go:embed description.txt
var description string

var _ tools.Tool = &Ls{}

type Ls struct {
	RepositoryDirectory string
	Debug               bool
}

func (l Ls) Name() string {
	return ToolName
}

func (l Ls) Description() string {
	return description
}

func (l Ls) Call(ctx context.Context, input string) (string, error) {
	path := input
	if path == "" {
		path = "."
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("could not read directory: %w", err)
	}

	var output string
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return "", fmt.Errorf("could not get file info: %w", err)
		}

		// Format permissions
		permissions := info.Mode().String()

		output += fmt.Sprintf("%s %s\n", permissions, file.Name())
	}

	return output, nil
}
