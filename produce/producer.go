package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"pojiang20/nsqDemo/common"
	"time"
)

const ADDR = "127.0.0.1:4150"

func main() {
	p1 := &Producer{}
	p2 := &Producer{}
	p1.producer, _ = InitProducer(ADDR)
	p2.producer, _ = InitProducer(ADDR)

	defer p1.producer.Stop()
	defer p2.producer.Stop()

	err := p1.publish("test", "hello")
	if err != nil {
		log.Println(err)
	}
	log.Println("Done")
}

type Producer struct {
	producer *nsq.Producer
}

func (p *Producer) publish(topic string, message string) error {
	if message == "" {
		return errors.New("message is empty")
	}
	data := &common.Data{Msg: message}
	data1, _ := json.Marshal(data)
	if err := p.producer.Publish(topic, data1); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 延迟消息
func (p *Producer) deferredPublish(topic string, delay time.Duration, message string) error {
	if message == "" {
		return errors.New("message is empty")
	}
	if err := p.producer.DeferredPublish(topic, delay, []byte(message)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InitProducer(addr string) (p *nsq.Producer, err error) {
	cfg := nsq.NewConfig()
	p, err = nsq.NewProducer(addr, cfg)
	if err != nil {
		return nil, err
	}
	return p, nil
}
