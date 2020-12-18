package kplclientgo

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"time"
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
	Host string
	Port string
	// ErrPort is the optional field, which is provided, will cause a server to start and
	// on this port and retrieve any error.
	// Provide ErrHandler if ErrPort is set.
	ErrPort    string
	ErrHost    string
	ErrHandler func(data string)
	Started    bool
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

		if c.ErrPort != "" {
			go c.processErrMessage()
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

func (c *KPLClient) processErrMessage() {
	time.Sleep(time.Second * 5)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", c.ErrHost, c.ErrPort))
	if err != nil {
		fmt.Println("Error listening to error port:", err.Error())
		return
	}

	// Close the listener when the application closes.
	defer conn.Close()

	log.Println("Error Socket Connection Established")

	for {
		//Read from err to socket
		content, err := Read(conn)
		if err != nil {
			log.Printf("Listener: Read error: %v", err)
			continue
		}

		go c.ErrHandler(content)
	}
}

func Read(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	return tp.ReadLine()
}

//PutRecord sends a data record to the KPL server
func (c *KPLClient) PutRecord(record string) error {
	if !c.Started {
		return errors.New("client is not started")
	}
	go func() { socketChannel <- record }()
	return nil
}
