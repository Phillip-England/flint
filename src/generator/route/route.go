package route

import (
	"fmt"

	"github.com/phillip-england/flint/src/generator/config"
)

type Route struct {
	Path string
}

func New(path string) (*Route, error) {
	route := &Route{
		Path: path,
	}
	return route, nil
}

func NewFromConfig(conf *config.Config) ([]*Route, error) {
	routes := make([]*Route, 0)
	for _, r := range conf.Routes {
		route, err := New(r)
		if err != nil {
			return routes, err
		}
		routes = append(routes, route)
	}
	return routes, nil
}

func (route *Route) Print() { fmt.Println(route.Path) }
