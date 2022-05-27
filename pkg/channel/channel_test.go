package channel

import (
	"charting/pkg/api"
	"fmt"
	"testing"
)

func TestChannelReadAndWrite(t *testing.T) {
	ch := NewChannel("test")
	name := "zhangsan"
	fmt.Println("now channel is empty")
	fmt.Println("read channel")
	ms := ch.Read(name)
	fmt.Println(fmt.Sprintf("read %d message", len(ms)))
	fmt.Println("write 1 message to channel")
	ch.Write(&api.Message{})
	fmt.Println("read  channel ")
	ms = ch.Read(name)
	fmt.Println(fmt.Sprintf("read %d message", len(ms)))
	fmt.Println("write 2 message to channel")
	ch.Write(&api.Message{})
	ch.Write(&api.Message{})
	fmt.Println("read channel")
	ms = ch.Read(name)
	fmt.Println(fmt.Sprintf("read %d message", len(ms)))
	fmt.Println("write 20 message to channel")
	for i := 0; i < 20; i++ {
		ch.Write(&api.Message{})
	}
	fmt.Println("read channel")
	ms = ch.Read(name)
	fmt.Println(fmt.Sprintf("read %d message", len(ms)))

}
