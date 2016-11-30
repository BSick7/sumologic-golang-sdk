package api

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	DEFAULT_SUMO_ADDRESS = "https://api.sumologic.com/api/v1"
)

type SessionImpl struct {
	address   string
	accessID  string
	accessKey string
}

func DefaultSession() *SessionImpl {
	s := &SessionImpl{}
	s.SetAddress(DEFAULT_SUMO_ADDRESS)
	s.SetCredentials(os.Getenv("SUMO_ACCESS_ID"), os.Getenv("SUMO_ACCESS_KEY"))
	return s
}

func NewSession() *SessionImpl {
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

func (s *SessionImpl) NewRequest(method string, endpoint string, params url.Values) (*http.Request, error) {
	uri := fmt.Sprintf("%s%s?%s", s.address, endpoint, params.Encode())
	uri = strings.TrimRight(uri, "?")
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s.accessID, s.accessKey)
	return req, nil
}
