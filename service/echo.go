package service

import (
  "log"
)

type Echo string

type Args struct {
  Name    string
}

type Reply struct {
  Greet   string
}

func (e *Echo) Say(args *Args, reply *Reply) error {
  log.Printf("Echo Say args %+v ------------------ ", args)
  reply.Greet = "hello"
  return nil
}


