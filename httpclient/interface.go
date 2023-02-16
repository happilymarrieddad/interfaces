package httpclient

import (
	"io"
	"net/http"
	"net/url"

	_ "github.com/golang/mock/mockgen/model"
)

//go:generate mockgen -destination=./mocks/HTTPClient.go -package=mocks github.com/happilymarrieddad/interfaces/httpclient HTTPClient
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Do(req *http.Request) (*http.Response, error)
}
