// Golang 实现抽象类 https://go.dev/play/p/aE8E1JSfrxN
package main

import (
	"fmt"
	"testing"
)

type ILog interface {
	Debuger
	Errorer
	Info()
}

type Debuger interface {
	Debug()
}

type Errorer interface {
	Error()
}

// 这个就是抽象类
type abstractLog struct {
	Debuger
	Errorer
}

func (a abstractLog) Info() {
	fmt.Println("abstractLog info")
}

type abstractDebug struct{}

func (a abstractDebug) Debug() {
	fmt.Println("abstractDebug Debug")
}

type abstractError struct{}

func (n abstractError) Error() {
	fmt.Println("abstractError Error")
}

// 用这个方法来使用授象类，实现 java 的 extend
func NewLog(mylog any) ILog {
	alog := &abstractLog{}
	if d, ok := mylog.(Debuger); ok {
		alog.Debuger = d
	} else {
		alog.Debuger = &abstractDebug{}
	}
	if e, ok := mylog.(Errorer); ok {
		alog.Errorer = e
	} else {
		alog.Errorer = &abstractError{}
	}
	return alog
}

// 如下3个是使用抽象类实现部分接口的类

type MyErrorLog struct{}

func (n MyErrorLog) Error() {
	fmt.Println("MyErrorLog Error")
}

type MyDebugLog struct{}

func (n MyDebugLog) Debug() {
	fmt.Println("MyDebugLog Debug")
}

type MyFullLog struct{}

func (n MyFullLog) Error() {
	fmt.Println("MyFullLog Error")
}

func (n MyFullLog) Debug() {
	fmt.Println("MyFullLog Debug")
}

func TestAbstract(t *testing.T) {
	c := NewLog(&MyDebugLog{})
	c.Debug()
	c.Info()
	c.Error()

	e := NewLog(&MyErrorLog{})
	e.Debug()
	e.Info()
	e.Error()

	f := NewLog(&MyFullLog{})
	f.Debug()
	f.Info()
	f.Error()
}
