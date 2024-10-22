package config

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Chain        int
	Rpc          string
	IpfsNode     string
	MainContract common.Address
	Database     databaseConfig
	Website      websiteConfig
}

type databaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type websiteConfig struct {
	ListenUrl string
}

const prefix = "RF_"

var cached *Config
var k = koanf.New(".")

func Load(configPath string) (*Config, error) {
	if cached != nil {
		return cached, nil
	}

	if configPath != "" {
		if err := k.Load(file.Provider(configPath), toml.Parser()); err != nil {
			return nil, fmt.Errorf("config: reading file %s: %w", configPath, err)
		}
	}

	translateEnv := func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, prefix)), "_", ".", -1)
	}
	if err := k.Load(env.Provider(prefix, ".", translateEnv), nil); err != nil {
		return nil, fmt.Errorf("config: parsing env: %w", err)
	}

	var out Config
	if err := k.Unmarshal("", &out); err != nil {
		return nil, fmt.Errorf("config: unmarshal: %w", err)
	}
	cached = &out
	return cached, nil
}
