package agent

import (
	"coding-assistant/internal/agent/prompt"
	"coding-assistant/internal/agent/tools/find"
	"coding-assistant/internal/agent/tools/findfiles"
	"coding-assistant/internal/agent/tools/grep"
	"coding-assistant/internal/agent/tools/ls"
	"coding-assistant/internal/agent/tools/read_file"
	"coding-assistant/internal/agent/tools/searchtext"
	"coding-assistant/internal/agent/tools/todoread"
	"coding-assistant/internal/agent/tools/todowrite"
	"coding-assistant/internal/agent/tools/write_file"
	"context"
	"fmt"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"

	"coding-assistant/internal/model"
)

type Agent struct {
	llm                 llms.Model
	memory              schema.ChatMessageHistory
	repositoryDirectory string
	debug               bool
	todoList            string
}

func NewAgent(modelType model.Model, directory string, debug bool) (*Agent, error) {
	var llm llms.Model
	var err error

	switch modelType {
	case model.ModelGPT4_1:
		llm, err = openai.New()
	case model.ModelCodex:
		// NOTE: This is an assumption. Adjust if Codex requires a different client or model name.
		llm, err = openai.New(openai.WithModel("code-davinci-002")) // Example model name for Codex
	case model.ModelClaude3_5Sonnet:
		// NOTE: This is an assumption. Adjust if Claude 3.5 Sonnet requires specific Anthropic client configuration.
		llm, err = anthropic.New()
	case model.ModelGemini2_5Pro:
		llm, err = googleai.New(context.Background(), googleai.WithDefaultModel("gemini-2.5-pro"))
	default:
		return nil, fmt.Errorf("unsupported model type: %v", modelType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create LLM: %w", err)
	}

	return &Agent{
		llm:                 llm,
		memory:              memory.NewChatMessageHistory(),
		repositoryDirectory: directory,
		debug:               debug,
	}, nil
}

func (agent *Agent) ResetToto() {
	agent.todoList = ""
}

func (agent *Agent) Handle(input string) (*model.AgentResponse, error) {
	messages := make([]llms.MessageContent, 0)
	messages = append(messages, llms.MessageContent{
		Role: llms.ChatMessageTypeSystem,
		Parts: []llms.ContentPart{
			llms.TextContent{Text: prompt.GetSystemPrompt()},
		},
	})
	messages = append(messages, llms.MessageContent{
		Role: llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{
			llms.TextContent{Text: input},
		},
	})

	agentTools := make([]tools.Tool, 0)
	agentTools = append(agentTools, &find.Find{
		RepositoryDirectory: agent.repositoryDirectory,
		Debug:               agent.debug,
	})
	agentTools = append(agentTools, &read_file.ReadFile{
		RepositoryDirectory: agent.repositoryDirectory,
		Debug:               agent.debug,
	})
	agentTools = append(agentTools, &write_file.WriteFile{
		RepositoryDirectory: agent.repositoryDirectory,
		Debug:               agent.debug,
	})
	agentTools = append(agentTools, &grep.Grep{
		RepositoryDirectory: agent.repositoryDirectory,
		Debug:               agent.debug,
	})
	agentTools = append(agentTools, &ls.Ls{
		RepositoryDirectory: agent.repositoryDirectory,
		Debug:               agent.debug,
	})
	agentTools = append(agentTools, &findfiles.FindFiles{
		RepositoryDirectory: agent.repositoryDirectory,
		Debug:               agent.debug,
	})
	agentTools = append(agentTools, &todoread.TodoRead{
		TodoList:            &agent.todoList,
		RepositoryDirectory: agent.repositoryDirectory,
		Debug:               agent.debug,
	})
	agentTools = append(agentTools, &todowrite.TodoWrite{
		TodoList:            &agent.todoList,
		RepositoryDirectory: agent.repositoryDirectory,
		Debug:               agent.debug,
	})
	agentTools = append(agentTools, &searchtext.SearchText{
		RepositoryDirectory: agent.repositoryDirectory,
		Debug:               agent.debug,
	})

	a := agents.NewOneShotAgent(agent.llm, agentTools, agents.WithMaxIterations(50))
	executor := agents.NewExecutor(a)

	answer, err := chains.Run(context.Background(), executor, input)
	if err != nil {

		if agent.debug {
			fmt.Printf("agent failed to handle task")
		}
		return nil, err
	}

	return &model.AgentResponse{
		Response: answer,
	}, nil

}
