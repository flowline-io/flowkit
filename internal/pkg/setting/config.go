package setting

import (
	"encoding/json"
	"fmt"
	"github.com/adrg/xdg"
	"io/fs"
	"os"
	"path/filepath"
)

type Config struct {
	ServerHost string `json:"serverHost"`
}

func DefaultConfig() Config {
	return Config{
		ServerHost: "localhost:6060",
	}

}

type ConfigStore struct {
	configPath string
}

func NewConfigStore() (*ConfigStore, error) {
	configFilePath, err := xdg.ConfigFile("flowkit/config.json")
	if err != nil {
		return nil, fmt.Errorf("could not resolve path for config file: %w", err)
	}

	return &ConfigStore{
		configPath: configFilePath,
	}, nil
}

func (s *ConfigStore) Config() (Config, error) {
	_, err := os.Stat(s.configPath)
	if os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	dir, fileName := filepath.Split(s.configPath)
	if len(dir) == 0 {
		dir = "."
	}

	buf, err := fs.ReadFile(os.DirFS(dir), fileName)
	if err != nil {
		return Config{}, fmt.Errorf("could not read the configuration file: %w", err)
	}

	if len(buf) == 0 {
		return DefaultConfig(), nil
	}

	cfg := Config{}
	err = json.Unmarshal(buf, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshal the configuration file: %w", err)
	}

	return cfg, nil
}

func (s ConfigStore) Save(cfg Config) error {
	buf, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("could not marshal the configuration file: %w", err)
	}

	return os.WriteFile(s.configPath, buf, 0644)
}
