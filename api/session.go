package api

import "net/http"

type Session interface {
	SetAddress(address string)
	SetCredentials(accessID, accessKey string)
	NewRequest(method string, endpoint string) (*http.Request, error)
}
