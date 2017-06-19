package backup

import (
	"config"
)

import (
	"testing"
	"os"
)

func TestMysqlDumper(t *testing.T) {
	configuration := config.Configuration{}
	path, _ := os.Getwd()
	configuration.Load(path + "/../../.env.test/config.yaml")

	mysqlDumperProcessing := MysqlDumperProcessing{}
	mysqlDumperProcessing.LoadConfig(configuration)
	mysqlDumperProcessing.Run()
}
