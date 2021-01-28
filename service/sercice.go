package service

import (
  "log"
  "reflect"
  "sync"
  "unicode"
  "unicode/utf8"
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

// found methods can call for service
func FoundMethods(typ reflect.Type) map[string]*methodType {
  methods := make(map[string]*methodType)
  for m := 0; m < typ.NumMethod(); m++ {
    method := typ.Method(m)
    mtype := method.Type
    mname := method.Name
    // method 一定要允许外部调用
    if method.PkgPath != "" {
      continue
    }
    // todo: 省略一些methd的参数类型检查
    argType := mtype.In(2)      // 第二个参数
    if !isExportedOrBuiltinType(argType) {
      log.Println(mname, "的参数类型不允许外部调用:", argType)
      continue
    }

    replyType := mtype.In(3)
    if replyType.Kind() != reflect.Ptr {
      log.Println(mname, "的返回参数不是一个指针:", replyType)
      continue
    }

    // the return type of method must be error
    if returnType := mtype.Out(0); returnType != reflect.TypeOf((*error)(nil)).Elem() {
      log.Println("method", mname, "返回不是error类型")
      continue
    }

    methods[mname] = &methodType{method: method, ArgType: argType, ReplyType:  replyType}
  }
  return methods
}

// 检查类型是否允许外部访问或者是内置类型
func isExportedOrBuiltinType(t reflect.Type) bool {
  for t.Kind() == reflect.Ptr {
    t = t.Elem()
  }
  // PkgPat will be non-empty even for an exported type,
  // so we need to check the type name as well
  return isExported(t.Name()) || t.PkgPath() == ""
}

func isExported(name string) bool {
  runex, _ := utf8.DecodeRuneInString(name)
  return unicode.IsUpper(runex)    // 方法名首位是否大写
}





