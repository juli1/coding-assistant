# Simple Coding Assistant

A simple coding assistant written in Go. Totally experimental and work in progress.

## How to Run

There are two ways to run the application:

### 1. Build and Run

First, build the application:
```sh
go build -o coding-assistant cmd/coding-assistant/cmd.go
```

Then, run the compiled binary:
```sh
./coding-assistant [directory] --model [model_name] --debug
```

**Example:**
```sh
./coding-assistant . --model claude-3.5-sonnet --debug
```

### 2. Run directly with `go run`

You can also run the application directly without building it first:
```sh
go run cmd/coding-assistant/cmd.go [directory] --model [model_name] --debug
```

**Example:**
```sh
go run cmd/coding-assistant/cmd.go . --model claude-3.5-sonnet --debug
```

### Flags

*   `--model`: The model to use for coding assistance. Allowed values are `gpt-4.1`, `codex`, and `claude-3.5-sonnet`.
*   `--debug`: Enable debug mode (optional).
