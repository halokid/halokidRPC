package service

import (
  "context"
  "log"
)

type Echo string

type Args struct {
  Name    string
}

type Reply struct {
  Greet   string
}

func (e *Echo) Say(ctx context.Context, args *Args, reply *Reply) error {
  log.Printf("Echo Say args %+v ------------------ ", args)
  return nil
}


