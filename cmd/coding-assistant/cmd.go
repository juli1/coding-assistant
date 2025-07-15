package main

import (
	"bufio"
	"coding-assistant/internal/cli"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"coding-assistant/internal/agent"
	modelModule "coding-assistant/internal/model"
)

var (
	model string
	debug bool
)

var rootCmd = &cobra.Command{
	Use:   "coding-assistant [directory]",
	Short: "A CLI tool for coding assistance.",
	Long:  `A CLI tool that takes a directory as an argument and provides coding assistance using various models.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runAssistant(args[0], model, debug)
	},
}

func runAssistant(directory string, modelName string, debugMode bool) {
	modelType := modelModule.ParseModel(modelName)
	if modelType == modelModule.ModelUnknown {
		fmt.Printf("Invalid model: %s. Allowed models are: %s\n", modelName, modelModule.AllModels)
		os.Exit(1)
	}

	codingAgent, err := agent.NewAgent(modelType, directory, debugMode)
	if err != nil {
		fmt.Printf("Failed to create agent: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(cli.GetWelcomeMessage())

	if debugMode {
		fmt.Printf("Directory: %s\n", directory)
		fmt.Printf("Using model: %s\n", modelName)
		if debugMode {
			fmt.Println("Debug mode is enabled.")
		}
		fmt.Println("Type 'exit' or 'bye' to quit.")

	}
	fmt.Println("---------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		trimmedInput := strings.TrimSpace(input)

		if trimmedInput == "exit" || trimmedInput == "bye" {
			fmt.Println("Goodbye!")
			break
		}

		// Process the input here
		fmt.Printf("Processing: %s\n", input)
		response, err := codingAgent.Handle(input)
		if err != nil {
			fmt.Printf("Failed to handle request: %v\n", err)
			continue
		}

		fmt.Printf("Agent: %s\n", response.Response)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func init() {
	rootCmd.Flags().StringVar(&model, "model", "gpt-4.1", "The model to use for coding assistance (gpt-4.1, codex, claude-3.5-sonnet, gemini-2.5-pro).")
	rootCmd.Flags().BoolVar(&debug, "debug", false, "Enable debug mode.")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
