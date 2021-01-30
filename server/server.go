package server

import (
  "encoding/json"
  "errors"
  "fmt"
  "github.com/halokid/ColorfulRabbit"
  "golang.org/x/net/http2"
  "golang.org/x/net/http2/h2c"
  "io/ioutil"
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

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    w.Write([]byte("只支持POST方法"))
    return
  }

  svcName := r.Header.Get("svc")
  if svcName == "" {
    w.Write([]byte("未指定服务"))
    return
  }
  mdName := r.Header.Get("method")
  if mdName == "" {
    w.Write([]byte("未指定要调用的方法"))
    return
  }
  svc := s.services[svcName]

  mtype := svc.Method[mdName]
  //payload := r.Body
  //log.Printf("payload ---------- %+v, %+v", reflect.TypeOf(payload), payload)
  payload, err := ioutil.ReadAll(r.Body)
  log.Printf("payload ---------- %+v, %+v", reflect.TypeOf(payload), string(payload))
  ColorfulRabbit.CheckError(err, "读取payload错误")


  // 复用参数类型声明
  argv := service.ArgsReplyPools.Get(mtype.ArgType)
  replyv := service.ArgsReplyPools.Get(mtype.ReplyType)

  // 赋值argv
  err = json.Unmarshal(payload, argv)
  ColorfulRabbit.CheckError(err, "payload json反序列化错误")

  err = svc.Call(mtype, reflect.ValueOf(argv), reflect.ValueOf(replyv))
  log.Printf("HandleRequest err ------- %+v", err)

  rsp, err := json.Marshal(replyv)
  ColorfulRabbit.CheckError(err, "返回给客户端的序列化错误")
  w.Write(rsp)
}

// run server
func (s *Server) Run(addr string) {
  log.Println("服务启动", addr)
  /*
  // for HTTP1
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    log.Printf("------收到RPC请求------")
    s.HandleRequest(w, r)
  })

  sx := &http.Server{
    Addr:       addr,
    Handler:    mux,
  }
  http2.ConfigureServer(sx, &http2.Server{})
  log.Fatal(sx.ListenAndServe())
   */

  h2s := &http2.Server{}
  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("------收到RPC请求------")
    s.HandleRequest(w, r)
  })
  server := &http.Server{
    Addr:    addr,
    Handler: h2c.NewHandler(handler, h2s),
  }
  fmt.Printf("Listening [%s]...\n", addr)
  server.ListenAndServe()
}





