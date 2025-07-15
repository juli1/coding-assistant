package grep

import (
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
			return err
		}

		relPath, err := filepath.Rel(g.RepositoryDirectory, path)
		if err != nil {
			return err
		}

		if strings.HasPrefix(relPath, ".git") {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if strings.Contains(string(content), input) {
			if g.Debug {
				fmt.Printf("[find] path %s match input %s\n", path, input)
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
		fmt.Printf("[grep] returning\n%s\nend\n", res)
	}

	return res, nil
}
