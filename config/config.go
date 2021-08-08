package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Database struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"databaseName"`
}

type Server struct {
	Host                  string `yaml:"host"`
	Port                  int    `yaml:"port"`
	Prefork               bool   `yaml:"prefork"`
	ReadTimeout           int    `yaml:"timeout"`
	DisableStartupMessage bool   `yaml:"disableStartupMessage"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

func GetConfig(path string) (*Config, error) {
	if err := validateConfigPath(path); err != nil {
		return nil, err
	}

	config := &Config{}
	// Open config file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
