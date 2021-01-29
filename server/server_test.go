package server

import (
  "../service"
  "testing"
)

func TestServer_RegisterSvc(t *testing.T) {
  s := NewServer()
  err := s.RegisterSvc("Echo", new(service.Echo))
  t.Log("err ---------------", err)
  t.Logf("s --------------- %+v", s)
}

