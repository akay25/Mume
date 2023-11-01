package backend

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

// TODO: Setup default config values here
const APPLICATION_NAME = "Mume"
const CONFIG_FILE_LOCATION = "config.json"
const MIN_WINDOW_WIDTH = 1000
const MIN_WINDOW_HEIGHT = 563

var CONFIG_FILE_PATH = fmt.Sprintf("%s/%s", APPLICATION_NAME, CONFIG_FILE_LOCATION)
var logger = Logger

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
	configFilePath, err := xdg.ConfigFile(CONFIG_FILE_PATH)
	if err != nil {
		return nil, fmt.Errorf("could not resolve path for config file: %w", err)
	}

	return &ConfigStore{
		configPath: configFilePath,
	}, nil
}

func (s *ConfigStore) GetConfig() (Config, error) {
	_, err := os.Stat(s.configPath)
	if os.IsNotExist(err) {
		logger.Debugf("No config file or dir (%s) exists, loading default config", s.configPath)
		defaultConfig := DefaultConfig()
		// Save this default config in s.configPath
		s.UpdateConfig(defaultConfig)
		return defaultConfig, nil
	}

	dir, fileName := filepath.Split(s.configPath)
	if len(dir) == 0 {
		dir = "."
	}

	buf, err := fs.ReadFile(os.DirFS(dir), fileName)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read the configuration file: %w", err)
	}

	if len(buf) == 0 {
		logger.Debugf("No config file or dir (%s) exists, loading default config", s.configPath)
		return DefaultConfig(), nil
	}

	cfg := Config{}
	if err := json.Unmarshal(buf, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (s *ConfigStore) UpdateConfig(cfg Config) (bool, error) {
	// Check if file exists or not
	_, err := os.Stat(s.configPath)
	if err != nil {
		// Create new file if it doesn't exist
		f, err := os.Create(s.configPath)
		if err != nil {
			return false, fmt.Errorf("failed to create config file: %w", err)
		}
		f.Close()
	}

	// Update the config file
	rawJSON, _ := json.MarshalIndent(cfg, "", " ")
	_ = ioutil.WriteFile(s.configPath, rawJSON, 0644)

	return true, nil
}
