package rabbitmq

import (
	"gin-inject/repository"
	"reflect"
)

type Listener struct {
	Repo repository.IStartRepo
}

func (l *Listener) AutoFunc() {
	valueOf := reflect.ValueOf(l)
	numOfMethod := valueOf.NumMethod()
	for i := 1; i < numOfMethod; i++ {
		valueOf.Method(i).Call(nil)
	}
}

func (l *Listener) SomeListener() {
	mq := &Mq{}
	mq.BaseListener(
		&MqConsumeConfig{
			Queue:     "",
			consumer:  "",
			autoAck:   true,
			exclusive: false,
			noLocal:   false,
			noWait:    false,
			args:      nil,
		},
		func(msg []byte) {
			//do something...
		},
	)
}

func (l *Listener) OtherListener() {
	mq := &Mq{}
	mq.BaseListener(
		&MqConsumeConfig{
			Queue:     "",
			consumer:  "",
			autoAck:   true,
			exclusive: false,
			noLocal:   false,
			noWait:    false,
			args:      nil,
		},
		func(msg []byte) {
		},
	)
}
