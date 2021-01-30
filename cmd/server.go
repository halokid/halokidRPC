package main

import (
  "../server"
  "../service"
  "log"
)

func main() {
  s := server.NewServer()
  s.RegisterSvc("Echo", new(service.Echo))
  log.Printf("s -------------- %+v", s)
  s.Run(":9527")
}


