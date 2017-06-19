package backup

import (
	"strconv"
	"time"
	"fmt"
	"strings"
	"os/exec"
	"os"
	"syscall"
	"github.com/pkg/errors"
)

type MysqlDumper struct {
	containerName   string
	user            string
	password        string
	database        string
	mysqldumpParams []string
	destName        string
	destFolder      string
	destPath        string
}

func (md *MysqlDumper) LoadParams(
	containerName string,
	user string,
	password string,
	database string,
	mysqldumpParams []string,
	destName string,
	destFolder string,
) {
	md.containerName = containerName
	md.user = user
	md.password = password
	md.database = database
	md.mysqldumpParams = mysqldumpParams
	md.destName = destName
	md.destFolder = destFolder
}

func (md *MysqlDumper) Dump(process *chan string) error {
	*process <- fmt.Sprintf("Backup params validated")

	command := md.getBackupCmd()
	*process <- fmt.Sprintf("Command generated:")
	*process <- fmt.Sprintf(command)
	cmd := exec.Command("sh", "-c", command)

	var waitStatus syscall.WaitStatus
	if out, err := cmd.CombinedOutput(); err != nil {
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			*process <- fmt.Sprintf("Exit code: %d", waitStatus.ExitStatus())
			*process <- string(out)
		}
		return err
	} else {
		// Command was successful
		if len(out) > 0 {
			*process <- string(out)
			return errors.New("Unexcepted output")
		}
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		*process <- fmt.Sprintf("%d", waitStatus.ExitStatus())
	}

	*process <- fmt.Sprintf("Command executed:")
	return nil
}

func (md *MysqlDumper) getBackupCmd() string {

	mysqlParams := strings.Join(md.mysqldumpParams, " ")
	command := fmt.Sprintf("docker exec %s /usr/bin/mysqldump -u %s --password=%s %s  %s | gzip > %s",
		md.containerName,
		md.user,
		md.password,
		md.database,
		mysqlParams,
		md.generateDestPath(),
	)
	return command
}

func (md *MysqlDumper) generateDestPath() string {
	timeString := strconv.Itoa(int(time.Now().Unix()))
	fileName := md.destName + timeString
	md.destPath = md.destFolder + string(os.PathSeparator) + fileName + ".sql.gz"
	return md.getDestPath()
}
func (md *MysqlDumper) getDestPath() string {
	return md.destPath
}
