package utils

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient interface {
	ReadBytes(ctx context.Context, url string) ([]byte, error)
	IsAccessible(ctx context.Context, url string) bool
}

type httpClient struct{}

func NewHttpClient() HttpClient {
	return &httpClient{}
}

func (f *httpClient) ReadBytes(ctx context.Context, url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (f *httpClient) IsAccessible(ctx context.Context, url string) bool {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false
	}
	resp, err := http.DefaultClient.Do(req)
	return err == nil && 200 <= resp.StatusCode && resp.StatusCode < 300
}
