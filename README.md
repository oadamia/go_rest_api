# go_rest_api
Simple REST API built on Go native libraries.  
Step by step guid for presentation purposes

## Hello World
---
- Create Go Application with *go mod init* command.  
- Create main.go file with *echo package main > main.go*.  
- Write main function in main.go    
    >    func main() {
    >
    >   }
- print "Hello World"
- *go build*
- *go run .*

## HTTP Server
---
- a little research about Go [net/http](https://pkg.go.dev/net/http)
- Add http listener
    >   http.ListenAndServe(":3000", nil)
- add logger
- error 
    > listen tcp :3000: bind: address already in use
- kill listener 
    >   kill -9 $(lsof -ti:3000)

## Handler
---
- one more little research Go [http Handler](https://pkg.go.dev/net/http#Handler)
- add handler 
- return Hello World

## model and JSON
---
- one more little research Go [json/encoder](https://pkg.go.dev/encoding/json#Encoder)
- movie model
- encode movie model to json



