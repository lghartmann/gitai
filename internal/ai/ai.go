package ai

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"github.com/openai/openai-go/v2/packages/param"
	"github.com/sourcegraph/go-diff/diff"
)

func CallLLM(systemMessage string, userMessage string, maxTokens param.Opt[int64], temperature param.Opt[float64]) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		return "", ErrAPIKeyNotSet
	}

	client := openai.NewClient(option.WithAPIKey(apiKey))

	res, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT3_5Turbo0125,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemMessage),
			openai.UserMessage(userMessage),
		},
		MaxTokens:   maxTokens,
		Temperature: temperature,
	})

	if err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", ErrNoResponse
	}

	return res.Choices[0].Message.Content, nil

}

func GenerateCommitMessage(diff string, status string, detailed bool) (string, error) {
	var systemMessage string
	if detailed {
		systemMessage = "You are an expert software engineer specializing in writing clear, concise, and professional git commit messages. Given a git diff and status, generate a detailed commit message that accurately summarizes the changes. If possible, include bullet points for each significant change. Strictly output only the commit message itself, without any explanations, formatting, or additional text."
	} else {
		systemMessage = "You are an expert software engineer specializing in writing clear, concise, and professional git commit messages. Given a git diff and status, generate a commit message that accurately summarizes the changes. Strictly output only the commit message itself, without any explanations, formatting, or additional text."
	}
	userMessage := "diff: " + diff + "\n\nstatus: " + status
	maxTokens := param.NewOpt[int64](60)
	temperature := param.NewOpt(0.7)

	err := checkDiffSafety(diff)
	if err != nil {
		return "", err
	}

	return CallLLM(systemMessage, userMessage, maxTokens, temperature)
}

func checkDiffSafety(diffText string) (err error) {
	fileDiffs, err := diff.ParseMultiFileDiff([]byte(diffText))
	if err != nil {
		fmt.Printf("Error parsing diff: %v\n", err)
		return
	}

fileDiffLoop:
	for _, fileDiff := range fileDiffs {
		for _, hunk := range fileDiff.Hunks {
			for _, line := range hunk.Body {
				if strings.HasPrefix(string(line), "+") && !strings.HasPrefix(string(line), "+++") {
					content := strings.TrimPrefix(string(line), "+")

					if containsSensitiveData(content) {
						warning := fmt.Sprintf("⚠️  WARNING: Potential sensitive data detected in %s: %s\n",
							fileDiff.NewName, strings.TrimSpace(content))

						fmt.Println(warning)

						var option string
						fmt.Print("Are you sure you want to continue? [Y/n]")
						fmt.Scanln(&option)

						switch strings.ToLower(option) {
						case "", "y", "yes":
							continue
						case "n", "no":
							err = errors.New(warning)
							break fileDiffLoop
						}
					}
				}
			}
		}
	}
	return
}

func containsSensitiveData(content string) bool {
	lower := strings.ToLower(content)
	fmt.Printf("DEBUG: Checking if sensitive: %q\n", lower)

	sensitivePatterns := []string{
		"password", "passwd", "pwd",
		"api_key", "apikey", "api-key",
		"secret", "token", "auth",
		"private_key", "privatekey", "private-key",
		"access_key", "accesskey", "access-key",
		"client_secret", "clientsecret", "client-secret",
		"bearer", "oauth",
	}

	for _, pattern := range sensitivePatterns {
		if strings.Contains(lower, pattern) {
			return true
		}
	}

	fmt.Println("DEBUG: No sensitive patterns found")
	return false
}
