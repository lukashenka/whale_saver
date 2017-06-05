package backup

import "testing"

func TestBackupDb(t *testing.T)  {

	volumeBackup := VolumeBackup{
		Params: BackupParams{
			volumeName:"2kola_db",
			destFolder:"/home/karachun/backup",
			destName: "db",
			path:"/var/lib/mysql",
			tempImage: "busybox",
		},
	}
	err:=volumeBackup.Backup()
	if err!=nil {
		t.Errorf(err.Error())
	}
}

func TestBackupWeb(t *testing.T)  {

	volumeBackup := VolumeBackup{
		Params: BackupParams{
			volumeName:"2kola_web",
			destFolder:"/home/karachun/backup",
			destName: "web",
			path:"/var/www/html",
			tempImage: "busybox",
		},
	}
	err:=volumeBackup.Backup()
	if err!=nil {
		t.Errorf(err.Error())
	}
}