package backup

import (
	"testing"
	"config"
	"os"
	"fmt"
)

func TestBackupDb(t *testing.T) {
	configuration := config.Configuration{}
	path, _ := os.Getwd()
	configuration.Load(path + "/.env.test/config.yaml")

	volumeBackup := FolderBackup{
		Params: BackupParams{
			volumeName: "2kola_db",
			destFolder: "/home/karachun/backup",
			destName:   "db",
			path:       "/var/lib/mysql",
			tempImage:  "busybox",
		},
	}
	process := make(chan string)
	err := volumeBackup.Backup(&process)
	if err != nil {
		t.Errorf(err.Error())
	}
	for {
		fmt.Println(<-process)
	}

}

func TestBackupWeb(t *testing.T) {

	volumeBackup := FolderBackup{
		Params: BackupParams{
			volumeName: "2kola_web",
			destFolder: "/home/karachun/backup",
			destName:   "web",
			path:       "/var/www/html",
			tempImage:  "busybox",
		},
	}
	process := make(chan string)
	err := volumeBackup.Backup(&process)
	if err != nil {
		t.Errorf(err.Error())
	}
	for {
		fmt.Println(<-process)
	}
}
