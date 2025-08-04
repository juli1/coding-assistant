package model

import "fmt"

type Model int

const (
	ModelUnknown Model = iota
	ModelGPT4_1
	ModelClaude3_5Sonnet
	ModelGemini2_5Pro
	ModelOllama = 5
)

var AllModels []Model = []Model{
	ModelGPT4_1, ModelClaude3_5Sonnet, ModelGemini2_5Pro, ModelOllama,
}

func (m Model) String() string {
	switch m {
	case ModelGPT4_1:
		return "gpt-4.1"
	case ModelClaude3_5Sonnet:
		return "claude-3.5-sonnet"
	case ModelGemini2_5Pro:
		return "gemini-2.5-pro"
	case ModelOllama:
		return "ollama"

	default:
		return "Unknown"
	}
}

func FromString(s string) (Model, error) {
	switch s {
	case "gpt-4.1":
		return ModelGPT4_1, nil
	case "claude-3.5-sonnet":
		return ModelClaude3_5Sonnet, nil
	case "gemini-2.5-pro":
		return ModelGemini2_5Pro, nil
	case "ollama":
		return ModelOllama, nil
	default:
		return ModelUnknown, fmt.Errorf("unknown model: %s", s)
	}
}
