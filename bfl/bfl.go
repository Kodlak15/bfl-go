package bfl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Key     string
	BaseURL string
}

func NewClient(key string, baseURL string) *Client {
	return &Client{
		Key:     key,
		BaseURL: baseURL,
	}
}

type AsyncResponse struct {
	ID         string `json:"id"`
	PollingURL string `json:"polling_url"`
}

type AsyncWebhookResponse struct {
	ID         string `json:"id"`
	PollingURL string `json:"polling_url"`
	WebhookURL string `json:"webhook_url"`
}

type HTTPValidationError struct {
	Detail []ValidationError `json:"detail"`
}

func (e *HTTPValidationError) Error() string {
	n := len(e.Detail)
	if n == 0 {
		return "Validation error"
	}
	msg := ""
	for i, detail := range e.Detail {
		msg += fmt.Sprintf("Validation error (%d/%d): %v", i+1, n, detail.Msg)
		if i != n-1 {
			msg += "\n"
		}
	}
	return msg
}

type Result interface {
	*GenerateResult | *FinetuneResult
}

type Details interface {
	*GenerateDetails | *FinetuneDetails
}

type ResultResponse[T Result, D Details] struct {
	ID       string         `json:"id"`
	Status   StatusResponse `json:"status"`
	Result   T              `json:"result"`
	Progress float64        `json:"progress"`
	Details  D              `json:"details"`
}

type StatusResponse string

const (
	StatusTaskNotFound     StatusResponse = "Task not found"
	StatusPending          StatusResponse = "Pending"
	StatusRequestModerated StatusResponse = "Request Moderated"
	StatusContentModerated StatusResponse = "Content Moderated"
	StatusReady            StatusResponse = "Ready"
	StatusError            StatusResponse = "Error"
)

type ValidationError struct {
	Loc []interface{} `json:"loc"`
	Msg string        `json:"msg"`
	Typ string        `json:"type"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %v", e.Msg)
}

type AsyncTask interface {
	GetActionURL(baseURL string) string
}

func (c *Client) AsyncRequest(url string, inputs AsyncTask) (*AsyncResponse, error) {
	if c.Key == "" {
		return nil, fmt.Errorf("API key is not set")
	}
	data, err := json.Marshal(inputs)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Key", c.Key)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	switch res.StatusCode {
	case 200:
		var ar AsyncResponse
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}
		return &ar, nil
	case 422:
		var httpValidationError HTTPValidationError
		if err = json.Unmarshal(body, &httpValidationError); err != nil {
			return nil, err
		}
		return nil, &httpValidationError
	default:
		return nil, fmt.Errorf("status code: %d, body: %s", res.StatusCode, string(body))
	}
}

func GetResult[T Result, D Details](client *Client, taskID string) (*ResultResponse[T, D], error) {
	url := fmt.Sprintf("%s/v1/get_result?id=%s", client.BaseURL, taskID)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	switch res.StatusCode {
	case 200:
		var resultResponse ResultResponse[T, D]
		err = json.Unmarshal(body, &resultResponse)
		if err != nil {
			return nil, err
		}
		return &resultResponse, nil
	case 422:
		var httpValidationError HTTPValidationError
		if err = json.Unmarshal(body, &httpValidationError); err != nil {
			return nil, err
		}
		return nil, &httpValidationError
	default:
		return nil, fmt.Errorf("status code: %d, body: %s", res.StatusCode, string(body))
	}
}

// Poll the BFL API for the result of an async task every second.
func Poll[T Result, D Details](client *Client, ar *AsyncResponse, verbose bool) (*ResultResponse[T, D], error) {
	sleepTimeSeconds := 1
	attempts := 0
	for {
		res, err := http.Get(ar.PollingURL)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		switch res.StatusCode {
		case 200:
			var resultResponse ResultResponse[T, D]
			err = json.Unmarshal(body, &resultResponse)
			if err != nil {
				return nil, err
			}
			if resultResponse.Status == StatusReady {
				return &resultResponse, nil
			}
		case 422:
			var httpValidationError HTTPValidationError
			if err = json.Unmarshal(body, &httpValidationError); err != nil {
				return nil, err
			}
			return nil, &httpValidationError
		default:
			return nil, fmt.Errorf("status code: %d, body: %s", res.StatusCode, string(body))
		}
		if verbose {
			if attempts%10 == 0 {
				fmt.Printf("Polling for result... (Wait time: %d seconds)\n", sleepTimeSeconds*attempts)
			}
		}
		time.Sleep(time.Duration(sleepTimeSeconds) * time.Second)
		attempts++
	}
}
