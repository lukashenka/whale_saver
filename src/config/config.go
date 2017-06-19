package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Instances struct {
	Volumes map[string]Volume `yaml:"volumes"`
}

type Folder struct {
	Path       string `yaml:"path"`
	DestFolder string `yaml:"destFolder"`  // to path in local storage
	DestName   string `yaml:"destName"` // to file name in local storage
	TempImage  string `yaml:"tempImage"` // image where backup temporary stored
}
type Volume struct {
	Folders map[string]Folder `yaml:"folders"`
}
type Config struct {
	Instances Instances `yaml:"instanсes"`
	BackupFileSend interface{} `yaml:"backup_file_send"`
}

type Configuration struct {
	filename string
	config   Config
}

func (c *Configuration) Load(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	config := Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return err
	}
	c.config = config
	return nil
}


func (c *Configuration) GetVolumes() map[string]Volume {
	return c.config.Instances.Volumes
}
