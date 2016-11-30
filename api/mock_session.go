package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type MockSession struct {
	Session
	address   string
	accessID  string
	accessKey string
	server    *httptest.Server
	handler   http.HandlerFunc
}

func NewMockSession() *MockSession {
	s := &MockSession{
		handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}
	s.SetAddress(DEFAULT_SUMO_ADDRESS)
	s.SetCredentials("mockaccessid", "mockaccesskey")

	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.handler(w, r)
	}))

	return s
}

func (s *MockSession) Handle(handler http.HandlerFunc) {
	s.handler = handler
}

func (s *MockSession) SetAddress(address string) {
	base, _ := url.Parse(address)
	s.address = fmt.Sprintf("%s%s", s.server.URL, base.Path)
}

func (s *MockSession) SetCredentials(accessID, accessKey string) {
	s.accessID = accessID
	s.accessKey = accessKey
}

func (s *MockSession) NewRequest(method string, endpoint string) (*http.Request, error) {
	uri := fmt.Sprintf("%s%s", s.address, endpoint)
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s.accessID, s.accessKey)
	return req, nil
}
