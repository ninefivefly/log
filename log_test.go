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
	"testing"
)

func TestCreateLog(t *testing.T) {
	Debug("this is debug")
	Debugf("this is %s", "debugf")

	Info("this is info")
	Infof("this is %s", "infof")

	Warn("this is Warn")
	Warnf("this is %s", "Warnf")

	Error("this is Error")
	Errorf("this is %s", "Errorf")

	Fatal("this is Fatal")
	Fatalf("this is %s", "Fatalf")
}

func TestFileLog(t *testing.T) {
	file := NewFileLogger()
	SetLogger("file", file)
	defer file.Close()

	Debug("this is debug")
	Debugf("this is %s", "debugf")

	Info("this is info")
	Infof("this is %s", "infof")

	Warn("this is Warn")
	Warnf("this is %s", "Warnf")

	Error("this is Error")
	Errorf("this is %s", "Errorf")

	Fatal("this is Fatal")
	Fatalf("this is %s", "Fatalf")
}
