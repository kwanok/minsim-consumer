package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const retryCount = 100

type HttpClient struct {
	client          *http.Client
	inferenceApiUrl string
}

func NewHttpClient(
	client *http.Client,
	inferenceApiUrl string,
) *HttpClient {
	return &HttpClient{
		client:          client,
		inferenceApiUrl: inferenceApiUrl,
	}
}

type PredictRequest struct {
	Text string `json:"text"`
}

type PredictResponse struct {
	Negative float64 `json:"negative"`
	Positive float64 `json:"positive"`
	Neutral  float64 `json:"neutral"`
}

func (c *HttpClient) Predict(
	text string,
) (*PredictResponse, error) {
	body := PredictRequest{
		Text: text,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		c.inferenceApiUrl+"/predict",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response

	for i := 0; i < retryCount; i++ {
		resp, err = c.client.Do(req)
		if err == nil {
			break
		}

		log.Println("retrying predict request", i)
		time.Sleep(128 * time.Millisecond * time.Duration(i))
	}

	if err != nil {
		log.Fatal("predict: ", err)

		return nil, err
	}

	defer resp.Body.Close()

	response := PredictResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)

	if err != nil {
		fmt.Println(text)
		fmt.Println(resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		log.Fatal(string(body))

		return nil, err
	}

	return &response, nil
}
