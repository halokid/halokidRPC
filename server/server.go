package server

import (
  "errors"
  "fmt"
  "golang.org/x/net/http2"
  "log"
  "net/http"
  "reflect"
  "sync"

  "../service"
)

type Server struct {
  servicesMu        sync.RWMutex
  services          map[string]*service.Service
}

func NewServer() *Server {
  return &Server{
    services:    make(map[string]*service.Service),
  }
}

// no register center, just register local
func (s *Server) RegisterSvc(name string, rcvr interface{}) error {
  s.servicesMu.Lock()
  defer s.servicesMu.Unlock()
  svc := new(service.Service)
  svc.Typ = reflect.TypeOf(rcvr)
  svc.Rcvr =  reflect.ValueOf(rcvr)
  svc.Name = name
  svc.Method = service.FoundMethods(svc.Typ)

  s.services[name] = svc

  if len(svc.Method) == 0 {
    errStr := fmt.Sprintf("服务%s没有可执行的方法", name)
    return errors.New(errStr)
  }
  return nil
}

// run server
func (s *Server) Run(addr string) {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    log.Printf("------收到RPC请求------")
  })

  sx := &http.Server{
    Addr:       addr,
    Handler:    mux,
  }
  http2.ConfigureServer(sx, &http2.Server{})
  log.Fatal(sx.ListenAndServe())
}



