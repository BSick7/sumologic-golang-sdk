package api

type Client struct {
	session Session
}

func NewClient(session Session) *Client {
	return &Client{
		session: session,
	}
}

