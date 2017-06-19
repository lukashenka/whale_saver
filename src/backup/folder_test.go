package backup

import (
	"testing"
	"config"
	"os"
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
	_, err := volumeBackup.Backup()
	if err != nil {
		t.Errorf(err.Error())
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
	_, err := volumeBackup.Backup()
	if err != nil {
		t.Errorf(err.Error())
	}
}
