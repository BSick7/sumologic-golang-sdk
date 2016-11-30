package api

import (
	"net/http"
	"net/url"
)

type Session interface {
	SetAddress(address string)
	SetCredentials(accessID, accessKey string)
	NewRequest(method string, endpoint string, params url.Values) (*http.Request, error)
}
