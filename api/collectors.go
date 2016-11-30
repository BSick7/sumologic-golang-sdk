package api

import (
	"net/url"
	"strconv"
)

type Collectors struct {
	executor *ClientExecutor
}

func NewCollectors(executor *ClientExecutor) *Collectors {
	return &Collectors{
		executor: executor,
	}
}

type Collector struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	CollectorType    string          `json:"collectorType"`
	Alive            bool            `json:"alive"`
	Links            []CollectorLink `json:"links"`
	CollectorVersion string          `json:"collectorVersion"`
	Ephemeral        bool            `json:"ephemeral"`
	Description      string          `json:"description"`
	OsName           string          `json:"osName"`
	OsArch           string          `json:"osArch"`
	OsVersion        string          `json:"osVersion"`
	Category         string          `json:"category"`
}

type CollectorLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func (c *Collectors) List(offset int, limit int) ([]*Collector, error) {
	params := url.Values{
		"offset": []string{strconv.Itoa(offset)},
		"limit":  []string{strconv.Itoa(limit)},
	}
	type listResponse struct {
		collectors []*Collector `json:"collectors"`
	}
	list := &listResponse{}
	if err := c.executor.Get("/collectors", params, list); err != nil {
		return nil, err
	}
	return list.collectors, nil
}
