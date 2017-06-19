package backup

import (
	"config"
	"fmt"
	"time"
)

type VolumeBackuper struct {
	configuration config.Configuration
}

func (vb *VolumeBackuper) LoadConfig(config config.Configuration) {
	vb.configuration = config
}

func (vb *VolumeBackuper) Run() {
	errors := make(chan string)
	processing := make(chan string)
	vb.processFolders(processing, errors)
	go func() {
		for {
			select {
			case errorString := <-errors:
				fmt.Println("Error:", errorString)
			case process := <-processing:
				fmt.Println(process)
			}
		}
	}()
	time.Sleep(10000000000)
}

func (vb *VolumeBackuper) processFolders(process chan string, errors chan string) {
	folders := vb.getFolders()
	for _, folder := range folders {
		go func() {
			process <- fmt.Sprintf("Backup folder %s in volume %s started",
				folder.Params.path,
				folder.Params.volumeName,
			)
			err := folder.Backup(process)
			process <- fmt.Sprintf("Backup folder %s in volume %s completed. Backup file in %s",
				folder.Params.path,
				folder.Params.volumeName,
				folder.Params.destFolder,
			)
			if err != nil {
				errors <- err.Error()
			}
		}()
	}
}

func (vb *VolumeBackuper) getFolders() []FolderBackup {
	folders := []FolderBackup{}
	volumes := vb.configuration.GetVolumes()
	for volumeName, volumeConfig := range volumes {
		for _, folder := range volumeConfig.Folders {
			folderBackupParams := BackupParams{
				volumeName: volumeName,
				path:       folder.Path,
				destFolder: folder.DestFolder,
				destName:   folder.DestName,
				tempImage:  folder.TempImage,
			}
			folderBackup := FolderBackup{
				Params: folderBackupParams,
			}
			folders = append(folders, folderBackup)
		}
	}
	return folders
}
