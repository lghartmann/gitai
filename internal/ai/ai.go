package ai

import (
	"context"
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

	return CallLLM(systemMessage, userMessage, maxTokens, temperature)
}

func CheckDiffSafety(diffText string) (sensitiveData []string, err error) {
	fileDiffs, err := diff.ParseMultiFileDiff([]byte(diffText))
	if err != nil {
		return nil, nil
	}

	var foundSensitive []string

	for _, fileDiff := range fileDiffs {
		for _, hunk := range fileDiff.Hunks {
			lines := strings.Split(string(hunk.Body), "\n")

			newLineNum := int(hunk.NewStartLine)

			for _, line := range lines {
				if line == "" {
					continue
				}

				if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
					content := strings.TrimPrefix(line, "+")

					if containsSensitiveData(content) {
						relativeFile := strings.TrimPrefix(fileDiff.NewName, "b/")
						relativeFile = strings.TrimPrefix(relativeFile, "a/")
						clickableFile := fmt.Sprintf("%s:%d:1", relativeFile, newLineNum)
						sensitive := fmt.Sprintf("%s: %s", clickableFile, strings.TrimSpace(content))
						foundSensitive = append(foundSensitive, sensitive)
					}
					newLineNum++
				} else if strings.HasPrefix(line, " ") {
					newLineNum++
				}
			}
		}
	}
	return foundSensitive, nil
}

func containsSensitiveData(content string) (isSensitive bool) {
	lower := strings.ToLower(content)

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
			isSensitive = true
			return
		}
	}

	return
}
