// Copyright 2015 tsuru-autoscale authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"log"
	"os"

	"github.com/getsentry/raven-go"
)

var lg *Logger

type Logger struct {
	lg *log.Logger
}

func (l *Logger) Print(v ...interface{}) {
	l.lg.Print(v)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.lg.Printf(format, v)
}

func (l *Logger) Error(err error) {
	raven.CaptureError(err, nil)
	l.Print(err)
}

func New() *Logger {
	return &Logger{lg: log.New(os.Stdout, "[autoscale] ", 0)}
}

func Log() *Logger {
	if lg == nil {
		lg = New()
	}
	return lg
}
