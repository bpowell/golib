package multilog

import (
	"os"
	"log"
	"io"
)

func NewMultiLog(prefix string, flag int, filename string) (mlog *log.Logger, err error){
	//Open up log file for writing/appending
	logf, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
        return nil, err
	}

	//Setup the new log to use a MultiWriter to write to log file and stdout
	mlog = log.New(io.MultiWriter(logf, os.Stdout), prefix, flag)

	return mlog, err
}
