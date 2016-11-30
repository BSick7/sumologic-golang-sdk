package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"net/http/httputil"
	"os"
	"log"
)

type ClientExecutor struct {
	session Session
	client  *http.Client
	debug   bool
}

func NewClientExecutor(session Session, client *http.Client) *ClientExecutor {
	return &ClientExecutor{
		session: session,
		client:  client,
		debug:   os.Getenv("SUMO_DEBUG") == "1",
	}
}

func (c *ClientExecutor) Get(endpoint string, params url.Values, out interface{}) error {
	req, err := c.session.NewRequest("GET", endpoint, params)
	if err != nil {
		return fmt.Errorf("error creating request %s: %s", req.URL, err)
	}
	if c.debug {
		raw, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(raw))
	}
	res, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error requesting %s: %s", req.URL, err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("error response %d: %s", res.StatusCode, res.Status)
	}

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %s", err)
	}
	fmt.Println(string(raw))

	return json.Unmarshal(raw, out)
}
