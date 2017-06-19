package backup

import (
	"fmt"
	"os/exec"
	"os"
	"time"
	"strconv"
)

type BackupParams struct {
	volumeName string
	path       string // from path in volume
	destFolder string // to path in local storage
	destName   string // to file name in local storage
	tempImage  string // image where backup temporary stored
}

type FolderBackup struct {
	Params BackupParams
}

func (vb *FolderBackup) Backup(process *chan string) (error) {

	err := vb.validateParams()
	if err != nil {
		return err
	}
	*process <- fmt.Sprintf("Backup params validated")

	command := vb.getBackupCmd()
	*process <- fmt.Sprintf("Command generated:")
	*process <- fmt.Sprintf(command)
	out, err := exec.Command("sh", "-c", command).Output()
	*process <- fmt.Sprintf("Command executed:")
	*process <- fmt.Sprintf(string(out))
	if err != nil {
		return err
	}
	return nil
}


func (vb *FolderBackup) validateParams() error {
	destFolder := vb.Params.destFolder
	if _, err := os.Stat(destFolder); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			return err
		}
	}
	return nil
}

func (vb *FolderBackup) getBackupCmd() string {
	timeString := strconv.Itoa(int(time.Now().Unix()))
	fileName := vb.Params.destName + timeString
	command := fmt.Sprintf("docker run --rm --volume %s:%s -v %s:/backup %s tar cvf /backup/%s.tar.gz %s",
		vb.Params.volumeName,
		vb.Params.path,
		vb.Params.destFolder,
		vb.Params.tempImage,
		fileName,
		vb.Params.path)
	return command
}
