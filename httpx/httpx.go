package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Common Content Types
const (
	HeaderContentTypeJson = "application/json"
	HeaderContentTypeXml  = "application/xml"
)

// HttpClient is an http client interface
type HttpClient interface {
	Do(ctx context.Context, method string, url url.URL, header http.Header, body any) (response Response, err error)
}

type httpClient struct {
	client        http.Client
	retryAttempts int
	defaultHeader http.Header
}

// New initiates new http client
func New(client http.Client, retryAttempts int) HttpClient {
	return httpClient{
		client:        client,
		retryAttempts: retryAttempts,
	}
}

func (c httpClient) Do(ctx context.Context, method string, url url.URL, header http.Header, body any) (response Response, err error) {
	return c.retry(ctx, c.retryAttempts, func() (response Response, err error) {
		fullHeader := c.buildRequestHeader(header)
		if err != nil {
			return Response{}, err
		}

		reqBody, err := c.buildRequestBody(fullHeader.Get("Content-Type"), body)

		req, err := http.NewRequestWithContext(ctx, method, url.String(), bytes.NewReader(reqBody))
		if err != nil {
			return
		}

		req.Header = fullHeader

		res, err := c.client.Do(req)
		if err != nil {
			return
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return
		}
		defer res.Body.Close()

		if res.StatusCode >= http.StatusInternalServerError {
			err = fmt.Errorf("status code %d, resp %s", res.StatusCode, resBody)
			return
		}

		response = Response{Body: resBody, StatusCode: res.StatusCode}
		return
	})
}

// retry is a recursive function that makes up to a certain number of retries
func (c httpClient) retry(ctx context.Context, attempts int, f func() (response Response, err error)) (response Response, err error) {
	response, err = f()
	if err != nil {
		if attempts--; attempts > 0 {
			return c.retry(ctx, attempts, f)
		}
		return
	}
	return
}

func (c httpClient) buildRequestHeader(header http.Header) http.Header {
	fullHeader := make(http.Header, len(c.defaultHeader)+len(header))

	// Add default header
	for k, v := range c.defaultHeader {
		for _, item := range v {
			fullHeader.Add(k, item)
		}
	}

	// Add custom header
	for k, v := range header {
		for _, item := range v {
			fullHeader.Add(k, item)
		}
	}

	return fullHeader
}

func (c httpClient) buildRequestBody(contentType string, body any) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch contentType {
	case "application/xml":
		return xml.Marshal(body)
	case "application/json":
		fallthrough
	default:
		return json.Marshal(body)
	}
}

type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (r *Response) Bytes() []byte {
	return r.Body
}

func (r *Response) String() string {
	return string(r.Body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}
