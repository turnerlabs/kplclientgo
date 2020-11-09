# kplclientgo

A go client library for [kplserver](https://github.com/turnerlabs/kplserver)


### usage

```go
package main

import "github.com/turnerlabs/kplclientgo"

func main() {

  //create a client
  kpl := kplclientgo.NewKPLClient("127.0.0.1", "3000")

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
AWS_PROFILE=my-profile KINESIS_STREAM=my-stream DLQ_URL=aws-sqs-url go test
```
