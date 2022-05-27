package client

import (
	"charting/pkg/api"
	"charting/pkg/profile"
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	ctx := &profile.Context{
		User:    "chenyang",
		Server:  "127.0.0.1:8080",
		Channel: "default",
	}
	client, err := NewClient(ctx)
	if err != nil {
		t.Error(err)
	}
	if err := client.SendMessage("", &api.Message{Content: []byte("hello")}); err != nil {
		t.Error(err)
	}
}
func TestReceive(t *testing.T) {
	ctx := &profile.Context{
		User:    "chenyang",
		Server:  "127.0.0.1:8080",
		Channel: "default",
	}
	client, err := NewClient(ctx)
	if err != nil {
		t.Error(err)
	}
	msg, err := client.ReceiveMessage("")
	if err != nil {
		t.Error(err)
	}
	for i := range msg {
		fmt.Println(msg[i].Content)
	}
}

func TestConnect(t *testing.T) {
	ctx := &profile.Context{
		User:    "chenyang",
		Server:  "127.0.0.1:8080",
		Channel: "default",
	}
	client, err := NewClient(ctx)
	if err != nil {
		t.Error(err)
	}
	if err := client.Connect("default"); err != nil {
		t.Error(err)
	}
}

