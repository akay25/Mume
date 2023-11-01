package backend

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

// TODO: Setup default config values here
const APPLICATION_NAME = "Mume"
const CONFIG_FILE_LOCATION = "config.json"
const MIN_WINDOW_WIDTH = 1000
const MIN_WINDOW_HEIGHT = 563

// Config setup
type Config struct {
	LibraryPath  string `json:"LibraryPath"`
	LogLevel     string `json:"LogLevel"`
	WindowWidth  int    `json:"WindowWidth"`
	WindowHeight int    `json:"WindowHeight"`
}

func DefaultConfig() Config {
	return Config{
		LibraryPath:  "",
		LogLevel:     "debug",
		WindowWidth:  1000,
		WindowHeight: 563,
	}
}

type ConfigStore struct {
	configPath string
}

func NewConfigStore() (*ConfigStore, error) {
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", APPLICATION_NAME, CONFIG_FILE_LOCATION))
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
	if err := json.Unmarshal(buf, &cfg); err != nil {
		return Config{}, fmt.Errorf("configuration file does not have a valid format: %w", err)
	}

	return cfg, nil

}
