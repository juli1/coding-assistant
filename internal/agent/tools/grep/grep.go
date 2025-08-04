package grep

import (
	toolsModule "coding-assistant/internal/agent/tools"
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tmc/langchaingo/tools"
)

//go:embed description.txt
var description string

const (
	ToolName = "Grep"
)

// Grep is a tool that executes a grep command.
// It implements the tools.Tool interface.
type Grep struct {
	RepositoryDirectory string
	Debug               bool
}

var _ tools.Tool = Grep{}

// Name returns the name of the tool.
func (g Grep) Name() string {
	return ToolName
}

// Description returns the description of the tool.
func (g Grep) Description() string {
	return description
}

// Call calls the tool with the given input.
func (g Grep) Call(ctx context.Context, input string) (string, error) {
	var foundFiles []string

	if g.Debug {
		fmt.Printf("[grep] %s\n", input)
	}

	err := filepath.Walk(g.RepositoryDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		relPath, err := filepath.Rel(g.RepositoryDirectory, path)
		if err != nil {
			return nil
		}

		for _, d := range toolsModule.DirectoriesToIgnore {
			if strings.HasPrefix(relPath, d) {
				return nil
			}
		}

		if info.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		for _, i := range toolsModule.FindFilesIgnoreSuffix {
			if strings.HasSuffix(relPath, i) {
				return nil
			}
		}

		if strings.Contains(string(content), input) {
			if g.Debug {
				fmt.Printf("[grep] path %s match input %s\n", path, input)
			}
			foundFiles = append(foundFiles, path)
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	res := strings.Join(foundFiles, "\n")
	if g.Debug {
		fmt.Printf("[grep] returning\n%s\n", res)
		fmt.Printf("[grep] end returning\n")
	}

	return res, nil
}
