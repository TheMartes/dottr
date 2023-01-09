package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/TheMartes/dottr/dirutils"
	yaml "gopkg.in/yaml.v3"
)

var DottrConfigName = ".dottr.yaml"

type DottrConfig struct {
	DottrFolder   string `yaml:"repo_dir"`
	DottrGitUrl   string `yaml:"repo_url"`
	IncludeConfig bool   `yaml:"include_config"`

	FoldersToSync []string `yaml:"folders_to_sync"`
	FilesToSync   []string `yaml:"files_to_sync"`
}

func LoadConfig() DottrConfig {
	config := DottrConfig{}
	fileName := filepath.Join(dirutils.GetHomeDirectory(), DottrConfigName)

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Print(err)
	}

	err = yaml.Unmarshal([]byte(fileContent), &config)

	return config
}
