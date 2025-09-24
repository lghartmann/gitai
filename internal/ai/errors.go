package ai

import "errors"

var (
	ErrAPIKeyNotSet = errors.New("OPENAI_API_KEY not set")
	ErrNoResponse   = errors.New("no response from OpenAI")
)
