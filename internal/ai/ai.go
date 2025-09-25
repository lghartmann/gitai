package ai

import (
	"context"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"github.com/openai/openai-go/v2/packages/param"
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

func GenerateCommitMessage(diff string, status string) (string, error) {
	systemMessage := "You are a helpful assistant that generates concise and meaningful git commit messages based on the provided git diff and status."
	userMessage := "Generate a concise and meaningful git commit message based on the following git diff and status:\n\nGit Diff:\n" + diff + "\n\nGit Status:\n" + status
	maxTokens := param.NewOpt[int64](60)
	temperature := param.NewOpt(0.7)

	return CallLLM(systemMessage, userMessage, maxTokens, temperature)
}
