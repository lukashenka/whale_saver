package config

import (
	"fmt"
	"whale_saver/vendor/src/gopkg.in/yaml.v2"
	"io/ioutil"
)

type instanses struct {
	Volumes map[string]volume `yaml:"volumes"`
}

type volume struct {
	Container string   `yaml:"container_name"`
	Folders   []string `yaml:"folders"`
}
type config struct {
	Instanses instanses `yaml:"instanses"`
}

type Configuration struct {
	filename string
	config   config
}

func (c *Configuration) Load(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	config := config{}

	yaml.Unmarshal([]byte(data), &config)

	c.config = config



	return nil
}

func (c *Configuration)GetVolumes() map[string]volume  {
	return c.config.Instanses.Volumes
}