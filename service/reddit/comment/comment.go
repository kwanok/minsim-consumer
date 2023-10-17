package comment

import (
	"context"
	"encoding/json"
	_kafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/kwanok/minsim-consumer/client"
	"github.com/kwanok/minsim-consumer/config"
	"github.com/kwanok/minsim-consumer/kafka"
	"github.com/kwanok/minsim-consumer/service/minsim"
	"github.com/kwanok/minsim-consumer/utils"
	"log"
	"strings"
)

type redditCommentConsumer struct {
	*config.Config
	minsim.MinsimClient
	*client.HttpClient
}

func (c *redditCommentConsumer) Start(ctx context.Context) error {
	cs, err := _kafka.NewConsumer(&_kafka.ConfigMap{
		"bootstrap.servers": strings.Join(c.Kafka.Servers, ","),
		"group.id":          "reddit-group",
		"auto.offset.reset": "earliest",
	})

	defer cs.Close()

	if err != nil {
		return err
	}

	err = cs.SubscribeTopics([]string{c.Kafka.RedditCommentsTopic}, nil)

	if err != nil {
		return err
	}

	messageChannel := make(chan *kafka.RedditCommentMessage)

	go func() {
		for {
			msg, err := cs.ReadMessage(-1)
			if err == nil {
				message := &kafka.RedditCommentMessage{}
				err := json.Unmarshal(msg.Value, message)
				if err != nil {
					log.Fatal("unmarshal: ", err)
				}
				messageChannel <- message
			} else {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}()

	for {
		select {
		case message := <-messageChannel:
			go func() {
				comment := message.Payload.After

				chunks := utils.Preprocess(comment.Body)

				responses := make([]*client.PredictResponse, 0)

				for _, chunk := range chunks {
					resp, err := c.HttpClient.Predict(chunk)
					if err != nil {
						log.Fatal("predict: ", err)
					}
					responses = append(responses, resp)
				}

				if len(responses) == 0 {
					return
				}

				resp := utils.Average(responses)

				_, err = c.MinsimClient.NewMinsim(ctx, &minsim.NewMinsimRequest{
					Type:      "reddit-comment",
					Positive:  float32(resp.Positive),
					Negative:  float32(resp.Negative),
					Neutral:   float32(resp.Neutral),
					Content:   strings.Join(chunks, " "),
					User:      comment.Author,
					Url:       comment.Permalink,
					Id:        comment.Id,
					CreatedAt: int64(comment.CreatedUtc),
					Subreddit: comment.Subreddit,
				})
			}()
		case <-ctx.Done():
			return nil
		}
	}
}

func NewRedditCommentConsumer(
	config *config.Config,
	minsimClient minsim.MinsimClient,
	httpClient *client.HttpClient,
) kafka.Consumer {
	return &redditCommentConsumer{
		Config:       config,
		MinsimClient: minsimClient,
		HttpClient:   httpClient,
	}
}
