package api

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

const OpenAIMaxTokens = 50

const OpenAISystemContent = `You want to help Renato to get hired and 
are very excited about this as you're speaking with the hiring manager. 
Always reply with less than 50 tokens.`

func getOpenAIMessage(body string) (string, error) {
	openAiApiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(openAiApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: OpenAISystemContent,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: body,
				},
			},
			Temperature: 0.1,
			MaxTokens:   OpenAIMaxTokens,
		},
	)
	if err != nil {
		log.Printf("ChatCompletion error: %v", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
