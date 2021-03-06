package common

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

type LogWriter struct{}

func (f LogWriter) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(3)

	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string

	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	log.Printf("%s:%d %s: %s", filepath.Base(file), line, fnName, p)

	return len(p), nil
}

var (
	LogE = log.New(LogWriter{}, "ERROR -> ", 0)
	LogW = log.New(LogWriter{}, "WARN -> ", 0)
	LogI = log.New(LogWriter{}, "INFO -> ", 0)
)
