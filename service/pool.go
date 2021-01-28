package service

import (
  "reflect"
  "sync"
)

type typePools struct {
  mu      sync.RWMutex
  pools   map[reflect.Type]*sync.Pool
  New     func(t reflect.Type) interface{}
}

var argsReplyPools = &typePools{
  pools:   make(map[reflect.Type]*sync.Pool),

  New:     func(t reflect.Type) interface{} {
    var argv reflect.Value

    if t.Kind() == reflect.Ptr {
      argv = reflect.New(t.Elem())
    } else {
      argv = reflect.New(t)
    }

    return argv.Interface()     // 返回类型的 interface{} 描述
  },
}

func (p *typePools) Init(t reflect.Type) {
  tp := &sync.Pool{}
  tp.New = func() interface{} {
    return p.New(t)
  }
  p.mu.Lock()
  defer p.mu.Unlock()
  p.pools[t] = tp
}