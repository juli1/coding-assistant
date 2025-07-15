package find

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
	ToolName = "Find"
)

type Find struct {
	RepositoryDirectory string
	Debug               bool
}

var _ tools.Tool = Find{}

// Name returns the name of the tool.
func (g Find) Name() string {
	return ToolName
}

// Description returns the description of the tool.
func (g Find) Description() string {
	return description
}

// Call calls the tool with the given input.
func (g Find) Call(ctx context.Context, input string) (string, error) {
	var foundFiles []string

	if g.Debug {
		fmt.Printf("[find] %s\n", input)
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

		if strings.Contains(path, input) {
			foundFiles = append(foundFiles, path)
			return nil
		}

		matched, err := filepath.Match(input, relPath)
		if err != nil {
			return err
		}

		if matched {
			foundFiles = append(foundFiles, path)
			return nil
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	res := strings.Join(foundFiles, "\n")
	if g.Debug {
		fmt.Printf("[find] returning\n%s\nend\n", res)
	}

	return res, nil
}
