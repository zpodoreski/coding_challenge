package config

import (
	"os"

	"github.com/coding_challenge/pkg/model"
	"gopkg.in/yaml.v2"
)

// Getting config data from config.yaml and decoding it into struct
func GetConfig(cfg *model.Config) error {
	f, err := os.Open(os.Getenv("GOPATH") + "/src/github.com/coding_challenge/config.yml")
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		return err
	}
	return nil
}
