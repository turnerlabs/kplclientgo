package kplclientgo

import (
	"fmt"
	"testing"
	"time"
)

func TestHappy(t *testing.T) {

	//give the test harness a chance to boot
	time.Sleep(10 * time.Second)

	c := NewKPLClient("127.0.0.1", "3000")

	err := c.Start()
	if err != nil {
		t.Error(err)
	}

	err = c.PutRecord("test")
	if err != nil {
		t.Error(err)
	}

	//wait before shutting down
	time.Sleep(20 * time.Second)
}

func TestNegative(t *testing.T) {

	//give the test harness a chance to boot
	time.Sleep(10 * time.Second)

	c := NewKPLClient("127.0.0.1", "3000")

	c.ErrPort = ":3001"
	c.ErrHandler = func(data string) {
		fmt.Println("errObj:", data)
	}

	err := c.Start()
	if err != nil {
		t.Error(err)
	}

	err = c.PutRecord("test")
	if err != nil {
		t.Error(err)
	}

	//wait before shutting down
	time.Sleep(20 * time.Second)
}
