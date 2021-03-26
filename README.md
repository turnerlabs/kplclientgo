# kplclientgo

A go client library for [kplserver](https://github.com/turnerlabs/kplserver)

If ErrPort is set, this library can be used to retrieve any message that kplserver may fail to send to kinesis. You can set a callback function ErrHandler that will be invoked for every message. See the example usage below on how to set these vaiables. 

### Usage

```go
package main

import (
	"log"
	"github.com/turnerlabs/kplclientgo"
)

func main() {

	//create a client
	kpl := kplclientgo.NewKPLClient("127.0.0.1", "3000")

	// Optionally handle the failed messages
	kpl.ErrPort = "3000"
	kpl.ErrHost = "127.0.0.1"
	
	kpl.ErrHandler = func(data string) {
		log.Print("Could not send the message to kinesis", data)
	}

	//start it up
	err := kpl.Start()
	if err != nil {
		panic(err)
	}

	//send a record
	err = kpl.PutRecord("some data")
	if err != nil {
		panic(err)
	}
}
```

### development

Requires:

- Go 1.13
- docker-compose

To run tests

```sh
AWS_PROFILE=my-profile KINESIS_STREAM=my-stream ERROR_SOCKET_PORT=3001 go test
```
