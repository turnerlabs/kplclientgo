# kplclientgo

A go client library for [kplserver](https://github.com/turnerlabs/kplserver)


### usage

```go
package main

import "github.com/turnerlabs/kplclientgo"

func main() {

  //create a client
  kpl := NewKPLClient("127.0.0.1:3000", "my-kinesis-data-stream")

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