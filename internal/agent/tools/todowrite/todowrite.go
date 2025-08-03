package todowrite

import (
	"context"
	_ "embed"

	"github.com/tmc/langchaingo/tools"
)

//go:embed description.txt
var toolDescription string

const ToolName = "TodoWrite"

var _ tools.Tool = &TodoWrite{}

type TodoWrite struct {
	TodoList            *string
	RepositoryDirectory string
	Debug               bool
}

func (t *TodoWrite) Name() string {
	return ToolName
}

func (t *TodoWrite) Description() string {
	return toolDescription
}

func (t *TodoWrite) Call(ctx context.Context, input string) (string, error) {
	*t.TodoList = input
	return "TODO list updated.", nil
}
