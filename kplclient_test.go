package kplclientgo

import (
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
	//Enabled it when testing for negative scenario. Pass invalid credential to AWS
	t.SkipNow()
	//give the test harness a chance to boot
	time.Sleep(10 * time.Second)

	c := NewKPLClient("127.0.0.1", "3000")

	c.ErrPort = ":3001"
	c.ErrHandler = func(data string) {
		if data != "test" {
			t.Error("Expected test", "Got", data)
			t.Fail()
		}
	}

	err := c.Start()
	if err != nil {
		t.Error(err)
	}

	err = c.PutRecord("test")
	if err != nil {
		t.Error(err)
	}

	//wait for kpl server to send the message via error channel
	time.Sleep(time.Minute)
}
