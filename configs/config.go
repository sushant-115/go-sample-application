package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Config struct {
	Server  *ServerConfig `json:"server"`
	MongoDB *MongoConfig  `json:"mongo"`
}

type ServerConfig struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	ReadTimeout  string `json:"read_timeout"`
	WriteTimeout string `json:"write_timeout"`
	IdleTimeout  string `json:"idle_timeout"`
}

type MongoConfig struct {
	DbName   string `json:"db_name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func (s *ServerConfig) GetURI() (uri string, err error) {
	if s.Host == "" || s.Port == "" {
		err = fmt.Errorf("Error server config not initialized")
		return
	}
	uri = s.Host + ":" + s.Port
	return
}

func parseDuration(t string) (duration time.Duration) {
	duration, err := time.ParseDuration(t)
	if err != nil {
		duration = 0 * time.Second
	}
	return
}

func (s *ServerConfig) GetReadTimeout() time.Duration {
	return parseDuration(s.ReadTimeout)
}

func (s *ServerConfig) GetWriteTimeout() time.Duration {
	return parseDuration(s.WriteTimeout)
}

func (s *ServerConfig) GetIdleTimeout() time.Duration {
	return parseDuration(s.IdleTimeout)
}

func (m *MongoConfig) GetURI() (uri string, err error) {
	if m.Host == "" || m.Port == "" {
		err = fmt.Errorf("Error mongo config not initialized")
		return
	}

	if m.User != "" && m.Password != "" {
		uri = "mongodb://" + m.User + ":" + m.Password + "@" + m.Host + ":" + m.Port
	} else {
		uri = "mongodb://" + m.Host + ":" + m.Port
	}
	return
}

func (config *Config) SetDefaults() {
	server := config.Server
	mongo := config.MongoDB

	localhost := "127.0.0.1"
	defaultTimeout := "1s"

	if server.Host == "" {
		server.Host = localhost
	}
	if server.Port == "" {
		server.Host = "8080"
	}
	if mongo.Host == "" {
		mongo.Host = localhost
	}

	if mongo.Port == "" {
		mongo.Port = "27201"
	}

	if server.ReadTimeout == "" {
		server.ReadTimeout = defaultTimeout
	}
	if server.WriteTimeout == "" {
		server.WriteTimeout = defaultTimeout
	}
	if server.IdleTimeout == "" {
		server.IdleTimeout = defaultTimeout
	}
}

func InitConfig(configPath string) (config *Config, err error) {
	config = &Config{}
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, config)
	return
}
