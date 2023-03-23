package test

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func produce1(addr string) {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		log.Fatalln(err)
	}
	msgs := []string{"x", "y"}
	for _, msg := range msgs {
		err := producer.Publish("test", []byte(msg))
		if err != nil {
			log.Fatalln("publish error: " + err.Error())
		}
	}
}

func produce2(addr string) {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		log.Fatalln(err)
	}
	msgs := []string{"z"}
	for _, msg := range msgs {
		err := producer.Publish("test", []byte(msg))
		if err != nil {
			log.Fatalln("publish error: " + err.Error())
		}
	}
}

func consumer1(NSQDsAddrs []string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "channel1", cfg)
	if err != nil {
		log.Fatalln(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		log.Printf("msg: %s from %s,%s", msg.Body, msg.NSQDAddress, msg.ID)
		return nil
	}))
	err = consumer.ConnectToNSQDs(NSQDsAddrs)
	if err != nil {
		log.Fatalln(err)
	}
	<-consumer.StopChan
}

func consumer2(lookupd string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "channel1", cfg)
	if err != nil {
		log.Fatalln(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		log.Printf("msg: %s from %s,%s", msg.Body, msg.NSQDAddress, msg.ID)
		return nil
	}))
	err = consumer.ConnectToNSQLookupd(lookupd)
	if err != nil {
		log.Fatalln(err)
	}
	<-consumer.StopChan
}

func consumer3(NSQDsAddrs []string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "channel1", cfg)
	if err != nil {
		log.Fatalln(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		log.Printf("msg: %s from %s,%s", msg.Body, msg.NSQDAddress, msg.ID)
		return nil
	}))
	err = consumer.ConnectToNSQDs(NSQDsAddrs)
	if err != nil {
		log.Fatalln(err)
	}
	<-consumer.StopChan
}
