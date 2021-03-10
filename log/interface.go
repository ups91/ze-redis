package log

import "fmt"

type Logger interface {
	Log(msg interface{})
}

// STUB
type Log struct{}

func (l *Log) Log(msg interface{}) { fmt.Println(msg) }
