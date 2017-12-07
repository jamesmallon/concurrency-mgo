package log

import (
	"fmt"
	"os"
	"runtime"
)

type sysLog struct {
	file *os.File
}

func NewSysLog() (*sysLog, error) {
	file, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return &sysLog{file}, err
}

func (lg *sysLog) GetTraceMsg() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
}

func (lg *sysLog) Debug() error {
	_, err := lg.file.Write([]byte("appended some data\n"))

	if err != nil {
		fmt.Println("")
	}
	return err
}

func (lg *sysLog) close() error {
	err := lg.file.Close()
	if err != nil {
		fmt.Println("")
	}
	return err
}
