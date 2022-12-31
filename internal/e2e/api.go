package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const host = "http://localhost"

var jar, _ = cookiejar.New(nil)
var client = &http.Client{Jar: jar}

type getAPI[R any] struct {
	url url.URL
}

func NewGetAPI[R any](path string) getAPI[R] {
	url := url.URL{}
	url.Scheme = "http"
	url.Host = "localhost"
	url.Path = path
	return getAPI[R]{url}
}

func (api getAPI[R]) getURL() string {
	return api.url.String()
}

func (api getAPI[R]) Request() (*R, error) {
	resp, err := client.Get(api.getURL())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("wrong status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r R
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type postAPI[B any, R any] struct {
	path string
}

func NewPostAPI[B any, R any](path string) postAPI[B, R] {
	return postAPI[B, R]{path}
}

func (api postAPI[B, R]) Request(b B) (*R, error) {
	bs, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(host+api.path, "application/json", bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("wrong status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil
	}
	var r R
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
