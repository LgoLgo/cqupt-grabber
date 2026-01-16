package client

import (
	"bytes"
	"log"
	"strings"
	"sync"
)

type LogLevel int

const (
	LogInfo LogLevel = iota
	LogWarn
	LogError
	LogSuccess
)

type uiLogWriter struct {
	mu    sync.Mutex
	buf   bytes.Buffer
	logf  func(level LogLevel, msg string)
	level func(line string) LogLevel
}

func (w *uiLogWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.buf.Write(p)
	s := w.buf.String()
	lines := strings.Split(s, "\n")

	w.buf.Reset()
	if tail := lines[len(lines)-1]; tail != "" {
		w.buf.WriteString(tail)
	}
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		if strings.TrimSpace(line) == "" {
			continue
		}

		lv := LogInfo
		if w.level != nil {
			lv = w.level(line)
		}
		w.logf(lv, line)
	}
	return len(p), nil
}

func HookStdLog(append func(level LogLevel, msg string)) {
	w := &uiLogWriter{
		logf: append,
		level: func(line string) LogLevel {
			if strings.Contains(line, "抢课成功") || strings.Contains(line, "ok") {
				return LogSuccess
			}
			if strings.Contains(line, "失败") {
				return LogWarn
			}
			if strings.Contains(strings.ToLower(line), "error") || strings.Contains(line, "错误") {
				return LogError
			}
			return LogInfo
		},
	}

	log.SetOutput(w)

	// log.SetFlags(0)
}
