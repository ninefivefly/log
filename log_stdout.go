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
)

// Stdout is a stdout logger
type Stdout struct {
	level  Level
	logger *log.Logger
}

// NewStdout create new default stdout logger
func NewStdout() *Stdout {
	return &Stdout{
		logger: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		level:  LevelDebug,
	}
}

// Output create log
func (s *Stdout) Output(lv Level, log string) {
	if lv >= s.level {
		s.logger.Output(4, fmt.Sprintf("%s: %s", lv.String(), log))
	}
}

// SetLogLevel set log level
func (s *Stdout) SetLogLevel(lv Level) {
	s.level = lv
}
