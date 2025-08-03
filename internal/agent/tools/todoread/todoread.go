package todoread

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/tmc/langchaingo/tools"
)

//go:embed description.txt
var toolDescription string

const ToolName = "TodoRead"

var _ tools.Tool = &TodoRead{}

type TodoRead struct {
	TodoList            *string
	RepositoryDirectory string
	Debug               bool
}

func (t *TodoRead) Name() string {
	return ToolName
}

func (t *TodoRead) Description() string {
	return toolDescription
}

func (t *TodoRead) Call(ctx context.Context, input string) (string, error) {
	if t.Debug {
		fmt.Printf("[todoread] %s\n", input)
	}
	return *t.TodoList, nil
}
