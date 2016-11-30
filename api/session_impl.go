package api

import (
	"fmt"
	"net/http"
)

const (
	DEFAULT_SUMO_ADDRESS = "https://api.sumologic.com/api/v1"
)

type SessionImpl struct {
	address   string
	accessID  string
	accessKey string
}

func NewSession() Session {
	s := &SessionImpl{}
	s.SetAddress(DEFAULT_SUMO_ADDRESS)
	return s
}

func (s *SessionImpl) SetAddress(address string) {
	s.address = address
}

func (s *SessionImpl) SetCredentials(accessID, accessKey string) {
	s.accessID = accessID
	s.accessKey = accessKey
}

func (s *SessionImpl) NewRequest(method string, endpoint string) (*http.Request, error) {
	uri := fmt.Sprintf("%s%s", s.address, endpoint)
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s.accessID, s.accessKey)
	return req, nil
}
