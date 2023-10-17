package utils

import (
	"github.com/kwanok/minsim-consumer/client"
	"strings"
)

func Preprocess(text string) []string {
	words := strings.Fields(text)
	var newWords []string
	var chunks []string
	var currentChunk string

	for _, word := range words {
		if strings.HasPrefix(word, "@") && len(word) > 1 {
			newWords = append(newWords, "@user")
		} else if strings.HasPrefix(word, "http") {
			newWords = append(newWords, "http")
		} else {
			newWords = append(newWords, word)
		}
	}

	for _, newWord := range newWords {
		if len(currentChunk)+len(newWord)+1 > 300 {
			chunks = append(chunks, currentChunk)
			currentChunk = newWord
		} else {
			if currentChunk != "" {
				currentChunk += " "
			}
			currentChunk += newWord
		}
	}

	if currentChunk != "" {
		chunks = append(chunks, currentChunk)
	}

	return chunks
}

func Average(responses []*client.PredictResponse) *client.PredictResponse {
	result := &client.PredictResponse{
		Positive: 0.0,
		Negative: 0.0,
		Neutral:  0.0,
	}

	for _, response := range responses {
		result.Positive += response.Positive
		result.Negative += response.Negative
		result.Neutral += response.Neutral
	}

	result.Positive = result.Positive / float64(len(responses))
	result.Negative = result.Negative / float64(len(responses))
	result.Neutral = result.Neutral / float64(len(responses))

	return result
}
