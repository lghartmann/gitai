# Gitai 🤖✨

Your AI-Powered Git Companion

---

## What is Gitai?

**Gitai** is a modern CLI tool that leverages AI to help you write clear, concise, and professional git commit messages based on your repository’s changes.  
No more writer’s block or vague commit messages—let AI do the heavy lifting!

---

## Features 🚀

- **AI-Generated Commit Messages:**  
  Generate meaningful commit messages from your git diff and status.
- **Detailed or Concise:**  
  Use `--detailed` for a more comprehensive message.
- **Automatic Commit & Staging:**  
  Use `--commit` to commit instantly, and `--add` to stage all changes before committing.
- **Interactive Loader:**  
  Enjoy a smooth CLI experience with a live spinner while AI works.

---

## Planned Features 🛠️

- **Conventional Commit Support** 📝  
  Option to generate messages in [Conventional Commits](https://www.conventionalcommits.org/) format.
- **Branch Name Suggestions** 🌿  
  Let AI suggest branch names based on your changes.
- **Pre-commit Hook Integration** 🪝  
  Seamlessly integrate Gitai as a git pre-commit hook.
- **Commit Message Editing** ✏️  
  Approve or edit the AI-generated message before committing.
- **Multi-language Support** 🌐  
  Generate commit messages in different languages.
- **Summary/Explanation Mode** 📄  
  Summarize code changes or explain diffs in plain English.
- **History and Undo** ⏪  
  Show a history of generated messages and allow undoing the last commit.
- **Custom AI Prompts** 🛠️  
  Customize the prompt sent to the AI for tailored messages.
- **Integration with Issue Trackers** 🔗  
  Automatically reference issue numbers or pull request IDs in messages.
- **Quality Checks** ✅  
  Lint or check the generated message for length, clarity, or forbidden words.
- **Batch Mode** 📦  
  Generate commit messages for multiple commits (e.g., for rebasing or squashing).
- **Config File Support** ⚙️  
  Allow user configuration via a `.gitai.yaml` or similar file.
- **Stats and Analytics** 📊  
  Show stats about commit message usage, length, or AI performance.

---

## Installation 🏗️

### 1. Build from Source

```sh
git clone https://github.com/yourusername/gitai.git
cd gitai
make install
```

This will build and move the `gitai` binary to `/usr/local/bin/`.

### 2. Or Add to Your PATH

```sh
export PATH="$PATH:/path/to/gitai/bin"
```

---

## Usage 🧑‍💻

```sh
gitai gen commit_message [flags]
```

### Common Flags

- `--detailed`  
  Generate a more detailed commit message.
- `--commit`  
  Commit with the generated message.
- `--add`  
  Stage all changes before committing.

### Examples

```sh
gitai gen commit_message --detailed
gitai gen commit_message --commit
gitai gen commit_message --add --commit
```

---

## Help

See all commands and options:

```sh
gitai --help
gitai gen commit_message --help
```

---

## Contributing 🤝

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

---

## License

MIT © Vusal Huseynov
