package prompt

import _ "embed"

//go:embed txt/system.txt
var systemPrompt string

func GetSystemPrompt() string {
	return systemPrompt
}
