package service

import (
  "../codec"
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
  methods := FoundMethods(svcTyp)   // put in services.method, NumMethod只需要传入strutc的type定义
  t.Logf("methods ----------- %d, %+v", len(methods), methods)
  //methodName := "Say"
  //mtype := methods[methodName]

  mtype := methods["Say"]
  t.Logf("mtype ---------------- %+v", mtype)
  var argv = ArgsReplyPools.Get(mtype.ArgType)

  payload := `{"name": "halokid"}`
  payloadB := []byte(payload)      // 客户端传过来的数据
  codecx := &codec.JSONCodec{}
  err := codecx.Decode(payloadB, argv)    // 按照编码方式decode
  t.Log("err -----------", err)
  t.Logf("argv ----------- %+v", argv)

  replyv := ArgsReplyPools.Get(mtype.ReplyType)
  t.Logf("replyv ----------- %+v", replyv)

  s := &Service{}
  s.Rcvr = reflect.ValueOf(rcvr)    // todo: 取得struct的实体

  if mtype.ArgType.Kind() != reflect.Ptr {
    // 如果method arg的类型不是指针
    err = s.Call(mtype, reflect.ValueOf(argv).Elem(), reflect.ValueOf(replyv))
  } else {
    // 如果method arg的类型是指针
    err = s.Call(mtype, reflect.ValueOf(argv), reflect.ValueOf(replyv))
  }

  // 成功使用 argv, replyv之后， 放入pool复用
  ArgsReplyPools.Put(mtype.ArgType, argv)
  ArgsReplyPools.Put(mtype.ReplyType, replyv)

  t.Logf("replyv call after ------------- %+v, %+v", reflect.TypeOf(replyv), replyv)
  data, err := codecx.Encode(replyv)     // 服务端按照客户端的codec方式 encode 返回结果
  t.Logf("data ------------- %+v, %+v, %+v", reflect.TypeOf(data), data, string(data))
}



