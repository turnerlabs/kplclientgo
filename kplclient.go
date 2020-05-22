package main

import (
	"errors"
	"fmt"
	"net"
)

var socket net.Conn
var socketChannel chan string

//NewKPLClient returns a new KPLClient
func NewKPLClient(host, port string) *KPLClient {
	return &KPLClient{
		Host: host,
		Port: port,
	}
}

//KPLClient represents a client to the KPL Server
type KPLClient struct {
	Host    string
	Port    string
	Started bool
}

//Start starts up a communication channel to the server
func (c *KPLClient) Start() error {

	if !c.Started {
		address := fmt.Sprintf("%s:%s", c.Host, c.Port)
		var err error
		socket, err = net.Dial("tcp", address)
		if err != nil {
			return err
		}

		//synchronize records written across the socket
		socketChannel = make(chan string)
		go processChannel()
	}

	c.Started = true
	return nil
}

//Stop shutsdown the communication channel to the server
func (c *KPLClient) Stop() {
	if c.Started {
		socket.Close()
	}
}

func processChannel() {
	for {

		//read record from channel
		r := <-socketChannel

		//write to socket
		socket.Write([]byte(r + "\n"))
	}
}

//PutRecord sends a data record to the KPL server
func (c *KPLClient) PutRecord(record string) error {
	if !c.Started {
		return errors.New("client is not started")
	}
	go func() { socketChannel <- record }()
	return nil
}
