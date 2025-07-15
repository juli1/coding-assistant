package write_file

import (
	"context"
	_ "embed"
	"os"
	"strings"

	"github.com/tmc/langchaingo/tools"
)

const (
	ToolName = "WriteFile"
)

//go:embed description.txt
var description string

// WriteFile is a tool that writes content to a file.
// It implements the tools.Tool interface.
type WriteFile struct {
	RepositoryDirectory string
	Debug               bool
}

var _ tools.Tool = WriteFile{}

// Name returns the name of the tool.
func (w WriteFile) Name() string {
	return ToolName
}

// Description returns the description of the tool.
func (w WriteFile) Description() string {
	return description
}

// Call calls the tool with the given input.
func (w WriteFile) Call(ctx context.Context, input string) (string, error) {
	parts := strings.SplitN(input, "\n", 2)
	if len(parts) != 2 {
		return "", os.ErrInvalid
	}

	filePath := parts[0]
	content := parts[1]

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return "", err
	}

	return "File written successfully", nil
}
