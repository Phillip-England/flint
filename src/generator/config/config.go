package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/phillip-england/purse"
)

type Config struct {
	Host    string   `json:"host"`
	Static  string   `json:"static"`
	Favicon string   `json:"favicon"`
	Out     string   `json:"out"`
	Target  string   `json:"target"`
	Routes  []string `routes:"routes"`
}

func New() (*Config, error) {
	f, err := os.ReadFile("./flint.json")
	if err != nil {
		return nil, err
	}
	var config *Config
	err = json.Unmarshal(f, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing configuration: %w", err)
	}
	return config, nil
}

func (conf *Config) Print() {
	fmt.Println(purse.Fmt(`
		host: %s
		static: %s
		out: %s
		routes: %s
	`, conf.Host, conf.Static, conf.Out, conf.Routes))
}
