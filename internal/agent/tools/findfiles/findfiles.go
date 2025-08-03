package findfiles

import (
	"context"
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/tmc/langchaingo/tools"
)

//go:embed description.txt
var toolDescription string

const ToolName = "FindFiles"

var _ tools.Tool = &FindFiles{}

type FindFiles struct {
	RepositoryDirectory string
	Debug               bool
}

func (f *FindFiles) Name() string {
	return ToolName
}

func (f *FindFiles) Description() string {
	return toolDescription
}

func (f *FindFiles) Call(ctx context.Context, input string) (string, error) {
	files, err := filepath.Glob(input)
	if err != nil {
		return "", fmt.Errorf("could not glob pattern: %w", err)
	}

	if f.Debug {
		fmt.Printf("[findfiles] %s\n", input)
	}

	var output string
	for _, file := range files {
		output += file + "\n"
	}

	return output, nil
}
