package service

import (
  "reflect"
  "testing"
)

func TestFoundMethods(t *testing.T) {
  rcvr := new(Echo)
  svcTyp := reflect.TypeOf(rcvr)
  methods := FoundMethods(svcTyp)
  t.Logf("methods ----------- %d, %+v", len(methods), methods)
}

func TestCall(t *testing.T) {
  rcvr := new(Echo)
  svcTyp := reflect.TypeOf(rcvr)
  methods := FoundMethods(svcTyp)
  t.Logf("methods ----------- %d, %+v", len(methods), methods)
  //methodName := "Say"
  //mtype := methods[methodName]

}