package conf

import (
	"sync"
)

var Conf *Config
var once sync.Once

type Config struct {
	ApiServer  *ApiServerConfig
	Prometheus *PrometheusConfig
}

type ApiServerConfig struct {
	Port string `json:"port"`
}
type PrometheusConfig struct {
	URL string `json:"url"`
}

func init() {
	once.Do(func() {
		Conf = &Config{}
	})
	Conf.ApiServer = &ApiServerConfig{
		Port: "8092",
	}
	Conf.Prometheus = &PrometheusConfig{
		URL: "127.0.0.1:31945",
	}
}
