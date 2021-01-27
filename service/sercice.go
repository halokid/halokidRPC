package service

import (
  "reflect"
  "sync"
)

type methodType struct {
  sync.Mutex // protects counters
  method     reflect.Method     // the method of struct pointer
  ArgType    reflect.Type
  ReplyType  reflect.Type
  // numCalls   uint
}

type functionType struct {
  sync.Mutex // protects counters
  fn         reflect.Value
  ArgType    reflect.Type
  ReplyType  reflect.Type
}

type service struct {
  name      string              // name of sercice
  // receiver of methods for the service, like new(service), a struct pointer
  rcvr      reflect.Value
  typ       reflect.Type        // type of the receiver
  method    map[string]*methodType
  function  map[string]*functionType
}

func (s *service) call(mtype *methodType, argv, replyv reflect.Value) error {
  function := mtype.method.Func
  // invoke the method, providing a new value for the reply
  returnVal := function.Call([]reflect.Value{s.rcvr, argv, replyv})
  // the return val for the method is an error
  err := returnVal[0].Interface()
  if err != nil {    // invoke fail
    return err.(error)
  }
  return nil
}






