package pkg

import (
	"fmt"
	"time"
)

var (
	LogSavePath = "runtime/logs"
	LogSaveName = "log"
	LogFileExt  = "log"
	Format      = "2006/01/02"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := LogFileExt
	fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(Format), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}
