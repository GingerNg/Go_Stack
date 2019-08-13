package basic

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func DemoGoroutine() {
	/*
		两个 goroutine 在执行
	*/
	go say("world")
	say("hello")
}

func DemoChannel() {
	/*
		goroutine间通信
	*/
}
