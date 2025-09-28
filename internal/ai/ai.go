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
	systemMessage := `You are an expert software engineer. Given a git diff and git status, produce one or more candidate commit messages following Conventional Commits format. Output only the commit message(s), nothing else. Rules:

Use types: feat, fix, docs, style, refactor, perf, test, chore.
Optionally include a scope in parentheses (e.g., feat(auth):).
Subject must be imperative, ≤ 50 characters (try to stay short).
If needed, include a 1-2 paragraph body (wrap at ~72 chars) that explains why the change was made and any important context or migration steps.`

	userMessage := "diff: " + diff + "\n\nstatus: " + status
	maxTokens := param.NewOpt[int64](60)
	temperature := param.NewOpt(0.7)

	return CallLLM(systemMessage, userMessage, maxTokens, temperature)
}
