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
	return http.NewRequest(method, uri, nil)
}

func (s *SessionImpl) CreateTransport() http.RoundTripper {
	return NewAnonymousTransport(func(req *http.Request) (*http.Response, error) {
		req.SetBasicAuth(s.accessID, s.accessKey)
		return http.DefaultTransport.RoundTrip(req)
	})
}
