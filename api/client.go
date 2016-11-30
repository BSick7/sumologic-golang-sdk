package api

import (
	"net/http"
)

type Client struct {
	executor *ClientExecutor
}

func NewClient(session Session) *Client {
	return &Client{
		executor: NewClientExecutor(session, http.DefaultClient),
	}
}

func (c *Client) Collectors() *Collectors {
	return NewCollectors(c.executor)
}
