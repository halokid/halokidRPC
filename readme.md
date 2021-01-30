HaloKidRPC
=========================================

a RPC sample base on http/2, client direct connect server,
no register center. use local map store register info,
this will show you how to build a RPC server & client use pure go code.

### Features
* use HTTP/2 without TLS(h2c)
* protocol parse as http arguments
* support codec

### Server
```shell
cd cmd & go run server.go
```
this will make a server, register service

### Client
```shell
cd cmd & go run client.go
```
then will invoke the Echo service Say method

### Summary
check this link: http://www.pangulab.com/post/kb-yo8n7r.html



