# log

Structured leveled logging package for Go.

## Installation

`go get -u github.com/ninefivefly/log`

## Quick Start

```go
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
```

When logging to file.

```go
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
```
