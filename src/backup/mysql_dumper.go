package backup

import (
	"fmt"
	"config"
	"sync"
)

type MysqlDumperProcessing struct {
	configuration config.Configuration
}

func (md *MysqlDumperProcessing) LoadConfig(config config.Configuration) {
	md.configuration = config
}

func (md *MysqlDumperProcessing) Run() {
	var wg sync.WaitGroup
	errors := make(chan string)
	processing := make(chan string)
	finished := make(chan bool)
	md.dumpProcessing(&wg, processing, errors, finished)
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

func (vb *MysqlDumperProcessing) dumpProcessing(wg *sync.WaitGroup, process chan string, errors chan string, finished chan bool) {
	dumpers := vb.getDumpers()
	for _, dumper := range dumpers {
		wg.Add(1)
		go func(dumper MysqlDumper) {
			process <- fmt.Sprintf("Backup mysql database %s",
				dumper.database,
			)
			err := dumper.Dump(&process)
			process <- fmt.Sprintf("Backup mysql database %s completed. Backup file in %s",
				dumper.database,
				dumper.getDestPath(),
			)
			if err != nil {
				errors <- err.Error()
			}
			wg.Done()
		}(dumper)
	}
}

func (md *MysqlDumperProcessing) getDumpers() []MysqlDumper {
	mysqlDumpers := []MysqlDumper{}
	mysqlDumpersConfig := md.configuration.GetMysqlDumpers()
	for _, mysqlDumperConfig := range mysqlDumpersConfig {
		mysqlDumper := MysqlDumper{}
		mysqlDumper.LoadParams(
			mysqlDumperConfig.Container,
			mysqlDumperConfig.User,
			mysqlDumperConfig.Pass,
			mysqlDumperConfig.Database,
			mysqlDumperConfig.MysqldumpParams,
			mysqlDumperConfig.DestName,
			mysqlDumperConfig.DestFolder,
		)
		mysqlDumpers = append(mysqlDumpers, mysqlDumper)

	}
	return mysqlDumpers
}
