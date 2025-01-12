package asset

import (
	"bytes"
	"fmt"

	"github.com/phillip-england/flint/src/generator/response"

	"github.com/PuerkitoBio/goquery"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

type Asset struct {
	Html string
	Path string
}

func NewFromResponse(response *response.Response, targetServer string) (*Asset, error) {
	doc := response.Doc
	var potErr error
	// adjusting our hrefs
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if len(href) == 0 {
			return
		}
		firstChar := string(href[0])
		if firstChar != "/" {
			return
		}
		if href == "/" {
			s.SetAttr("href", targetServer+"/index.html")
		} else {
			s.SetAttr("href", targetServer+href+".html")
		}
	})
	if potErr != nil {
		return nil, potErr
	}
	// adjusting our link href's
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if len(href) == 0 {
			return
		}
		firstChar := string(href[0])
		if firstChar == "/" {
			s.SetAttr("href", targetServer+href)
		}
	})
	if potErr != nil {
		return nil, potErr
	}
	// adjusting our script src's
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if len(src) == 0 {
			return
		}
		firstChar := string(src[0])
		if firstChar == "/" {
			s.SetAttr("src", targetServer+src)
		}
	})
	if potErr != nil {
		return nil, potErr
	}
	// adjusting image tags
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if len(src) == 0 {
			return
		}
		firstChar := string(src[0])
		if firstChar == "/" {
			s.SetAttr("src", targetServer+src)
		}
	})
	if potErr != nil {
		return nil, potErr
	}
	// adjusting iframe src's
	doc.Find("iframe").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if len(src) == 0 {
			return
		}
		firstChar := string(src[0])
		if firstChar == "/" {
			s.SetAttr("src", targetServer+src)
		}
	})
	if potErr != nil {
		return nil, potErr
	}
	// adjusting <object> data attr
	doc.Find("object").Each(func(i int, s *goquery.Selection) {
		data, _ := s.Attr("data")
		if len(data) == 0 {
			return
		}
		firstChar := string(data[0])
		if firstChar == "/" {
			s.SetAttr("data", targetServer+data)
		}
	})
	if potErr != nil {
		return nil, potErr
	}
	htmlStr, err := doc.Html()
	if err != nil {
		return nil, err
	}
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	var buf bytes.Buffer
	err = m.Minify("text/html", &buf, bytes.NewBufferString(htmlStr))
	if err != nil {
		return nil, err
	}
	var path string
	if response.Path == "/" {
		path = "/index.html"
	} else {
		path = response.Path + ".html"
	}
	asset := &Asset{
		Html: buf.String(),
		Path: path,
	}
	return asset, nil
}

func NewFromResponses(responses []*response.Response, targetServer string) ([]*Asset, error) {
	assets := make([]*Asset, 0)
	for _, res := range responses {
		asset, err := NewFromResponse(res, targetServer)
		if err != nil {
			return assets, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

func (asset *Asset) Print() { fmt.Println(asset.Html) }
