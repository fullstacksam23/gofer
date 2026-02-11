# Gofer

A lightweight, autonomous CLI agent written in Go. This tool leverages LLMs (via OpenRouter/OpenAI-compatible APIs) to perform system tasks. It allows the AI to **read files**, **write files**, and **execute shell commands** to fulfill your prompts.

## 🚀 Features

- **Function Calling:** Implements the OpenAI Tool Use standard to let the LLM decide when to act.
- **File System Access:**
  - `Read`: Allows the LLM to inspect file contents.
  - `Write`: Allows the LLM to create or overwrite files (automatically creates directories).
- **Shell Execution:**
  - `Bash`: Gives the LLM access to the system shell to run commands.
- **OpenRouter Integration:** Pre-configured to use `anthropic/claude-haiku-4.5` via OpenRouter, but compatible with any OpenAI-compliant endpoint.

## ⚠️ Security Warning

**Use with caution.** This tool grants an AI model the ability to execute `bash` commands and modify your file system.

- Do not run this with root/sudo privileges.
- Review the code before running it against sensitive production environments.
- The AI could theoretically run destructive commands (e.g., `rm -rf`) if prompted or hallucinated.

## 🛠️ Prerequisites

- [Go](https://go.dev/dl/) (1.21 or higher recommended)
- An API Key from [OpenRouter](https://openrouter.ai/) (or an OpenAI-compatible provider).

## 📦 Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/fullstacksam23/gofer.git
   cd gofer
   ```
2. **Install dependencies:**

   ```bash
   go mod tidy

   ```

3. **Create .env:** This tool uses Go's embed feature to bake your configuration into the binary, allowing you to run it from any directory. Create a .env file in the root of the project:

   ```bash
   touch .env
   ```

4. **Add your API key**:

   .env

   ```bash
   # Required
   OPENROUTER_API_KEY=sk-or-v1-your-key-here

   # Optional (Defaults to OpenRouter)
   OPENROUTER_BASE_URL=[https://openrouter.ai/api/v1]
   ```

   Note: The model is currently hardcoded to anthropic/claude-haiku-4.5 in main.go. You can modify the Model field in the code to use other models.

5. **Compile and install the binary:**
   ```bash
   go install .
   ```

## 🏃 Usage

Once installed, you can run the tool from any terminal window in any folder.

**Example**

```bash
gofer -p "Add a comment above the readFileTool function in main.go explaining what it does"
```

## 🧩 How it Works

The application runs a loop:

1. Sends the user prompt to the LLM.
2. Checks if the LLM wants to call a tool (Read, Write, Bash).
3. If yes: Executes the Go function corresponding to the tool, feeds the result back to the LLM, and repeats the loop.
4. If no: Prints the final text response from the LLM and exits.
