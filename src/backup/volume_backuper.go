package backup

import (
	"config"
	"fmt"
	"sync"
)

type VolumeBackuper struct {
	configuration config.Configuration
}

func (vb *VolumeBackuper) LoadConfig(config config.Configuration) {
	vb.configuration = config
}

func (vb *VolumeBackuper) Run() {
	var wg sync.WaitGroup
	errors := make(chan string)
	processing := make(chan string)
	finished := make(chan bool)
	vb.processFolders(&wg, processing, errors, finished)
	go func() {
		for {
			select {
			case errorString := <-errors:
				fmt.Println("Error:", errorString)
			case process := <-processing:
				fmt.Println(process)
			case <-finished:
				fmt.Println("Jobs done!")
				return
			}
		}
	}()
	wg.Wait()
	finished <- true
}

func (vb *VolumeBackuper) processFolders(wg *sync.WaitGroup, process chan string, errors chan string, finished chan bool) {
	folders := vb.getFolders()
	for _, folder := range folders {
		wg.Add(1)
		go func() {
			process <- fmt.Sprintf("Backup folder %s in volume %s started",
				folder.Params.path,
				folder.Params.volumeName,
			)
			err := folder.Backup(&process)
			process <- fmt.Sprintf("Backup folder %s in volume %s completed. Backup file in %s",
				folder.Params.path,
				folder.Params.volumeName,
				folder.Params.destFolder,
			)
			if err != nil {
				errors <- err.Error()
			}
			wg.Done()
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
