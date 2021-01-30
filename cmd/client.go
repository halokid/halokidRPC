package main

import (
  "bytes"
  "crypto/tls"
  "encoding/json"
  "github.com/halokid/ColorfulRabbit"
  "golang.org/x/net/http2"
  "io/ioutil"
  "log"
  "net"
  "net/http"
)

func main() {
  client := http.Client{
    Transport:     &http2.Transport{
      AllowHTTP:  true,
      DialTLS: func(network, addr string, cfg *tls.Config) (conn net.Conn, err error) {
        return net.Dial(network, addr)
      },
    },
  }

  url := "http://127.0.0.1:9527"
  payload, err := json.Marshal(map[string]string {
    "name": "halokid",
  })
  ColorfulRabbit.CheckError(err, "客户端序列化payload请求错误")

  req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
  req.Header.Add("svc", "Echo")
  req.Header.Add("method", "Say")
  rsp, _ := client.Do(req)
  log.Printf("RPC请求返回 --------- %+v", rsp)
  defer rsp.Body.Close()
  reply, _ := ioutil.ReadAll(rsp.Body)
  log.Printf("RPC请求返回reply --------- %+v", string(reply))
}