/**********************************************************************************
MIT License

Copyright (c) 2019 JIANG PENG CHENG

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
************************************************************************************/

package log

import (
	"fmt"
)

// Level log level
type Level int

const (
	// LevelDebug log level is debug
	LevelDebug Level = iota
	// LevelInfo log level is info
	LevelInfo
	// LevelWarn log level is warn
	LevelWarn
	// LevelError log level is error
	LevelError
	// LevelFatal log level is fatal
	LevelFatal
)

// String return level string.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "FATAL"
	}
}

// Logger is
type Logger interface {
	Output(lv Level, log string)
}

var loggers = make(map[string]Logger)

func init() {
	SetLogger("std", NewStdout())
}

// SetLogger set logger with key
func SetLogger(key string, l Logger) {
	loggers[key] = l
}

// Debug create debug log
func Debug(v ...interface{}) {
	output(LevelDebug, v...)
}

// Debugf create debug log
func Debugf(format string, v ...interface{}) {
	outputf(LevelDebug, format, v...)
}

// Info create info log
func Info(v ...interface{}) {
	output(LevelInfo, v...)
}

// Infof create Info log
func Infof(format string, v ...interface{}) {
	outputf(LevelInfo, format, v...)
}

// Warn create warn log
func Warn(v ...interface{}) {
	output(LevelWarn, v...)
}

// Warnf create Warn log
func Warnf(format string, v ...interface{}) {
	outputf(LevelWarn, format, v...)
}

// Error create Error log
func Error(v ...interface{}) {
	output(LevelError, v...)
}

// Errorf create Error log
func Errorf(format string, v ...interface{}) {
	outputf(LevelError, format, v...)
}

// Fatal create Fatal log
func Fatal(v ...interface{}) {
	output(LevelFatal, v...)
}

// Fatalf create Fatal log
func Fatalf(format string, v ...interface{}) {
	outputf(LevelFatal, format, v...)
}

func output(level Level, v ...interface{}) {
	log := fmt.Sprint(v...)
	for _, l := range loggers {
		l.Output(level, log)
	}
}

func outputf(level Level, format string, v ...interface{}) {
	log := fmt.Sprintf(format, v...)
	for _, l := range loggers {
		l.Output(level, log)
	}
}
