# Contributing to Coding Assistant

First off, thank you for considering contributing to Coding Assistant! It's people like you that make open source software such a great community.

We welcome any type of contribution, not just code. You can help with:
*   **Reporting a bug**
*   **Discussing the current state of the code**
*   **Submitting a fix**
*   **Proposing new features**
*   **Becoming a maintainer**

## Getting Started

1.  **Fork the repository** on GitHub.
2.  **Clone your fork** locally:
    ```sh
    git clone https://github.com/YOUR_USERNAME/coding-assistant.git
    ```
3.  **Navigate to the project directory**:
    ```sh
    cd coding-assistant
    ```
4.  **Install dependencies**:
    This project uses Go modules. To install the necessary dependencies, run the following command:
    ```sh
    go mod tidy
    ```

## How to Contribute

### Proposing a Change

If you have an idea for a new feature or want to report a bug, please open an issue on GitHub. This allows for discussion and ensures that your work is not in vain.

### Making Changes

1.  **Create a new branch** for your changes. Use a descriptive name, like `feat/add-new-model` or `fix/improve-error-handling`.
    ```sh
    git checkout -b your-branch-name
    ```
2.  **Make your changes** to the code.
3.  **Ensure your code is formatted** according to Go standards. You can use `go fmt` to format your code:
    ```sh
    go fmt ./...
    ```
4.  **Commit your changes** with a clear and concise commit message.
    ```sh
    git commit -m "feat: Add support for the new XYZ model"
    ```
5.  **Push your changes** to your fork on GitHub.
    ```sh
    git push origin your-branch-name
    ```

### Submitting a Pull Request

Once you've pushed your changes, open a pull request from your fork to the main `coding-assistant` repository. In the pull request description, please explain the changes you've made and link to any relevant issues.

## Code Style

Please try to match the existing code style. This project follows the standard Go conventions.

## Questions?

If you have any questions, feel free to open an issue on GitHub.

Thank you for your contribution!
