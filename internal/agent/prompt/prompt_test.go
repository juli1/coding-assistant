package prompt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSystemPrompt(t *testing.T) {
	prompt := GetSystemPrompt()
	assert.Contains(t, prompt, "You are a coding assistant CLI tool")
}
