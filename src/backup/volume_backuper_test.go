package backup

import (
	"config"
)

import (
	"testing"
	"os"
)

func TestVolumeBackuper(t *testing.T) {
	configuration := config.Configuration{}
	path, _ := os.Getwd()
	configuration.Load(path + "/../../.env.test/config.yaml")

	volumeBackuper := VolumeBackuper{}
	volumeBackuper.LoadConfig(configuration)
	volumeBackuper.Run()
}
