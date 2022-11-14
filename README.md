# go_rest_api
Simple REST API built on Go native libraries.  
Step by step guid

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
- a little research about Go http package [net/http](https://pkg.go.dev/net/http)
- Add http listener
    >   http.ListenAndServe(":3000", nil)
- add logger
- error 
    > listen tcp :3000: bind: address already in use
- kill listener 
    >   kill -9 $(lsof -ti:3000)


