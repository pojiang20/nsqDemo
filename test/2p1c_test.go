package test

import (
	"log"
	"testing"
	"time"
)

func TestNSQ1(t *testing.T) {
	addrs := []string{"127.0.0.1:4150", "127.0.0.1:4152"}
	for i := 0; i < 50; i++ {
		log.Println("---------------", i)
		go consumer1(addrs)
		go produce1(addrs[0])
		go produce2(addrs[1])
		time.Sleep(time.Second)
	}
}

func TestNSQ2(t *testing.T) {
	lookupAddr := "127.0.0.1:4161"
	addr := "127.0.0.1:4150"
	addr1 := "127.0.0.1:4152"
	for i := 0; i < 50; i++ {
		log.Println("---------------", i)
		go consumer2(lookupAddr)
		go produce1(addr)
		go produce2(addr1)
		time.Sleep(time.Second)
	}
}
