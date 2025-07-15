package read_file

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/langchaingo/tools"
)

//go:embed description.txt
var description string

const (
	ToolName = "ReadFile"
)

// ReadFile is a tool that reads the content of a file.
// It implements the tools.Tool interface.
type ReadFile struct {
	RepositoryDirectory string
	Debug               bool
}

var _ tools.Tool = ReadFile{}

// Name returns the name of the tool.
func (r ReadFile) Name() string {
	return ToolName
}

// Description returns the description of the tool.
func (r ReadFile) Description() string {
	return description
}

// Call calls the tool with the given input.
func (r ReadFile) Call(ctx context.Context, input string) (string, error) {
	path := filepath.Join(r.RepositoryDirectory, input)

	if r.Debug {
		fmt.Printf("[read tool] Reading file %s\n", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
