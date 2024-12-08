package response

import (
	"flint/src/generator/route"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Response struct {
	Url  string
	Path string
	Doc  *goquery.Document
}

func NewFromRoute(serverUrl string, route *route.Route) (*Response, error) {
	url := serverUrl + route.Path
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return nil, fmt.Errorf(
			"endpoint %s returned content type %s, expected text/html",
			url, contentType,
		)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf(
			"⚠️ Warning: Expected status code 200, but got %d for URL %s\n",
			resp.StatusCode, url,
		)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	response := &Response{
		Url:  url,
		Path: route.Path,
		Doc:  doc,
	}
	return response, nil
}

func NewFromRoutes(serverUrl string, routes []*route.Route) ([]*Response, error) {
	responses := make([]*Response, 0)
	for _, route := range routes {
		res, err := NewFromRoute(serverUrl, route)
		if err != nil {
			return responses, err
		}
		responses = append(responses, res)
	}
	return responses, nil
}
