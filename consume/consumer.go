package main

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"pojiang20/nsqDemo/common"
)

const ADDR = "127.0.0.1:4150"

func main() {
	InitConsumer("test", "test")
	InitConsumer("test", "lc")
	select {}
}

type MyTestHandler struct {
	q          *nsq.Consumer
	MsgReceive int
}

func (h *MyTestHandler) HandleMessage(message *nsq.Message) error {
	data := common.Data{}
	err := json.Unmarshal(message.Body, &data)
	fmt.Println(data)
	if err != nil {
		log.Println("Unmarshal error: ", message.ID, err)
	}
	message.Finish()
	return nil
}

func InitConsumer(topic string, channel string) {
	handler := &MyTestHandler{}
	cfg := nsq.NewConfig()
	ret, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Println(err)
		return
	}
	handler.q = ret
	handler.q.AddHandler(handler)
	err = handler.q.ConnectToNSQD(ADDR)
	if err != nil {
		log.Println("ConnectToNSQD Error: ", err)
	}
	log.Println("Start consume")
	return
}
