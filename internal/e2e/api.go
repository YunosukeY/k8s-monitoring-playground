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

func NewGetAPIWithUser[R any](username string, password string, host string, path string) getAPI[R] {
	u := url.URL{}
	u.Scheme = "http"
	u.User = url.UserPassword(username, password)
	u.Host = host
	u.Path = path
	return getAPI[R]{url: u}
}

func (api getAPI[R]) Request() (*R, error) {
	return api.RequestWithParam("")
}

func (api getAPI[R]) RequestWithParam(query string) (*R, error) {
	api.url.RawQuery = query
	defer func() {
		api.url.RawQuery = ""
	}()
	resp, err := client.Get(api.url.String())
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
	url url.URL
}

func NewPostAPIWithUser[B any, R any](username string, password string, host string, path string) postAPI[B, R] {
	u := url.URL{}
	u.Scheme = "http"
	u.User = url.UserPassword(username, password)
	u.Host = host
	u.Path = path
	return postAPI[B, R]{url: u}
}

func (api postAPI[B, R]) Request(b B) (*R, error) {
	bs, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(api.url.String(), "application/json", bytes.NewBuffer(bs))
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
