package model

type Model int

const (
	ModelUnknown Model = iota
	ModelGPT4_1
	ModelCodex
	ModelClaude3_5Sonnet
	ModelGemini2_5Pro
)

func (m Model) String() string {
	switch m {
	case ModelGPT4_1:
		return "gpt-4.1"
	case ModelCodex:
		return "codex"
	case ModelClaude3_5Sonnet:
		return "claude-3.5-sonnet"
	case ModelGemini2_5Pro:
		return "gemini-2.5-pro"
	default:
		return "Unknown"
	}
}

func ParseModel(s string) Model {
	switch s {
	case "gpt-4.1":
		return ModelGPT4_1
	case "codex":
		return ModelCodex
	case "claude-3.5-sonnet":
		return ModelClaude3_5Sonnet
	case "gemini-2.5-pro":
		return ModelGemini2_5Pro
	default:
		return ModelUnknown
	}
}
