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

type VolumeBackup struct {
	Params BackupParams
}

func (vb *VolumeBackup) Backup() error {

	err := vb.validateParams()
	if err != nil {
		return err
	}

	command := vb.getBackupCmd()
	out, err := exec.Command("sh", "-c", command).Output()
	fmt.Println(out)
	if err != nil {
		return err
	}
	return nil
}

func (vb *VolumeBackup) validateParams() error {
	destFolder := vb.Params.destFolder
	if _, err := os.Stat(destFolder); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			return err
		}
	}
	return nil
}

func (vb *VolumeBackup) getBackupCmd() string {
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
