package api

import (
	"net/http"
)

type Client struct {
	executor *ClientExecutor
}

func NewClient(session Session) *Client {
	return &Client{
		executor: NewClientExecutor(session, createHttpClient(session)),
	}
}

func createHttpClient(session Session) *http.Client {
	return &http.Client{
		Transport: session.CreateTransport(),
	}
}

func (c *Client) Collectors() *Collectors {
	return NewCollectors(c.executor)
}
