package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/spf13/cobra"

	"coding-assistant/internal/agent"
	"coding-assistant/internal/cli"
	"coding-assistant/internal/model"
)

var (
	debug     bool
	modelType string
)

var rootCmd = &cobra.Command{
	Use:   "coding-assistant",
	Short: "A powerful and flexible coding assistant",
	Long:  cli.GetWelcomeMessage(),
	Run:   run,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")
	rootCmd.PersistentFlags().StringVarP(&modelType, "model", "m", "gemini-2.5-pro", "Specify the model to use (e.g., gpt-4.1, codex, claude-3.5-sonnet, gemini-2.5-pro")
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println(cli.GetWelcomeMessage())
	directory := args[0]
	parsedModel, err := model.FromString(modelType)
	if err != nil {
		log.Fatalf("Failed to parse model type: %v", err)
	}

	a, err := agent.NewAgent(parsedModel, directory, debug)
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	if debug {
		fmt.Printf("Directory: %s\n", directory)
		fmt.Printf("Using model: %s\n", modelType)
		if debug {
			fmt.Println("Debug mode is enabled.")
		}
		fmt.Println("Type 'exit' or 'bye' to quit.")
	}

	// Create a temporary file in the default temp directory
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	tmpFilename := tmpFile.Name()
	defer func() { _ = os.Remove(tmpFilename) }()

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		HistoryFile:     "/tmp/coding-assistant.tmp",
		HistoryLimit:    50,
		AutoComplete:    nil,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatalf("Failed to create readline: %v", err)
	}
	defer func() { _ = rl.Close() }()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF, readline.ErrInterrupt
			break
		}

		if strings.EqualFold(line, "exit") {
			println("kthxbye")
			break
		}

		response, err := a.Handle(line)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Println(response.Response)
	}
}
