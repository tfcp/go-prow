package mq

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestPusher_Kafka(t *testing.T) {
	var err error
	pusher := NewPusher("kafka")
	gtest.C(t, func(t *gtest.T) {
		err = pusher.Push([]byte("this is demo data"), "demo-topic")
		t.Assert(err, nil)
	})
}
