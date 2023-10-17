package submission

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

type redditSubmissionConsumer struct {
	*config.Config
	*client.GrpcClient
	*client.HttpClient
}

func (c *redditSubmissionConsumer) Start(ctx context.Context) error {
	cs, err := _kafka.NewConsumer(&_kafka.ConfigMap{
		"bootstrap.servers": strings.Join(c.Kafka.Servers, ","),
		"group.id":          "reddit-group",
		"auto.offset.reset": "earliest",
	})

	defer cs.Close()

	if err != nil {
		return err
	}

	err = cs.SubscribeTopics([]string{c.Kafka.RedditSubmissionsTopic}, nil)

	if err != nil {
		return err
	}

	minsimClient := minsim.NewMinsimClient(c.GrpcClient.ClientConn)

	for {
		msg, err := cs.ReadMessage(-1)
		if err == nil {
			message := &kafka.RedditSubmissionMessage{}
			err := json.Unmarshal(msg.Value, message)
			if err != nil {
				log.Fatal("unmarshal: ", err)
			}

			submission := message.Payload.After

			chunks := utils.Preprocess(submission.Selftext)

			responses := make([]*client.PredictResponse, 0)

			for _, chunk := range chunks {
				resp, err := c.HttpClient.Predict(chunk)
				if err != nil {
					log.Fatal("predict: ", err)
				}
				responses = append(responses, resp)
			}

			if len(responses) == 0 {
				continue
			}

			resp := utils.Average(responses)

			_, err = minsimClient.NewMinsim(ctx, &minsim.NewMinsimRequest{
				Type:      "reddit-submission",
				Positive:  float32(resp.Positive),
				Negative:  float32(resp.Negative),
				Neutral:   float32(resp.Neutral),
				Content:   strings.Join(chunks, " "),
				User:      submission.Author,
				Id:        submission.Id,
				Url:       submission.Url,
				CreatedAt: int64(submission.CreatedUtc),
				Subreddit: submission.Subreddit,
			})

			if err != nil {
				log.Fatal("grpc: ", err)
			}
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func NewRedditSubmissionConsumer(
	config *config.Config,
	grpcClient *client.GrpcClient,
	httpClient *client.HttpClient,
) kafka.Consumer {
	return &redditSubmissionConsumer{
		Config:     config,
		GrpcClient: grpcClient,
		HttpClient: httpClient,
	}
}
