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
	"log"
	"os"
	"time"
)

// FileLogger is a file logger
type FileLogger struct {
	level  Level
	logger *log.Logger
	file   *os.File
}

// NewFileLogger create default file logger
func NewFileLogger(v ...interface{}) *FileLogger {
	var (
		fileName string
		ok       bool
	)

	if len(v) > 0 {
		fileName, ok = v[0].(string)
		if !ok {
			fmt.Println("set error file log path.")
		}
	}

	if len(fileName) == 0 {
		fileName = "./" + time.Now().Format("20180102") + ".log"
	}

	fp, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		fmt.Println(err)
	}

	logger := log.New(fp, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	return &FileLogger{
		level:  LevelDebug,
		logger: logger,
		file:   fp,
	}
}

// Output create log
func (f *FileLogger) Output(lv Level, log string) {
	if lv >= f.level {
		err := f.logger.Output(4, fmt.Sprintf("%s: %s", lv.String(), log))
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Close close file
func (f *FileLogger) Close() {
	f.file.Close()
}
