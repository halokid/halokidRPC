package main

import (
  "./server"
)

func main() {
  s := server.NewServer()
  s.Run(":9527")
}


