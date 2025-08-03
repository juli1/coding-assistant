package todowrite

import (
	"context"
	_ "embed"
	"fmt"

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
	if t.Debug {
		fmt.Printf("[todowrite] writing the following to-do list\n---\n%s\n---\n", input)
	}
	return "to-do list updated.", nil
}
