package config

import (
	"testing"
	"os"
)

func TestConfig(t *testing.T) {
	configuration := Configuration{}
	path, _ := os.Getwd()
	err := configuration.Load(path + "/../../.env.test/config.yaml")
	if err != nil {
		t.Error(err.Error())
	}
	if len(configuration.GetVolumes()) == 0 {
		t.Error("Volumes must me more than 0")
	}

}

